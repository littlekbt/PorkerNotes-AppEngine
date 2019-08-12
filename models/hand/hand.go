package hand

import (
  "time"
  "strconv"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
  "github.com/littlekbt/PorkerNotes-AppEngine/models/action"
  "github.com/littlekbt/PorkerNotes-AppEngine/models/card"
  // "github.com/littlekbt/PorkerNotes-AppEngine/models/porker_notes"
)

type Hand struct {
  ID int64 `json:"id" db:"id, primarykey"`
  TableID int64 `json:"table_id" db:"table_id"`
  Name string `json:"name" db:"name"`
  Pos int64 `json:"pos" db:"pos"`
  Tip int64 `json:"tip" db:"tip"`
  Hands []card.Card `json:"hands"`
  Boards []card.Card `json:"boards"`
  Hand1 int `db:"hand1"`
  Hand2 int `db:"hand2"`
  Board1 int `db:"board1"`
  Board2 int `db:"board2"`
  Board3 int `db:"board3"`
  Actions []action.Action `json:"actions"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}


func Select(offset int) ([]Hand, error) {
  limit := 10
  hands := []Hand{}
  db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3307)/porker_notes?parseTime=true&loc=Asia%2FTokyo")
  defer db.Close()
	if err != nil {
		return hands, err
	}

	rows, err := db.Query("SELECT id, name FROM hands limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset))
	if err != nil {
		return hands, err
  }
  for rows.Next() {
    h := Hand{}
    if err := rows.Scan(&h.ID,&h.Name); err != nil {
      return hands, err
    }
    hands = append(hands, h)
  }

  return hands, nil
}

func (h Hand) Insert() (Hand, error){
  db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3307)/porker_notes?parseTime=true&loc=Asia%2FTokyo")
  defer db.Close()
	if err != nil {
		return h, err
	}

  // TODO: Transaction
	ins, err := db.Prepare("INSERT INTO hands(table_id, name, pos, tip, hand1, hand2, board1, board2, board3, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return h, err
	}
  ins.Exec(h.TableID, h.Name, h.Pos, h.Tip, h.Hand1, h.Hand2, h.Board1, h.Board2, h.Board3, h.CreatedAt, h.UpdatedAt)
  

  for _, action := range h.Actions {
    ins, err = db.Prepare("INSERT INTO actions(hand_id, created_at, updated_at) VALUES(?, ?, ?)")
    if err != nil {
		  return h, err
	  }
    ins.Exec(action.HandID, h.CreatedAt, h.UpdatedAt)
  }
	return h, nil
}

func (h Hand) Valid() bool {
  if h.TableID == 0 || h.Name == "" {
    return false
  }
  return true
}
