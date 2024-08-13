package models

import "time"

//go:generate go-enum --marshal --
// ENUM(view, click, add_to_cart)
type ActionType int

type UserInteraction struct {
	UserID     string     `json:"user_id"`
	ProductSKU string     `json:"product_sku"`
	Action     ActionType `json:"action"`
	Timestamp  time.Time  `json:"timestamp"`
	Duration   int        `json:"duration,omitempty"`
}
