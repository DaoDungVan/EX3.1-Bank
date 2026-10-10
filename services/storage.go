package services

import (
	"Bank/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

const dataFile = "./data/users.json"

func LoadAccounts() ([]models.Account, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, fmt.Errorf(
				"data file does not exist: %s: %w", dataFile, err,
			)
		}
		if errors.Is(err, fs.ErrPermission) {
			return nil, fmt.Errorf(
				"cannot read data file %s: permission denied: %w",
				dataFile, err,
			)
		}
		return nil, fmt.Errorf(
			"cannot read data file %s: %w", dataFile, err,
		)
	}

	var accounts []models.Account
	if err := json.Unmarshal(data, &accounts); err != nil {
		return nil, fmt.Errorf(
			"invalid JSON in %s: %w", dataFile, err,
		)
	}

	return accounts, nil
}

func SaveAccounts(accounts []models.Account) error {
	data, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return fmt.Errorf(
			"cannot convert accounts to JSON: %w", err,
		)
	}

	if err := os.WriteFile(dataFile, data, 0600); err != nil {
		return fmt.Errorf(
			"cannot write accounts to file %s: %w",
			dataFile,
			err,
		)
	}

	return nil
}
