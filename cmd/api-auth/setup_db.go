package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// DbSetup makes and ensures successful connection to the database and returrns
// database connection pool via func sql.Open(driverName string, dataSourceName string) (*sql.DB, error)
func DbSetup(ctx context.Context) (*sql.DB, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	var dbLogin = make(map[string]string, 3)

	for _, v := range [3]string{"HOSTNAME", "PASSWORD", "DBNAME"} {
		if val, ok := os.LookupEnv(v); ok && val != "" {
			dbLogin[v] = val
		}
	}

	var databaseURL string = fmt.Sprintf("postgres://%v:%v@localhost:5432/%v", dbLogin["HOSTNAME"], dbLogin["PASSWORD"], dbLogin["DBNAME"])

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	sqlFilePath, _ := filepath.Abs("./query.sql")

	sql, ioErr := os.ReadFile(sqlFilePath)
	if ioErr != nil {
		return nil, fmt.Errorf("unable to read .sql to setup database table: %v", err)
	}

	sqlQuery := string(sql)

	_, err = db.ExecContext(ctx, sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("unable to setup database table: %v", err)
	}

	return db, nil
}
