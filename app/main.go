package main

import (
	"encoding/json"
	"net/http"

	"github.com/littlekbt/PorkerNotes-AppEngine/handlers/table_handler"
	"github.com/littlekbt/PorkerNotes-AppEngine/models/response"
	"google.golang.org/appengine"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/tables", table_handler.Handle)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(response.Response{Status: "ok", Message: "Hello world."})
}
