// Package kv provides a key-value store with optional authentication and HTTP server functionality.
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

// Option defines a function signature for options used to configure a KV instance.
type Option func(*KV)

// KV represents a key-value store with optional authentication and network address configuration.
type KV struct {
	mu      *sync.RWMutex
	data    *hashmap.Map
	auth    bool
	token   string
	address string
	limit   int
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// generateRandomString creates a random string of a specified length using a predefined charset.
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

// New creates a new KV instance with the provided options.
func New(options ...Option) *KV {
	k := &KV{
		mu:      &sync.RWMutex{},
		data:    hashmap.New(),
		auth:    false,
		token:   "",
		address: "127.0.0.1",
		limit:   10000,
	}

	for _, option := range options {
		option(k)
	}

	return k
}

// WithAuth configures the KV instance to require authentication with the provided token.
func WithAuth(token string) Option {
	return func(k *KV) {
		k.auth = true
		k.token = token
		logger.Warn(styles.Warning.Styled("AUTH ENABLED"))
	}
}

// WithAddress sets the network address for the KV instance.
func WithAddress(address string) Option {
	return func(k *KV) {
		k.address = address
	}
}

// WithLimit sets the maximum number of items allowed in the KV store.
func WithLimit(limit int) Option {
	return func(k *KV) {
		k.limit = limit
	}
}

// WithRandomAuth configures the KV instance to require authentication with a randomly generated token.
func WithRandomAuth() Option {
	return func(k *KV) {
		k.auth = true
		k.token = generateRandomString(256)
		logger.Warn(styles.Warning.Styled("AUTH ENABLED"), "token", k.token)
	}
}

// Data returns a snapshot of the current data in the KV store.
func (k *KV) Data() *hashmap.Map {
	return k.data
}

// Set stores a value associated with a key in the KV store.
func (k *KV) Set(key string, value interface{}) error {
	if k.limit > 0 && k.data.Size() >= k.limit {
		logger.Warn(styles.Warning.Styled("TABLE FULL"))
		return errors.ErrTableFull
	}
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data.Put(key, value)

	logger.Debug(styles.Warning.Styled("SET"), key, value)
	return nil
}

// Get retrieves a value associated with a key from the KV store.
func (k *KV) Get(key string) (interface{}, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	if value, found := k.data.Get(key); found != false {
		logger.Debug(styles.AccentGreen.Styled("GET"), key, value)
		return value, nil
	}
	return nil, errors.ErrKeyNotFound
}

// Has checks if a key exists in the KV store.
func (k *KV) Has(key string) bool {
	k.mu.RLock()
	defer k.mu.RUnlock()
	_, ok := k.data.Get(key)
	return ok
}

// Remove removes a key and its associated value from the KV store.
func (k *KV) Remove(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data.Remove(key)
	logger.Debug(styles.AccentRed.Styled("DELETE"), key)
	return nil
}

// Keys returns a slice of all keys currently stored in the KV store.
func (k *KV) Keys() []interface{} {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.data.Keys()
}

// Values returns a slice of all values currently stored in the KV store.
func (k *KV) Values() []interface{} {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.data.Values()
}

// Clear removes all keys and values from the KV store.
func (k *KV) Clear() error {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data.Clear()
	logger.Warn(color.Bold.Sprint("CLEARED TABLE"))
	return nil
}

// Size returns the number of items currently stored in the KV store.
func (k *KV) Size() int {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.data.Size()
}

// ToJSON returns the KV store data as a JSON string.
func (k *KV) ToJSON() ([]byte, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	j := json.New()
	j.Set("data", k.data)
	return j.MarshalJSON()
}

// handleGetKey processes HTTP GET requests for retrieving a value by key.
func (k *KV) handleGetKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	res, err := k.Get(params["key"])
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

// handleSetKey processes HTTP POST requests for setting a key-value pair.
func (k *KV) handleSetKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	err := k.Set(params["key"], params["value"])
	if err != nil {
		logger.Error("SET ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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

// handleRemoveKey processes HTTP DELETE requests for removing a key-value pair.
func (k *KV) handleRemoveKey(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	k.Remove(p["key"])
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

// handleKvData processes HTTP GET requests for retrieving all key-value pairs.
func (k *KV) handleGetKv(w http.ResponseWriter, _ *http.Request) {
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

// handleClearKv processes HTTP DELETE requests for clearing all key-value pairs.
func (k *KV) handleClearKv(w http.ResponseWriter, _ *http.Request) {
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

// handleGetKvSize processes HTTP GET requests for retrieving the size of the store.
func (k *KV) handleGetSize(w http.ResponseWriter, _ *http.Request) {
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

// AuthMiddleware returns a middleware function that enforces authentication for HTTP requests.
func (k *KV) AuthMiddleware(_ *mux.Router) mux.MiddlewareFunc {
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

// Serve starts the HTTP server on the specified port with configured routes and middleware.
func (k *KV) Serve(port int) error {
	r := mux.NewRouter()
	// r.PathPrefix("/kv")

	r.HandleFunc("/kv/{key}", k.handleGetKey).Methods("GET")
	r.HandleFunc("/kv/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
	r.HandleFunc("/kv/{key}/{value}", k.handleSetKey).Methods("POST")
	r.HandleFunc("/kv/{key}", k.handleRemoveKey).Methods("DELETE")
	r.HandleFunc("/adm/kv", k.handleGetKv).Methods("GET")
	r.HandleFunc("/adm/kv", k.handleClearKv).Methods("DELETE")
	r.HandleFunc("/adm/size", k.handleGetSize).Methods("GET")
	r.Use(k.AuthMiddleware(r))
	fmt.Printf("%s\n\n", styles.Accent.Styled(banner))
	logger.Debug("Server started 🎉", "address", k.address, "port", port, "auth", k.auth)
	logger.Fatal(http.ListenAndServe(strings.Join([]string{k.address, strconv.Itoa(port)}, ":"), r))

	return nil
}
