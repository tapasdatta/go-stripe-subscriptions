package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type SubscriptionItem struct {
	gorm.Model
	SubscriptionID uint          `json:"subscription_id"`
	StripeID       string        `json:"stripe_id" gorm:"notNull;unique"`
	StripeProduct  string        `json:"stripe_product" gorm:"notNull"`
	StripePrice    string        `json:"stripe_price" gorm:"notNull"`
	Quantity       sql.NullInt32 `json:"quantity"`
	Subscription   Subscription
}
