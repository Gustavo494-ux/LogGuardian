package database

import (
	"LogGuardian/src/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Conectar() (*sqlx.DB, error) {
	db, err := sqlx.Open(config.DriverBanco, config.StringConexao)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestarConexao() (err error) {
	_, err = sqlx.Open(config.DriverBanco, config.StringConexao)
	return
}
