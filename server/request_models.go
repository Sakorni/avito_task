package server

type ChangeAmountModel struct {
	UserId uint `json:"user_id" binding:"required"`
	Amount int  `json:"amount" binding:"required"`
}
type TransactionModel struct {
	FromId uint `json:"from_id" binding:"required"`
	ToId   uint `json:"to_id" binding:"required"`
	Amount uint `json:"amount" binding:"required"`
}
