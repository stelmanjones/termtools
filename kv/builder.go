package kv

import (
	"math/rand"
	"sync"
	"time"

	"github.com/emirpasic/gods/maps/hashmap"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Builder is a builder for KV.
type Builder struct {
	token   string
	address string
	limit   int
	auth    bool
}

// WithAuth sets the authentication flag and token for the KV.
func (b *Builder) WithAuth(token string) *Builder {
	b.auth = true
	b.token = token
	return b
}

// WithAddress sets the address for the KV.
func (b *Builder) WithAddress(address string) *Builder {
	b.address = address
	return b
}

// WithLimit sets the max number of records that can be stored in the KV.
func (b *Builder) WithLimit(limit int) *Builder {
	b.limit = limit
	return b
}

// WithRandomBearerToken sets a random bearer token for the KV.
func WithRandomBearerToken() *Builder {
	return &Builder{
		auth:  true,
		token: generateRandomString(32),
	}
}

// Build returns a new KV instance with the configured options.
func (b *Builder) Build() *KV {
	return &KV{
		mux:  &sync.RWMutex{},
		data: hashmap.New(),

		auth:    b.auth,
		token:   b.token,
		address: b.address,
		limit:   b.limit,
		batch:   []interface{}{},
	}
}

// New returns a new KV builder.
func New() *Builder {
	return &Builder{}
}

// generateRandomString creates a random string of a specified length using a predefined charset.
func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
