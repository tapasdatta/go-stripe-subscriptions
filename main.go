package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/tapasdatta/go-stripe-subscriptions/models"
)

func main() {

	godotenv.Load()

	models.ConnectDatabase()

	fmt.Println("Database migration done!")
}
