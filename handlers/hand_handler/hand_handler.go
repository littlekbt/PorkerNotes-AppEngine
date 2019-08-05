package hand_handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/littlekbt/PorkerNotes-AppEngine/models/hand"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	hands := []hand.Hand{}
	switch r.Method {
	case "GET":
		hands = append(hands, hand.Hand{ID: int64(1)})
	case "POST":
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)
		var h hand.Hand
		json.Unmarshal(buf.Bytes(), &h)

		db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/porker_notes")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		ins, err := db.Prepare("INSERT INTO hands(name) VALUES(?)")
		if err != nil {
			log.Fatal(err)
		}
		ins.Exec(h.Name)
		hands = append(hands, h)
	}
	json.NewEncoder(w).Encode(hands)
}
