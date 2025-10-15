package models

import "time"

type Subscription struct {
	ID          string    `json:"id" db:"id"`
	UserID      string    `json:"user_id" db:"user_id"`
	ServiceName string    `json"service_name" db:"service_name"`
	Price       int       `json:"price" db:"price"`
	StartDate   time.Time `json:"start_date" db:"start_date""`
}

type SubscriptionFilter struct {
	FromDate    *time.Time
	ToDate      *time.Time
	UserID      *string
	ServiceName *string
}
