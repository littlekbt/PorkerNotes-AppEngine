package table_handler

import (
	"bytes"
	"encoding/json"
	"io"
  "log"
  "time"
	"net/http"

	"github.com/littlekbt/PorkerNotes-AppEngine/models/table"
)

func Handle(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")

	switch r.Method {
  case "GET":
    tables, err := table.Select(0)
    if err != nil {
      log.Fatal(err)
    }
    json.NewEncoder(w).Encode(tables)
  case "POST":
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
    io.Copy(buf, body)
    var t table.Table
    json.Unmarshal(buf.Bytes(), &t)
    t.CreatedAt = time.Now()
    t.UpdatedAt = time.Now()
    if !t.Valid() {
      json.NewEncoder(w).Encode(table.Table{})
      return
    }
    it, err := t.Insert()
    if err != nil {
      log.Fatal(err)
    }
    json.NewEncoder(w).Encode(it)
  case "OPTIONS":
    json.NewEncoder(w).Encode(table.Table{})
	}
}
