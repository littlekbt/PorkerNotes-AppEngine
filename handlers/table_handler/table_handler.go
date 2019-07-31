package table_handler

import (
	"encoding/json"
	"net/http"

	//"github.com/littlekbt/PorkerNotes-AppEngine/models/response"
	"github.com/littlekbt/PorkerNotes-AppEngine/models/table"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	tables := make([]table.Table, 1)
	tables[0] = table.Table{ID: int64(1)}
	json.NewEncoder(w).Encode(tables)
	//json.NewEncoder(w).Encode(response.Response{Status: "ok", Message: "Hello world."})
}
