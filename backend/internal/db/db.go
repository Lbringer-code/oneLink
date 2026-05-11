package db

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect( databaseUrlString string ) ( *sqlx.DB , error ) {
	db , err := sqlx.Connect( "pgx" , databaseUrlString )
	if err != nil {
		return nil , fmt.Errorf("connect to db: %w" , err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db , nil
}