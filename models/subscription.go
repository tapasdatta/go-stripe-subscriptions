package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	UserID            uint           `json:"user_id"`
	Name              string         `json:"name" gorm:"size:50;notNull"`
	StripeId          string         `json:"stripe_id" gorm:"notNull;unique"`
	StripeStatus      string         `json:"stripe_status" gorm:"notNull"`
	StripePrice       sql.NullString `json:"stripe_price"`
	Quantity          sql.NullInt32  `json:"quantity"`
	TrialEndsAt       sql.NullTime   `json:"trial_ends_at"`
	EndsAt            sql.NullTime   `json:"ends_at"`
	SubscriptionItems []SubscriptionItem
	User              User
}
