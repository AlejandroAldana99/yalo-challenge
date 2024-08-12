package models

import "time"

type UserInteraction struct {
	UserID     string    `json:"user_id"`
	ProductSKU string    `json:"product_sku"`
	Action     string    `json:"action"`
	Timestamp  time.Time `json:"timestamp"`
	Duration   int       `json:"duration,omitempty"`
}
