package helpers

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, code int, data interface{}){
	response,_:=json.Marshal(data)
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(response)
}