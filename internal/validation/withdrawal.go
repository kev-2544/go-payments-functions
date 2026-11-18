package validation

import (
	"errors"
)

func ValidateWithdrawal(amount int64) error {
	if amount <= 0 {
		return errors.New("Amount must be positive")
	}
	if amount < 500 {
		return errors.New("minimum withdrawal is 500")
	}
	if amount > 500000 {
		return errors.New("maximum withdrawal is 500000")
	}
	return nil
}