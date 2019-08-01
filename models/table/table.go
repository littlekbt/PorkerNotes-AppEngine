package table

import ()

type Table struct {
	ID   int64  `json:"id" db:"id, primarykey"`
	Name string `json:"name" db:"name"`
}
