package services

import (
	"database/sql"
	"time"
)

var db *sql.DB

// time for db process with any transaction
const dbTimeout = time.Second * 3

func New(dbPool *sql.DB) Models {
	db = dbPool
	
	return Models {}
}

type Models struct {
	Coffee Coffee
	JsonResponse JsonResponse
}