package models

type Recommendation struct {
	UserID   string   `json:"user_id"`
	Products []string `json:"products"`
}
