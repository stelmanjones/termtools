package examples

import ( 
	"github.com/stelmanjones/termtools/kv"
	"github.com/charmbracelet/log"
)

func kvTest() {
	db := kv.New()
	log.Fatal(db.Serve(9999))


	db.Get("key")
	db.Set("key", "value")
	db.Delete("key")
	db.Clear()
	db.Keys()
	db.Values()
	db.Data()
}