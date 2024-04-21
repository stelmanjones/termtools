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
	mu      *sync.RWMutex // Guards access to the data map.
	data    *hashmap.Map  // The data stored in the key-value store.
	auth    bool          // Indicates if authentication is enabled.
	token   string        // The authentication token required if auth is enabled.
	address string        // The network address the server will listen on.
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
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.data
}

// Set stores a value associated with a key in the KV store.
func (k *KV) Set(key string, value interface{}) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data.Put(key, value)

	logger.Debug(styles.Warning.Styled("SET"), key, value)
}

// Get retrieves a value associated with a key from the KV store.
func (k *KV) Get(key string) (interface{}, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()
	if value, found := k.data.Get(key); found {
		logger.Debug(styles.AccentGreen.Styled("GET"), key, value)
		return value, nil
	}
	return nil, errors.ErrKeyNotFound
}

// Delete removes a key and its associated value from the KV store.
func (k *KV) Delete(key string) error {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data.Remove(key)
	logger.Debug(styles.AccentRed.Styled("DELETE"), key)
	return nil
}

// Keys returns a slice of all keys currently stored in the KV store.
func (k *KV) Keys() []any {
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

// handleDeleteKey processes HTTP DELETE requests for removing a key-value pair.
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

// handleKvData processes HTTP GET requests for retrieving all key-value pairs.
func (k *KV) handleKvData(w http.ResponseWriter, _ *http.Request) {
	logger.WithPrefix("ADMIN").Info("GET KV")
	payload, err := k.data.ToJSON()
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
	data.Set("result", "CLEARED TABLE")
	payload, err := data.MarshalJSON()
	if err != nil {
		logger.Error("DELETE ERROR", "err", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// handleGetKvSize processes HTTP GET requests for retrieving the size of the store.
func (k *KV) handleGetKvSize(w http.ResponseWriter, _ *http.Request) {
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
			if !k.auth {
				next.ServeHTTP(w, r)
				return
			}
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
	

	r.HandleFunc("/kv/{key}", k.handleGetKey).Methods("GET")
	r.HandleFunc("/kv/{key}/{value}", k.handleSetKey).Methods("POST")
	r.HandleFunc("/kv/{key}", k.handleDeleteKey).Methods("DELETE")
	r.HandleFunc("/adm/kv", k.handleKvData).Methods("GET")
	r.HandleFunc("/adm/kv", k.handleClearKv).Methods("DELETE")
	r.HandleFunc("/adm/size", k.handleGetKvSize).Methods("GET")
	r.Use(k.AuthMiddleware(r))
	fmt.Printf("%s\n\n", styles.Accent.Styled(banner))
	logger.Debug("Server started 🎉", "address", k.address, "port", port, "auth", k.auth)
	err := http.ListenAndServe(strings.Join([]string{k.address, strconv.Itoa(port)}, ":"), r)
	if err != nil {
		return err
	}

	return nil
}
