package kv

import (
	"net/http"
	

	"github.com/wI2L/jettison"
)

func kvResult(data interface{}) ([]byte, error) {
	return jettison.Marshal(map[string]interface{}{"result": data})
}

func kvResponse(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
