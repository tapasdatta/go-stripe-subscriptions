package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tapasdatta/go-stripe-subscriptions/models"
	"github.com/tapasdatta/go-stripe-subscriptions/utils"
)

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
	sub, err := user.CreateSubscription(input.PriceID, input.PaymentToken, 7 * 24)
	if err != nil {
		utils.RespondWithError(w, 420, "Can't created subscription")
	}
	data := map[string]uint{"subscription_id": *sub}
	utils.RespondWithJson(w, 201, data)
}