package mysql

import(
  	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
  db
}

func init() {
}