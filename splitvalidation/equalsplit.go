package splitvalidation

import (
	"errors"
	"splitwise/entity"
)

type EqualSplit struct{}

func NewEqualSplit() *EqualSplit {
	return &EqualSplit{}
}

func (es *EqualSplit) ValidateSplitRequest(splitDetails []entity.Split, expenseAmount float64) error {
	amountShouldBePresent := expenseAmount / (float64)(len(splitDetails))
	for _, split := range splitDetails {
		if split.Amount != amountShouldBePresent {
			return errors.New("Not equal split")
		}
	}
	return nil

}
