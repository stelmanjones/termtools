package examples

import (
	"github.com/charmbracelet/log"
	"github.com/stelmanjones/termtools/kv"
)

func kvTest() {

	db := kv.New().
		WithAuth("kekw1337").
		WithLimit(1000).
		Build()
	log.Fatal(db.Serve(9999))

	db.Get("key")
	db.Set("key", "value")
	db.Remove("key")
	db.Clear()
	db.Keys()
	db.Values()
	db.Data()
}
