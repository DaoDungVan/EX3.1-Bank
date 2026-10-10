package services

import (
	"Bank/models"
	"errors"
)

func Login(username, password string) (*models.Account, error) {
	accounts, err := LoadAccounts()
	if err != nil {
		return nil, err
	}

	for i := range accounts {
		if accounts[i].Username == username &&
			accounts[i].Password == password {
			return &accounts[i], nil
		}
	}

	return nil, errors.New("invalid username or password")
}
