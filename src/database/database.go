package database

import (
	"LogGuardian/src/config"

	"github.com/jmoiron/sqlx"
	// Drive para SQLite GCC free
	_ "modernc.org/sqlite"
)

func Conectar() (*sqlx.DB, error) {
	db, err := sqlx.Open(config.DriverBanco, config.StringConexao)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestarConexao() (err error) {
	_, err = Conectar()
	return
}
