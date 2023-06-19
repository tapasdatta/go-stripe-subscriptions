package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/tapasdatta/go-stripe-subscriptions/models"
	"github.com/tapasdatta/go-stripe-subscriptions/utils"
	"gorm.io/gorm/clause"
)

// var validate *validator.Validate

type SubscriptionInput struct {
	CustomerID  string `json:"customer_id" validate:"required"`
	PriceID string `json:"price_id" validate:"required"`
	PaymentToken string `json:"payment_token" validate:"required"`
}

func CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var input SubscriptionInput
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
	validate = validator.New()
	err := validate.Struct(input)
	if err != nil {
		utils.RespondWithError(w, 420, "validation error")
	}
	var user models.User
	if err := models.DB.First(&user, input.CustomerID).Error; err != nil{
		utils.RespondWithError(w, 404, "User not found")
		return
	}
	//**create subscription
	params := &stripe.SubscriptionParams{
		Customer: stripe.String("cus_9s6XFG2Qq6Fe7v"),
		Items: []*stripe.SubscriptionItemsParams{
		  {
			Price: stripe.String("price_1NKf0l2eZvKYlo2CCjzb0KEW"),
		  },
		},
	  }
	  s, err := sub.New(params)
	subscriptionParams := &stripe.SubscriptionParams{
        Customer: &customerID,
        Items: []*stripe.SubscriptionItemsParams{
            {
                Plan: &priceID,
            },
        },
        TrialEnd:             &trialEnd,
        DefaultPaymentMethod: &paymentMethodID,
    }
    sb, err := sub.New(subscriptionParams)}

	utils.RespondWithJson(w, 201, user)
}