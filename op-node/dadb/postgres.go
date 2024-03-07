package dadb

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var Db *DB

type DB struct {
	db *sql.DB
}

func init() {
	Db = &DB{db: openDB()}
}

func openDB() *sql.DB {
	user := os.Getenv("DA_POSTGRES_USER")
	pass := os.Getenv("DA_POSTGRES_PASSWORD")
	host := os.Getenv("DA_POSTGRES_HOST")
	port := os.Getenv("DA_POSTGRES_PORT")
	dbName := os.Getenv("DA_POSTGRES_DB")
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Error(err.Error())
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	migrationsPath := os.Getenv("DA_POSTGRES_MIGRATIONS_PATH")
	if len(migrationsPath) > 0 {
		m, err := migrate.NewWithDatabaseInstance(
			migrationsPath,
			"postgres", driver)
		if err != nil {
			panic(err)
		}
		err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
		if err != nil {
			log.Error(err.Error())
		}
	}

	return db
}

// Inserts the given calldata into the db and returns the 8 byte long daTxId
func (db *DB) InsertCallData(data []byte) (int64, error) {
	sql := "INSERT INTO da_transactions(da_calldata) VALUES ($1) RETURNING da_tx_id"
	row := db.db.QueryRow(sql, data)
	if row.Err() != nil {
		return 0, fmt.Errorf("could not insert calldata into db: %s", row.Err().Error())
	}

	var daTxId int64
	err := row.Scan(&daTxId)
	if err != nil {
		return 0, err
	}
	return daTxId, nil
}

// Reads calldata from the db by a given daTxId
func (db *DB) ReadCallData(daTxId int64) ([]byte, error) {
	sql := "SELECT da_calldata FROM da_transactions WHERE da_tx_id = $1 LIMIT 1"
	row := db.db.QueryRow(sql, daTxId)
	if row.Err() != nil {
		return nil, fmt.Errorf("could not read calldata from db: %s", row.Err().Error())
	}

	var data []byte
	err := row.Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
