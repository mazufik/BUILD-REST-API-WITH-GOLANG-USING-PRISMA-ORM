package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mazufik/BUILD-REST-API-WITH-GOLANG-USING-PRISMA-ORM/config"
	"github.com/mazufik/BUILD-REST-API-WITH-GOLANG-USING-PRISMA-ORM/helper"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}
	fmt.Printf("Start Server " + os.Getenv("PORT") + "\n")

	// Handle DB Connection
	db, err := config.ConnectDB()
	helper.ErrorPanic(err)

	defer db.Prisma.Disconnect()

	server := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}
}
