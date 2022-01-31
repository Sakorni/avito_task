package models

import "github.com/golang/protobuf/ptypes/timestamp"

type BalanceInfo struct {
	Id        uint                `json:"id"`
	UserId    uint                `json:"user_id"`
	Balance   int                 `json:"balance"`
	CreatedAt timestamp.Timestamp `json:"created_at"`
	UpdatedAt timestamp.Timestamp `json:"updated_at"`
}
