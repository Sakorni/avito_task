package models

import "github.com/golang/protobuf/ptypes/timestamp"

type Transaction struct {
	Id        uint                `json:"id"`
	From      uint                `json:"from_id"`
	To        uint                `json:"to_id"`
	Amount    int                 `json:"amount"`
	CreatedAt timestamp.Timestamp `json:"created_at"`
}
