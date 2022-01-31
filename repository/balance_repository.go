package repository

import (
	"avito_task/errors"
	"avito_task/models"
	"database/sql"
)

type BalanceRepository struct {
	db *sql.DB
}

func NewBalanceRepository(db *sql.DB) *BalanceRepository {
	return &BalanceRepository{db: db}
}

func (b *BalanceRepository) GetBalance(userId uint) (info *models.BalanceInfo, err error) {
	tx, err := b.db.Begin()
	if err != nil {
		return
	}
	return b.getBalance(tx, userId)
}

func (b *BalanceRepository) ChangeAmount(id uint, amount int) (info *models.BalanceInfo, err error) {
	tx, err := b.db.Begin()
	defer tx.Rollback()

	info, err = b.getBalance(tx, id)
	if err != nil {
		if err != errors.NoSuchUser || amount < 0 {
			return
		}
		info, err = b.createBalance(tx, id)
		if err != nil {
			return
		}
	}

	info.Balance += amount
	if info.Balance < 0 {
		return nil, errors.NotEnoughMoney
	}

	err = b.createTransaction(tx, &models.Transaction{
		From:   id,
		To:     id,
		Amount: amount,
	})

	if err != nil {
		return
	}

	err = b.updateBalance(tx, info)
	if err != nil {
		return
	}

	info, err = b.getBalance(tx, id)
	if err != nil {
		return
	}

	err = tx.Commit()
	return
}

func (b *BalanceRepository) Transfer(transaction *models.Transaction) (err error) {
	tx, err := b.db.Begin()
	defer tx.Rollback()
	//Checking if there is invalid from id
	balanceFrom, err := b.getBalance(tx, transaction.From)
	if balanceFrom.Balance < transaction.Amount {
		return errors.NotEnoughMoney
	}

	if err != nil {
		return err
	}

	//Check if to id exists and creating if it's not
	balanceTo, err := b.getBalance(tx, transaction.To)
	if err != nil {
		if err != errors.NoSuchUser {
			return err
		}
		balanceTo, err = b.createBalance(tx, transaction.To)
		if err != nil {
			return err
		}
	}

	err = b.createTransaction(tx, transaction)
	if err != nil {
		return err
	}

	balanceTo.Balance += transaction.Amount
	balanceFrom.Balance -= transaction.Amount
	err = b.updateBalance(tx, balanceFrom)
	if err != nil {
		return err
	}
	err = b.updateBalance(tx, balanceTo)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (b *BalanceRepository) createBalance(tx *sql.Tx, id uint) (*models.BalanceInfo, error) {
	_, err := tx.Exec("INSERT INTO balance(user_id, amount) values(?,?)", id, 0)
	if err != nil {
		return nil, err
	}
	res, err := b.getBalance(tx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (b *BalanceRepository) getBalance(tx *sql.Tx, id uint) (*models.BalanceInfo, error) {
	query, err := tx.Query("Select * from balance where User_id = ?", id)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	res := &models.BalanceInfo{}
	if !query.Next() {
		return nil, errors.NoSuchUser
	}
	err = query.Scan(&res.Id, &res.UserId, &res.Balance, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BalanceRepository) createTransaction(tx *sql.Tx, transaction *models.Transaction) error {
	_, err := tx.Exec("INSERT INTO transaction(From_id, To_id, Amount) values(?, ?, ?)", transaction.From, transaction.To, transaction.Amount)
	return err
}

func (b *BalanceRepository) updateBalance(tx *sql.Tx, info *models.BalanceInfo) error {
	_, err := tx.Exec("UPDATE balance SET Amount = ? where Id = ?", info.Balance, info.Id)
	return err
}
