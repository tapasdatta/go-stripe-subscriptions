package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/stripe/stripe-go/v74"
	"github.com/tapasdatta/go-stripe-subscriptions/models"
	"github.com/tapasdatta/go-stripe-subscriptions/utils"
	"gorm.io/gorm/clause"
)

type SubscriptionInput struct {
	CustomerID  string `json:"customer_id" validate:"required"`
	PriceID string `json:"price_id" validate:"required"`
	PaymentToken string `json:"payment_token" validate:"required"`
}

type SubscriptionCancelInput struct {
	ID string `json:"id" validate:"required"`
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
	sub, err := user.CreateSubscription(input.PriceID, input.PaymentToken, 7 * 24)
	if err != nil {
		utils.RespondWithError(w, 420, "Can't created subscription")
	}
	data := map[string]uint{"subscription_id": *sub}
	utils.RespondWithJson(w, 201, data)
}

func CancelSubscription(w http.ResponseWriter, r *http.Request) {
	var input SubscriptionCancelInput
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
	validate = validator.New()
	err := validate.Struct(input)
	if err != nil {
		utils.RespondWithError(w, 420, "validation error")
	}
	var subscription models.Subscription
	if err := models.DB.First(&subscription, input.ID).Error; err != nil{
		utils.RespondWithError(w, 404, "Subscription not found")
		return
	}
	if err = subscription.CancelSubscription(); err != nil {
		utils.RespondWithError(w, 420, "Can't created subscription")
	}
	models.DB.Model(&subscription).Clauses(clause.Returning{}).Updates(models.Subscription{
		StripeStatus: string(stripe.SubscriptionStatusCanceled),
	})
	utils.RespondWithJson(w, 201, subscription)
	
}