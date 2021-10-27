package entity

import "time"

type Order struct {
	ID        int
	Code      string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
