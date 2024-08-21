package kv

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	sjson "github.com/bitly/go-simplejson"
	"github.com/charmbracelet/log"
	"github.com/emirpasic/gods/maps/hashmap"
	"github.com/gookit/color"
	"github.com/gorilla/mux"
	"github.com/stelmanjones/termtools/internal/theme"
	"github.com/stelmanjones/termtools/kv/errors"
)

// Option defines a function signature for options used to configure a KV instance.
type Option func(*KV)

// KV represents a key-value store with optional authentication and network address configuration.
type KV struct {
	mux     *sync.RWMutex
	data    *hashmap.Map
	token   string
	address string
	batch   []interface{}
	limit   int
	auth    bool
}

// func (k *KV) AddBatch(keyvals ...interface{}) error {
// 	k.mux.Lock()
// 	defer k.mux.Unlock()
// 	if len(keyvals)%2 != 0 {
// 		return errors.ErrMissingValue
// 	}
// 	logger.Debug("ADDED TO BATCH", "keyvals", keyvals)
// 	k.batch = append(k.batch, keyvals...)
// 	return nil
// }
//
// func (k *KV) CancelBatch() {
// 	k.mux.Lock()
// 	defer k.mux.Unlock()
// 	logger.Debug("Cancelled batch")
// 	k.batch = []interface{}{}
// }
//
// func (k *KV) CommitBatch() {
// 	k.mux.Lock()
// 	defer k.mux.Unlock()
// 	for i := 0; i < len(k.batch); i += 2 {
// 		k.data.Put(k.batch[i], k.batch[i+1])
// 	}
// 	logger.Debug("Committed batch")
// 	k.batch = []interface{}{}
// }

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
	Prefix:          "KV",
	ReportTimestamp: true,
})

// Data returns a snapshot of the current data in the KV store.
func (k *KV) Data() *hashmap.Map {
	return k.data
}

// Set stores a value associated with a key in the KV store.
func (k *KV) Set(key string, value interface{}) error {
	if k.limit > 0 && k.data.Size() >= k.limit {
		logger.Warn(theme.Warning.Render("TABLE FULL"))
		return errors.ErrTableFull
	}
	k.mux.Lock()
	defer k.mux.Unlock()
	if !k.Has(key) {
		k.data.Put(key, value)
		logger.Debug(theme.Warning.Render("SET"), key, value)
		return nil
	}
	logger.Error("Key '%v' already exists.", key)
	return errors.ErrKeyExists
}

// SetMany stores multiple key-value pairs in the KV store.
func (k *KV) SetMany(keyvals ...interface{}) error {
	if len(keyvals)%2 != 0 {
		return errors.ErrMissingValue
	}
	k.mux.Lock()
	defer k.mux.Unlock()
	for i := 0; i < len(keyvals); i += 2 {

		if k.limit > 0 && k.data.Size() >= k.limit {
			logger.Warn(theme.Warning.Render("TABLE FULL"))
			return errors.ErrTableFull
		}
		if !k.Has(keyvals[i].(string)) {
			k.data.Put(keyvals[i].(string), keyvals[i+1])
			logger.Debug(theme.Warning.Render("SET"), keyvals[i].(string), keyvals[i+1])
		}
		logger.Error("Key '%v' already exists.", keyvals[i].(string))
		return errors.ErrKeyExists
	}
	return nil
}

// Update updates the value associated with a key in the KV store.
func (k *KV) Update(key string, value interface{}) error {
	k.mux.Lock()
	defer k.mux.Unlock()
	if !k.Has(key) {
		return errors.ErrKeyNotFound
	}
	k.data.Put(key, value)
	logger.Debug(theme.AccentBlue.Render("UPDATED"), key, value)
	return nil
}

// UpdateMany updates multiple key-value pairs in the KV store.
func (k *KV) UpdateMany(keyvals ...interface{}) error {
	if len(keyvals)%2 != 0 {
		return errors.ErrMissingValue
	}
	for i := 0; i < len(keyvals); i += 2 {
		if !k.Has(keyvals[i].(string)) {
			logger.Error("Key '%v' does not exist.", keyvals[i].(string))
			return errors.ErrKeyNotFound
		}
		k.data.Put(keyvals[i].(string), keyvals[i+1])
		logger.Debug(theme.AccentBlue.Render("UPDATED"), keyvals[i].(string), keyvals[i+1])
	}
	return nil
}

// Get retrieves a value associated with a key from the KV store.
func (k *KV) Get(key string) (interface{}, error) {
	k.mux.RLock()
	defer k.mux.RUnlock()
	if value, found := k.data.Get(key); found {
		logger.Debug(theme.AccentGreen.Render("GET"), key, value)
		return value, nil
	}
	return nil, errors.ErrKeyNotFound
}

// GetMany retrieves multiple values from the KV store.
func (k *KV) GetMany(keys ...string) ([]interface{}, error) {
	k.mux.RLock()
	defer k.mux.RUnlock()
	var res []interface{}
	for _, key := range keys {
		if value, found := k.data.Get(key); found {
			res = append(res, value)
		}
	}
	logger.Debug(theme.AccentGreen.Render("GET"), "result", res)
	return res, nil
}

// Has checks if a key exists in the KV store.
func (k *KV) Has(key string) bool {
	k.mux.RLock()
	defer k.mux.RUnlock()
	_, ok := k.data.Get(key)
	return ok
}

