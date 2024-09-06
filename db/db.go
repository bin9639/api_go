package db

import "github.com/go-sql-driver/mysql"

func NewSQLDB() (cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("error occured while connecting to db: %v", err)
	}

	return db, nil
}