package services

import (
	"database/sql"
	"time"
)

var db *sql.DB

// time for db process with any transaction
const dbTimeout = time.Second * 3

type Models struct {
	Coffee Coffee
}