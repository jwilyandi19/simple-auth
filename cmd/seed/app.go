package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jwilyandi19/simple-auth/db/seed"
	repository "github.com/jwilyandi19/simple-auth/repository/user"
	"github.com/jwilyandi19/simple-auth/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	dbConn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect db: ", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping db: ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("failed to close db: ", err)
		}
	}()

	userRepo := repository.NewMySqlUserRepository(dbConn)
	seeder := seed.NewSeeder(userRepo)

	seeder.Seed()
}
