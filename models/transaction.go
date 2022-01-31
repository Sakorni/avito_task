package models

import (
	"time"
)

type Transaction struct {
	Id        uint      `json:"id"`
	From      uint      `json:"from_id" `
	To        uint      `json:"to_id"`
	Amount    int       `json:"amount" `
	CreatedAt time.Time `json:"created_at"`
}
