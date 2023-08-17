package db

import (
	"database/sql"
	"portfolio-api/src/config"

	_ "github.com/go-sql-driver/mysql" //Driver
)

func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.ConnString)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
