package entity

import (
	"errors"
)

type Investment struct{
	left uint8
	right uint8
	amount float32
}

func NewInvestment(left uint8, right uint8, amount float32) (*Investment, error) {
	if left > right {
		return nil, errors.New("left must be less than right")
	}
	if amount < 0 {
		return nil, errors.New("amount must be greater than 0")
	}
	return &Investment{left, right, amount}, nil
}
