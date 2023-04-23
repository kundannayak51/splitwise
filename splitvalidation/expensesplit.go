package splitvalidation

import "splitwise/entity"

type ExpenseSplit interface {
	ValidateSplitRequest(splitDetails []entity.Split, expenseAmount float64) error
}
