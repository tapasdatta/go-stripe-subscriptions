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
	//**Create stripe customer
	cID, err := user.CreateStripeCustomer()
	if err != nil {
		panic(err)
	}
	user.StripeId = *cID
	//**Save into DB
	models.DB.Create(user)
	utils.RespondWithJson(w, 201, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	userID := mux.Vars(r)["userID"]
	if err := models.DB.First(&user, userID).Error; err != nil{
		utils.RespondWithError(w, 404, "User not found")
		return
	}
	var input UserInput
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)
	validate = validator.New()
	err := validate.Struct(input)
	if err != nil {
		utils.RespondWithError(w, 420, "validation error")
	}
	models.DB.Model(&user).Clauses(clause.Returning{}).Updates(models.User{
		Name:       input.Name,
		Email:      input.Email,
	})
	//**Update stripe customer
	_, err = user.UpdateStripeCustomer(user.StripeId)
	if err != nil {
		panic(err)
	}
	utils.RespondWithJson(w, 201, user)
}
