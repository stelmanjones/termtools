package kv

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	json "github.com/bitly/go-simplejson"
	"github.com/charmbracelet/log"
	"github.com/emirpasic/gods/trees/btree"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
	"github.com/stelmanjones/termtools/kv/errors"
	"github.com/stelmanjones/termtools/styles"
)

type ServerOption func(*KV)

type KV struct {
	mu      *sync.RWMutex
	data    map[string]*btree.Tree
	auth    bool
	token   string
	address string
	port    int
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

var lvl = func() log.Level {
	lv := os.Getenv("LOG")
	switch strings.ToUpper(lv) {
	case "DEBUG":
		return log.DebugLevel
	case "INFO":
		return log.InfoLevel
	case "WARN":
		return log.WarnLevel
	case "ERROR":
		return log.ErrorLevel
	case "FATAL":
		return log.FatalLevel
	default:
		return log.InfoLevel
	}
}()

var logger = log.NewWithOptions(os.Stderr, log.Options{
	Level:           lvl,
	Prefix:          "EZKV-SERVER",
	ReportTimestamp: true,
})

func New(options ...ServerOption) *KV {
	k := &KV{
		mu:      &sync.RWMutex{},
		data:    make(map[string]*btree.Tree),
		auth:    false,
		token:   "",
		address: "127.0.0.1",
		port:    9999,
	}

	for _, option := range options {
		option(k)
	}

	return k
}

func WithAuth(token string) ServerOption {
	return func(k *KV) {
		k.auth = true
		k.token = token
		logger.Warn(styles.Warning.Styled("AUTH ENABLED"))

	}
}

func WithAddress(address string) ServerOption {
	return func(k *KV) {
		k.address = address
	}
}

func WithPort(port int) ServerOption {
	return func(k *KV) {
		k.port = port
	}
}

func WithRandomAuth() ServerOption {
	return func(k *KV) {
		k.auth = true
		k.token = generateRandomString(256)
		logger.Warn(styles.Warning.Styled("AUTH ENABLED"), "token", k.token)
	}
}

func (k *KV) Data() *map[string]*btree.Tree {
	return &k.data
}

func (k *KV) Set(table string, key string, value interface{}) {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.data[table] == nil {
		k.data[table] = btree.NewWithStringComparator(3)
	}
	k.data[table].Put(key, value)

	logger.Debug(styles.Warning.Styled("SET"), key, value)
}

func (k *KV) Get(table string, key string) (interface{}, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	if k.data[table] == nil {
		return nil, errors.ErrTableNotFound
	}
	value, found := k.data[table].Get(key)
	if !found {
		return value, errors.ErrKeyNotFound
	}
	logger.Debug(styles.AccentGreen.Styled("GET"), key, value)
	return value, nil
}

func (k *KV) Delete(table string, key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.data[table] == nil {
		return errors.ErrTableNotFound
	}

	k.data[table].Remove(key)
	logger.Debug(styles.AccentRed.Styled("DELETE"), key)
	return nil
}

func (k *KV) Keys() []map[string][]interface{} {
	k.mu.RLock()
	var keys []map[string][]interface{}
	defer k.mu.RUnlock()
	for table := range k.data {
		keys = append(keys, map[string][]interface{}{table: k.data[table].Keys()})
	}
	return keys
}

func (k *KV) Values() []map[string][]interface{} {
	k.mu.RLock()
	defer k.mu.RUnlock()
	var values []map[string][]interface{}
	defer k.mu.RUnlock()
	for table := range k.data {
		values = append(values, map[string][]interface{}{table: k.data[table].Values()})
	}
	return values
}

func (k *KV) Clear(table string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.data[table] == nil {
		return errors.ErrTableNotFound
	}

	k.data[table].Clear()
	logger.Warn(color.Bold.Sprintf("CLEARED TABLE %s", strings.ToUpper(table)))
	return nil
}

func (k *KV) Size() int {
	k.mu.RLock()
	defer k.mu.RUnlock()
	var size int
	for table := range k.data {
		size += k.data[table].Size()
	}
	return size
}

func (k *KV) handleGetKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	res, err := k.Get(params["table"], params["key"])
	if err != nil {
		logger.Error("GET ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	data := json.New()
	data.Set("result", map[string]interface{}{params["key"]: res})
	payload, err := data.MarshalJSON()
	if err != nil {
		logger.Error("GET ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (k *KV) handleSetKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	k.Set(params["table"], params["key"], params["value"])
	data := json.New()
	data.Set("result", map[string]interface{}{"table": params["table"], "key": params["key"], "value": params["value"]})
	payload, err := data.MarshalJSON()
	if err != nil {
		logger.Error("SET ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (k *KV) handleDeleteKey(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	k.Delete(p["table"], p["key"])
	data := json.New()
	data.Set("result", fmt.Sprintf("DELETED %s", p["key"]))
	payload, err := data.MarshalJSON()
	if err != nil {
		logger.Error("DELETE ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (k *KV) handleGetKv(w http.ResponseWriter, r *http.Request) {
	logger.WithPrefix("ADMIN").Info("GET KV")
	data := json.New()
	data.Set("result", k.Data())
	payload, err := data.MarshalJSON()
	if err != nil {
		logger.Error("DELETE ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (k *KV) handleClearKv(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	err := k.Clear(p["table"])
	if err != nil {
		logger.Error("CLEAR ERROR", "err", err)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logger.WithPrefix("ADMIN").Warnf("CLEARED TABLE %s", p["table"])
	data := json.New()
	data.Set("result", fmt.Sprintf("CLEARED TABLE %s", p["table"]))
	payload, err := data.MarshalJSON()
	if err != nil {
		logger.Error("DELETE ERROR", "err", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (k *KV) handleGetSize(w http.ResponseWriter, r *http.Request) {
	data := json.New()
	data.Set("result", map[string]interface{}{"size": k.Size()})
	payload, err := data.MarshalJSON()
	if err != nil {
		logger.Error("ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (k *KV) AuthMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			expected := fmt.Sprintf("Bearer %s", k.token)
			if token == "" || token != expected {
				logger.Warn("Unauthorized request", "path", r.URL.Path)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)

			return
		})
	}
}

func (k *KV) Serve() error {
	r := mux.NewRouter()
	//r.PathPrefix("/kv")

	r.HandleFunc("/kv/{table}/{key}", k.handleGetKey).Methods("GET")
	r.HandleFunc("/kv/{table}/{key}/{value}", k.handleSetKey).Methods("POST")
	r.HandleFunc("/kv/{table}/{key}", k.handleDeleteKey).Methods("DELETE")
	r.HandleFunc("/adm/kv", k.handleGetKv).Methods("GET")
	r.HandleFunc("/adm/kv/{table}", k.handleClearKv).Methods("DELETE")
	r.HandleFunc("/adm/size", k.handleGetSize).Methods("GET")
	r.Use(k.AuthMiddleware(r))
	fmt.Printf("%s\n\n", styles.Accent.Styled(Banner))
	logger.Debug("Server started 🎉", "address", k.address, "port", k.port, "auth", k.auth)
	logger.Fatal(http.ListenAndServe(strings.Join([]string{k.address, strconv.Itoa(k.port)}, ":"), r))

	return nil
}
