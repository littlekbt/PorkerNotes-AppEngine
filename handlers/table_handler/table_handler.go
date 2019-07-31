package table_handler

import (
  "net/http"
  "encoding/json"

  "models/response"
)

func Handle(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(response.Response{Status: "ok", Message: "Hello world."})
}