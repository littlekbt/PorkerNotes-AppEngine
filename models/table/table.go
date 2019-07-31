package table

import ()

type Table struct {
	ID int64 `json:"id,int" db:"id, primarykey"`
}
