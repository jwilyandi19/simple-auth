package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jwilyandi19/simple-auth/handler"
	"github.com/jwilyandi19/simple-auth/middleware"
	repository "github.com/jwilyandi19/simple-auth/repository/user"
	"github.com/jwilyandi19/simple-auth/usecase/user"
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
	userUsecase := user.NewUserUsecase(userRepo)
	handler := handler.NewHandler(userUsecase)
	r := gin.New()
	r.POST("/auth/signup", handler.SignUpHandler)
	r.POST("/auth/login", handler.LoginHandler)

	route := r.Group("/", middleware.AuthorizationMiddleware)
	route.GET("/user/userlist", handler.UserListHandler)

	r.Run()
}
