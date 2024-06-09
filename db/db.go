package db

import (
	"log"

	"github.com/euklid-dev/deep_chained/internal/config"
	"github.com/jmoiron/sqlx"
)

var (
	SQLx *sqlx.DB
)

func ConnectToDatabase() {
	var err error
	SQLx, err = sqlx.Connect("postgres", config.GlobalAppConfig.DB_URL)

	if err != nil {
		log.Fatal(err)
	}

	err = SQLx.Ping()

	if err != nil {
		log.Fatalln("Error: Could not establish a connection with the database", err)
	}

}
