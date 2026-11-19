package payment

import (
	"errors"
	"go-functions/internal/validation"
)

const feePercent = 2

func ProcessWithdrawal(amount, balance int64) (int64, int64, error) {
	err := validation.ValidateWithdrawal(amount)

	if err != nil {
		return 0, 0, err
	}
	fee  := amount * feePercent/100
	total := amount + fee

	if balance < total {
		return 0, 0,  errors.New("insufficient balance")
	}
	return fee, total, nil

}