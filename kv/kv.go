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
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
	"github.com/stelmanjones/termtools/kv/errors"
	"github.com/stelmanjones/termtools/styles"
)

type ServerOption func(*KV)

type KV struct {
	mu      *sync.RWMutex
	data    *hashmap.Map
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
		data:    hashmap.New(),
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

func (k *KV) Data() *hashmap.Map {
	return k.data
}
func (k *KV) Set(key string, value interface{}) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data.Put(key, value)

	logger.Debug(styles.Warning.Styled("SET"), key, value)
}

func (k *KV) Get(table string, key string) (interface{}, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	if value, found := k.data.Get(key); found != false {
		logger.Debug(styles.AccentGreen.Styled("GET"), key, value)
		return value, nil
	}
	return nil, errors.ErrKeyNotFound
}

func (k *KV) Delete(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data.Remove(key)
	logger.Debug(styles.AccentRed.Styled("DELETE"), key)
	return nil
}

func (k *KV) Keys() []interface{} {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.data.Keys()
}

func (k *KV) Values() []interface{} {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.data.Values()
}

func (k *KV) Clear() error {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data.Clear()
	logger.Warn(color.Bold.Sprint("CLEARED TABLE"))
	return nil
}

func (k *KV) Size() int {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.data.Size()
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

	k.Set(params["key"], params["value"])
	data := json.New()
	data.Set("result", map[string]interface{}{"key": params["key"], "value": params["value"]})
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
	k.Delete(p["key"])
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
	payload, err := k.data.MarshalJSON()
	if err != nil {
		logger.Error("DELETE ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (k *KV) handleClearKv(w http.ResponseWriter, r *http.Request) {
	err := k.Clear()
	if err != nil {
		logger.Error("CLEAR ERROR", "err", err)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logger.WithPrefix("ADMIN").Warn("CLEARED TABLE")
	data := json.New()
	data.Set("result", fmt.Sprint("CLEARED TABLE"))
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
		})
	}
}

func (k *KV) Serve() error {
	r := mux.NewRouter()
	// r.PathPrefix("/kv")

	r.HandleFunc("/kv/{key}", k.handleGetKey).Methods("GET")
	r.HandleFunc("/kv/{key}/{value}", k.handleSetKey).Methods("POST")
	r.HandleFunc("/kv/{key}", k.handleDeleteKey).Methods("DELETE")
	r.HandleFunc("/adm/kv", k.handleGetKv).Methods("GET")
	r.HandleFunc("/adm/kv", k.handleClearKv).Methods("DELETE")
	r.HandleFunc("/adm/size", k.handleGetSize).Methods("GET")
	r.Use(k.AuthMiddleware(r))
	fmt.Printf("%s\n\n", styles.Accent.Styled(Banner))
	logger.Debug("Server started 🎉", "address", k.address, "port", k.port, "auth", k.auth)
	logger.Fatal(http.ListenAndServe(strings.Join([]string{k.address, strconv.Itoa(k.port)}, ":"), r))

	return nil
}
