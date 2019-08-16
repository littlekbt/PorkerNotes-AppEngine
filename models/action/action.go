package action

import (
  "time"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
)

type Type string
const (
  PRE_FLOP = "pre_flop"
  BET_ROUND = "bet_round"
)

type Event int
const (
  _ Event = iota // ignore first value by assigning to blank identifier
  CALL
  CHECK
  FOLD
  RAISE
  BET
)

type Action struct {
  ID int64 `json:"id" db:"id, primarykey"`
  HandID int64 `json:"hand_id" db:"hand_id"`
  Person string `json:"person" db:"person"`
  Type Type `json:"type" db:"type"`
  Event Event `json:"event" db:"event"`
  Bet int64 `json:"bet" db:"bet"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (a Action) Insert() (Action, error) {
  a.CreatedAt = time.Now()
  a.UpdatedAt = time.Now()

  db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3307)/porker_notes?parseTime=true&loc=Asia%2FTokyo")
  defer db.Close()
	if err != nil {
		return Action{}, err
	}

  ins, err := db.Prepare("INSERT INTO actions(hand_id, person, type, event, bet, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)")
  if err != nil {
		return Action{}, err
	}
  r, err := ins.Exec(a.HandID, a.Person, a.Type, a.Event, a.Bet, a.CreatedAt, a.UpdatedAt)
  if err != nil {
		return Action{}, err
	}
  id, _ := r.LastInsertId()
  a.ID = id

  return a, nil

}