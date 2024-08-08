package main

import (
	"chat_application/db"
	"chat_application/router"
	"chat_application/user"
	"log"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialise database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRep)
	userHandler := user.NewHandler(userService)

	router.InitRouter(userHandler)
	router.Start("0.0.0.0:8080")
}
