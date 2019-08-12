package table

import (
  // "log"
  "strconv"
  "time"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"
  "github.com/littlekbt/PorkerNotes-AppEngine/models/porker_notes"
)

type Table struct {
	ID   int64  `json:"id" db:"id, primarykey"`
  Name string `json:"name" db:"name"`
  Memo string `json:"memo" db:"memo"`
  Type porker_notes.Table_Type `json:"type" db:"type"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
  UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func Select(offset int) ([]Table, error) {
  limit := 10
  tables := []Table{}
  db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3307)/porker_notes?parseTime=true&loc=Asia%2FTokyo")
  defer db.Close()
	if err != nil {
		return tables, err
	}

  rows, err := db.Query("SELECT id, name, type, memo, created_at, updated_at FROM tables limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset))
	if err != nil {
		return tables, err
  }
  for rows.Next() {
    t := Table{}
    if err := rows.Scan(&t.ID,&t.Name, &t.Type, &t.Memo, &t.CreatedAt, &t.UpdatedAt); err != nil {
      return tables, err
    }
    tables = append(tables, t)
  }
  return tables, nil
}

func (t Table) Insert() (Table, error){
  db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3307)/porker_notes?parseTime=true&loc=Asia%2FTokyo")
  defer db.Close()
	if err != nil {
		return t, err
	}

	ins, err := db.Prepare("INSERT INTO tables(name, type, memo, created_at, updated_at) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return t, err
	}
	ins.Exec(t.Name, t.Type, t.Memo, t.CreatedAt, t.UpdatedAt)
	return t, nil
}

func (t Table) Valid() bool {
  if t.Name == "" || t.Type == 0 {
    return false
  }
  return true
}