// Remove removes a key and its associated value from the KV store.
func (k *KV) Remove(key string) error {
	k.mux.Lock()
	defer k.mux.Unlock()
	k.data.Remove(key)
	logger.Debug(theme.AccentRed.Render("DELETE"), key)
	return nil
}

// RemoveMany removes multiple keys and their associated values from the KV store.
func (k *KV) RemoveMany(keys ...string) error {
	for _, key := range keys {
		k.Remove(key)
	}
	return nil
}

// Keys returns a slice of all keys currently stored in the KV store.
func (k *KV) Keys() []interface{} {
	k.mux.RLock()
	defer k.mux.RUnlock()
	return k.data.Keys()
}

// Values returns a slice of all values currently stored in the KV store.
func (k *KV) Values() []interface{} {
	k.mux.RLock()
	defer k.mux.RUnlock()
	return k.data.Values()
}

// Clear removes all keys and values from the KV store.
func (k *KV) Clear() error {
	k.mux.Lock()
	defer k.mux.Unlock()
	k.data.Clear()
	logger.Warn(color.Bold.Sprint("CLEARED TABLE"))
	return nil
}

// Size returns the number of items currently stored in the KV store.
func (k *KV) Size() int {
	k.mux.RLock()
	defer k.mux.RUnlock()
	return k.data.Size()
}

// ToJSON returns the KV store data as a JSON string.
func (k *KV) ToJSON() ([]byte, error) {
	k.mux.RLock()
	defer k.mux.RUnlock()
	j := sjson.New()
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
	payload, err := kvResult(map[string]interface{}{params["key"]: res})
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
	payload, err := kvResult(map[string]interface{}{params["key"]: params["value"]})
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
	payload, err := kvResult(fmt.Sprintf("DELETED %s", p["key"]))
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
	payload, err := kvResult("CLEARED TABLE")
	if err != nil {
		logger.Error("DELETE ERROR", "err", err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// handleGetKvSize processes HTTP GET requests for retrieving the size of the store.
func (k *KV) handleGetSize(w http.ResponseWriter, _ *http.Request) {
	payload, err := kvResult(map[string]interface{}{"size": k.Size()})
	if err != nil {
		logger.Error("ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (k *KV) handleJSON(w http.ResponseWriter, r *http.Request) {
	data, err := sjson.NewFromReader(r.Body)
	if err != nil {
		logger.Error(err.Error())
	}
	inserted := sjson.New()
	if v, err := data.Map(); err == nil {
		for key, val := range v {
			k.Set(key, val)
			inserted.Set(key, val)
		}
	} else if v, err := data.Array(); err == nil {
		for _, val := range v {
			switch val := val.(type) {
			case map[string]any:
				for key, val := range val {
					k.Set(key, val)
					inserted.Set(key, val)
				}
			}
		}
	} else {
		logger.Error("ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	payload, err := kvResult(inserted.Interface())
	if err != nil {
		logger.Error("ERROR", "err", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// TODO: Improve batching function.

func (k *KV) handleBatch(w http.ResponseWriter, r *http.Request) {
	inserted := sjson.New()
	switch r.Method {
	case "POST":
		data, err := sjson.NewFromReader(r.Body)
		if err != nil {
			logger.Error("ERROR", "err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if _, err := data.Map(); err == nil {

			b, err := data.MarshalJSON()
			if err != nil {
				logger.Error("ERROR", "err", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			k.data.FromJSON(b)
			inserted.UnmarshalJSON(b)
			logger.Debug(inserted.Interface())

		} else if v, err := data.Array(); err == nil {
			keyvals := make([]interface{}, 0)
			for _, val := range v {
				switch val := val.(type) {
				case map[string]any:
					for key, val := range val {
						keyvals = append(keyvals, key, val)
					}
					k.SetMany(keyvals...)
				}
			}

		} else {
			logger.Error("ERROR", "err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		payload, err := kvResult(inserted.Interface())
		if err != nil {
			logger.Error("ERROR", "err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)

	case "GET":
		payload, err := kvResult(k.batch)
		if err != nil {
			logger.Error("ERROR", "err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	default:
		http.Error(w, "Invalid method", http.StatusBadRequest)
	}
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
	//
	r.HandleFunc("/kv", k.handleJSON).Methods("POST")
	r.HandleFunc("/kv/{key}", k.handleGetKey).Methods("GET")
	r.HandleFunc("/kv/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")
	r.HandleFunc("/kv/tx", k.handleBatch).Methods("POST", "GET")
	r.HandleFunc("/kv/{key}/{value}", k.handleSetKey).Methods("POST")
	r.HandleFunc("/kv/{key}", k.handleRemoveKey).Methods("DELETE")
	r.HandleFunc("/adm/kv", k.handleGetKv).Methods("GET")
	r.HandleFunc("/adm/kv", k.handleClearKv).Methods("DELETE")
	r.HandleFunc("/adm/size", k.handleGetSize).Methods("GET")
	r.Use(k.AuthMiddleware(r))
	fmt.Printf("%s\n\n", theme.Accent.Render(banner))
	logger.Debug("Server started 🎉", "address", k.address, "port", port, "auth", k.auth)
	logger.Fatal(http.ListenAndServe(strings.Join([]string{k.address, strconv.Itoa(port)}, ":"), r))

	return nil
}
