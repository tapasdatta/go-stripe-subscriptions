package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tapasdatta/go-stripe-subscriptions/models"
	"github.com/tapasdatta/go-stripe-subscriptions/utils"
)

type UserInput struct {
	Name  string `json:"name" validate:"required,gte=1,lte=50"`
	Email string `json:"email" validate:"required,email,gte=1,lte=50"`
}

var validate *validator.Validate

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var input UserInput
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
	validate = validator.New()
	err := validate.Struct(input)
	if err != nil {
		utils.RespondWithError(w, 420, "validation error")
	}
	user := &models.User{
		Name:  input.Name,
		Email: input.Email,
	}
	cID, err := user.CreateStripeCustomer()
	if err != nil {
		//handle error
	}
	user.StripeId = *cID
	models.DB.Create(user)
	utils.RespondWithJson(w, 201, user)
}
