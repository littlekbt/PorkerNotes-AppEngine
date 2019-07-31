package main

import (
  "net/http"

  "handlers/table_handler"
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
