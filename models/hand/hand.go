package hand

import ()

type Hand struct {
	ID   int64  `json:"id" db:"id, primarykey"`
	Name string `json:"name" db:"name"`
}
