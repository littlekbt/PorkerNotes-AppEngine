package action

import (
  "time"
)

type Action struct {
  ID int64  `json:"id" db:"id, primarykey"`
  HandID int64 `json:"hand_id" db:"hand_id"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}