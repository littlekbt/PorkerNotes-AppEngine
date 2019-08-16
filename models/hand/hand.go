package hand

import (
  // "log"
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
  BlindL int64 `json:"blind_l" db:"blind_l"`
  BlindR int64 `json:"blind_r" db:"blind_r"`
  Ante int64 `json:"ante" db:"ante"`
  PreFlop []action.Action `json:"pre_flop"`
  Hands []card.Card `json:"hands"`
  Boards []card.Card `json:"boards"`
  Hand1 int `db:"hand1"`
  Hand2 int `db:"hand2"`
  Board1 int `db:"board1"`
  Board2 int `db:"board2"`
  Board3 int `db:"board3"`
  Board4 int `db:"board4"`
  Board5 int `db:"board5"`
  Actions []action.Action `json:"actions"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (h *Hand)NewCards() {
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
    case 3:
      h.Board4 = b
    case 4:
      h.Board5 = b
    }
  }
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
  h.CreatedAt = time.Now()
  h.UpdatedAt = time.Now()
  db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3307)/porker_notes?parseTime=true&loc=Asia%2FTokyo")
  defer db.Close()
	if err != nil {
		return Hand{}, err
	}

  // TODO: Transaction
	ins, err := db.Prepare("INSERT INTO hands(table_id, name, pos, blind_l, blind_r, ante, hand1, hand2, board1, board2, board3, board4, board5, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return Hand{}, err
	}
  r, err := ins.Exec(h.TableID, h.Name, h.Pos, h.BlindL, h.BlindR, h.Ante, h.Hand1, h.Hand2, h.Board1, h.Board2, h.Board3, h.Board4, h.Board5, h.CreatedAt, h.UpdatedAt) 
  if err != nil {
		return Hand{}, err
  }
  id, _ := r.LastInsertId()
  h.ID = id

  for i, pf := range h.PreFlop {
    pf.HandID = id
    pf.Type = action.PRE_FLOP
    p, err := pf.Insert()
    if err != nil {
      return Hand{}, err
    }
    h.PreFlop[i] = p
  }
  for i, act := range h.Actions {
    act.HandID = id
    act.Type = action.BET_ROUND
    a, err := act.Insert()
    if err != nil {
      return Hand{}, err
    }
    h.Actions[i] = a
  }
	return h, nil
}

func (h Hand) Valid() bool {
  if h.TableID == 0 || h.Name == "" {
    return false
  }
  return true
}
