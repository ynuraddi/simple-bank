package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"simple-bank/util"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalln("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v\n", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
