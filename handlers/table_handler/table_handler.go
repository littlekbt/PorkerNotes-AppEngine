package table_handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/littlekbt/PorkerNotes-AppEngine/models/table"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	tables := []table.Table{}
	switch r.Method {
	case "GET":
		tables = append(tables, table.Table{ID: int64(1)})
	case "POST":
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)
		var t table.Table
		json.Unmarshal(buf.Bytes(), &t)

		db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/porker_notes")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		ins, err := db.Prepare("INSERT INTO tables(name) VALUES(?)")
		if err != nil {
			log.Fatal(err)
		}
		ins.Exec(t.Name)
		tables = append(tables, t)
	}
	json.NewEncoder(w).Encode(tables)
}
