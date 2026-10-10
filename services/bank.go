package services

import (
	"Bank/models"
	"errors"
	"fmt"
	"strings"
)

// Homework 3: mỗi lần nạp tối đa 10000
const MaxDepositAmount = 10000

func UpdateAccount(updated models.Account) error {
	accounts, err := LoadAccounts()
	if err != nil {
		return err
	}

	for i := range accounts {
		if accounts[i].Username == updated.Username {
			accounts[i] = updated
			return SaveAccounts(accounts)
		}
	}

	return errors.New("account does not exist")
}

// Homework 2: tìm Account theo username.
func FindAccount(username string) (models.Account, error) {
	accounts, err := LoadAccounts()
	if err != nil {
		return models.Account{}, err
	}

	for i := range accounts {
		if accounts[i].Username == username {
			return accounts[i], nil
		}
	}

	return models.Account{}, fmt.Errorf("account %q does not exist", username)
}

func Deposit(acc *models.Account, amount float64) error {
	if acc == nil {
		return errors.New("account is nil")
	}
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	if amount > MaxDepositAmount {
		return fmt.Errorf(
			"deposit amount must not exceed %d per transaction",
			MaxDepositAmount,
		)
	}

	// cộng tiền trên bản copy trước, lưu file được rồi mới cập nhật acc
	// (không thì lưu lỗi mà số dư hiển thị vẫn tăng)
	updated := *acc
	updated.Balance += amount

	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Balance = updated.Balance
	return nil
}

func Withdraw(acc *models.Account, amount float64) error {
	if acc == nil {
		return errors.New("account is nil")
	}
	if amount <= 0 {
		return errors.New("withdraw amount must be greater than zero")
	}
	if amount > acc.Balance {
		return errors.New("account balance is not enough")
	}

	updated := *acc
	updated.Balance -= amount

	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Balance = updated.Balance
	return nil
}

func Transfer(from *models.Account, toUser string, amount float64) error {
	if from == nil {
		return errors.New("sender account is nil")
	}
	if amount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}
	if from.Username == toUser {
		return errors.New("sender and receiver must be different accounts")
	}

	accounts, err := LoadAccounts()
	if err != nil {
		return err
	}

	var sender *models.Account
	var receiver *models.Account

	for i := range accounts {
		if accounts[i].Username == from.Username {
			sender = &accounts[i]
		}
		if accounts[i].Username == toUser {
			receiver = &accounts[i]
		}
	}

	if sender == nil || receiver == nil {
		return errors.New("sender or receiver account does not exist")
	}
	if sender.Balance < amount {
		return errors.New("account balance is not enough")
	}

	sender.Balance -= amount
	receiver.Balance += amount

	if err := SaveAccounts(accounts); err != nil {
		return err
	}

	// from là account lúc Login, sender là account trong slice mới load
	// nên lưu xong phải gán lại số dư cho from
	from.Balance = sender.Balance
	return nil
}

// Homework 1: đổi mật khẩu, lưu qua UpdateAccount.
func ChangePassword(acc *models.Account, newPassword string) error {
	if acc == nil {
		return errors.New("account is nil")
	}
	if strings.TrimSpace(newPassword) == "" {
		return errors.New("new password must not be empty")
	}

	updated := *acc
	updated.Password = newPassword

	if err := UpdateAccount(updated); err != nil {
		return err
	}

	acc.Password = updated.Password
	return nil
}
