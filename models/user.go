package models

import (
	"database/sql"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/setupintent"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string         `json:"name" gorm:"size:50;notNull"`
	Email         string         `json:"email" gorm:"size:50;notNull"`
	StripeID      string         `json:"stripe_id"`
	PmType        sql.NullString `json:"pm_type"`
	PmLastFour    sql.NullString `json:"pm_last_four" gorm:"size:4"`
	TrialEndsAt   sql.NullTime   `json:"trial_ends_at"`
	Subscriptions []Subscription
}

func (u *User) CreateStripeCustomer() (*string, error) {
	params := &stripe.CustomerParams{
		Name:             stripe.String(u.Name),
		Email:            stripe.String(u.Email),
		PreferredLocales: stripe.StringSlice([]string{"en", "es"}),
	}
	c, err := customer.New(params)
	if err != nil { 
		return nil, err
	}
	return &c.ID, nil
}

func (u *User) UpdateStripeCustomer(stripeID string) (*string, error) {
	params := &stripe.CustomerParams{
		Name:             stripe.String(u.Name),
		Email:            stripe.String(u.Email),
		PreferredLocales: stripe.StringSlice([]string{"en", "es"}),
	}
	c, err := customer.Update(stripeID, params)
	if err != nil { 
		return nil, err
	}
	return &c.ID, nil
}

func (u *User) CreateSetupIntent() (*string, error) {
	params := &stripe.SetupIntentParams{
		PaymentMethodTypes: []*string{
		  stripe.String("card"),
		},
		Customer: &u.StripeID,
	}
	si, err := setupintent.New(params)
	if err != nil {
		return nil, err
	}
	return &si.ClientSecret, nil
}