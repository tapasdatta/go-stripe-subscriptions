package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string         `json:"name" gorm:"size:50;notNull"`
	Email         string         `json:"email" gorm:"size:50;notNull"`
	StripeId      sql.NullString `json:"stripe_id"`
	PmType        sql.NullString `json:"pm_type"`
	PmLastFour    sql.NullString `json:"pm_last_four" gorm:"size:4"`
	TrialEndsAt   sql.NullTime   `json:"trial_ends_at"`
	Subscriptions []Subscription
}
