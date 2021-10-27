package input

import "time"

type OrderInput struct {
	Code      string    `json:"code"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
