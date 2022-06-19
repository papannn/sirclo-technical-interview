package server

import (
	"database/sql"
	"net/http"

	"github.com/backend/domain/weight"
	ucweight "github.com/backend/usecase/weight"
	_ "github.com/mattn/go-sqlite3"
)

var (
	weightDomain weight.Domain

	WeightUsecase ucweight.Usecase
)

const (
	Query = `
	CREATE TABLE weight_data (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date string NOT NULL,
		min int NOT NULL,
		max int NOT NULL
	 );
	`
)

func InitDB() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", "./files/database.db")
	if err != nil {
		return db, err
	}

	// Create Table
	db.Exec(Query)
	return db, nil
}

func InitDomain(db *sql.DB) {
	weightDomain = *weight.NewDomain(db)
}

func InitUsecase() {
	WeightUsecase = *ucweight.NewUsecase(&weightDomain)
}

func Init() (handler *http.ServeMux, err error) {
	db, err := InitDB()
	if err != nil {
		return handler, err
	}

	InitDomain(db)
	InitUsecase()

	handler = http.NewServeMux()

	initRoute(handler)

	return handler, nil
}

func initRoute(handler *http.ServeMux) {
	handler.HandleFunc("/", HandleIndexWeightData)
	handler.HandleFunc("/create", HandleCreateWeightData)
	handler.HandleFunc("/edit", HandleEditWeightData)
	handler.HandleFunc("/reset", HandleReset)
}
