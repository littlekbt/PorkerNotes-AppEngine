package hand_test

import (
  "fmt"
  "testing"
  "os"
  "io/ioutil"
  "encoding/json"

  "github.com/littlekbt/PorkerNotes-AppEngine/models/hand"
)

func TestInsert(t *testing.T) {
  f, _ := os.Open("hand_test.json")
  byteValue, _ := ioutil.ReadAll(f)
  h := hand.Hand{}
  json.Unmarshal([]byte(byteValue), &h)
  h.NewCards()
  afterH, err := h.Insert()
  if err != nil {
    t.Errorf("Error: %s", err.Error())
  }
  fmt.Println(afterH)
}