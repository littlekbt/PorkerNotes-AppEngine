package hand_handler

import (
  "time"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/littlekbt/PorkerNotes-AppEngine/models/hand"
)

func Handle(w http.ResponseWriter, r *http.Request) {
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
		var h hand.Hand
    json.Unmarshal(buf.Bytes(), &h)
    for i, c := range h.Hands {
      b := c.MkBinary()
      switch i {
      case 0:
        h.Hand1 = b
      case 1:
        h.Hand2 = b
      }
    }

    for i, c := range h.Boards {
      b := c.MkBinary()
      switch i {
      case 0:
        h.Board1 = b
      case 1:
        h.Board2 = b
      case 2:
        h.Board3 = b
      }
    }
    h.CreatedAt = time.Now()
    h.UpdatedAt = time.Now()
    if !h.Valid() {
      json.NewEncoder(w).Encode(hand.Hand{})
      return
    }
    ih, err := h.Insert()
    if err != nil {
      log.Fatal(err)
    }
	  json.NewEncoder(w).Encode(ih)
	}
}
