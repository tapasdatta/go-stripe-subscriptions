package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v74"
	"github.com/tapasdatta/go-stripe-subscriptions/models"
)

func main() {

	godotenv.Load()

	stripe.Key = os.Getenv("STRIPE_KEY")

	models.ConnectDatabase()

	fmt.Println("Database migration done!")
}
