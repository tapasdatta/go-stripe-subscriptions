package models

import (
	"database/sql"
	"time"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/subscription"
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	UserID            uint           `json:"user_id"`
	Name              string         `json:"name" gorm:"size:50;notNull"`
	StripeID          string         `json:"stripe_id" gorm:"notNull;unique"`
	StripeStatus      string         `json:"stripe_status" gorm:"notNull"`
	TrialEndsAt       time.Time   `json:"trial_ends_at"`
	EndsAt            sql.NullTime   `json:"ends_at"`
	SubscriptionItems []SubscriptionItem
	User              User
}


func (u *User) CreateSubscription(priceID string, paymentToken string, trialEnd int64) (*uint, error) {
	params := &stripe.SubscriptionParams{
		Customer: &u.StripeID,
		Items: []*stripe.SubscriptionItemsParams{
		  {
			Price: &priceID,
		  },
		},
		TrialEnd: &trialEnd,
		DefaultPaymentMethod: &paymentToken,
	}
	s, err := subscription.New(params)
	if err != nil {
		return nil, err
	}
	subscription := &Subscription{
		Model:             gorm.Model{},
		UserID:            u.ID,
		Name:              "default",
		StripeID:          s.ID,
		StripeStatus:     string(s.Status),
		TrialEndsAt:       time.Now().Add(time.Hour * time.Duration(trialEnd)),
	}
	DB.Create(subscription)
	return &subscription.ID, nil
}
