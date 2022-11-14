package repsql

import "database/sql"

type SQLData struct {
	Connection *sql.DB
	Data       struct {
	}
}

func NewSQLData() *SQLData {
	return &SQLData{}
}
