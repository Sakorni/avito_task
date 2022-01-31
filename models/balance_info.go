package models

import (
	"time"
)

type BalanceInfo struct {
	Id        uint      `json:"id"`
	UserId    uint      `json:"user_id"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
