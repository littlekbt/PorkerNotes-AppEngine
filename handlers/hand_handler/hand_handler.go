package hand_handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/littlekbt/PorkerNotes-AppEngine/models/hand"
)

func Handle(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")

	switch r.Method {
  case "GET":
    hands, err := hand.Select(0)
    if err != nil {
      log.Fatal(err)
    }
    json.NewEncoder(w).Encode(hands)
  case "POST":
		body := r.Body
		defer body.Close()

		buf := new(bytes.Buffer)
		io.Copy(buf, body)
    var hs []hand.Hand
    var hss []hand.Hand
    json.Unmarshal(buf.Bytes(), &hs)
    
    for _, h := range hs {
      h.NewCards()
      if !h.Valid() {
        json.NewEncoder(w).Encode(hand.Hand{})
        return
      }
      ih, err := h.Insert()
      if err != nil {
        log.Fatal(err)
      }
      hss = append(hss, ih)
    }
    json.NewEncoder(w).Encode(hss)
  case "OPTIONS":
    json.NewEncoder(w).Encode(hand.Hand{})
	}
}
