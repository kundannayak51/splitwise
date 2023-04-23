package service

import (
	"splitwise/entity"
	"splitwise/repository"
	"splitwise/splitvalidation"
)

type ExpenseService struct {
	BalanceService BalanceService
	ExpenseRepo    repository.ExpenseRepo
}

func NewExpenseService(balanceService BalanceService, expenseRepo repository.ExpenseRepo) *ExpenseService {
	return &ExpenseService{
		BalanceService: balanceService,
		ExpenseRepo:    expenseRepo,
	}
}

func (es *ExpenseService) CreateExpense(expenseId string, desc string, expenseAmount float64, splitDetails []entity.Split,
	splitType entity.SPLITTYPE, paidBy entity.User) (*entity.Expense, error) {
	expenseSplit := splitvalidation.GetSplitObjectFromFactory(splitType)
	err := expenseSplit.ValidateSplitRequest(splitDetails, expenseAmount)
	if err != nil {
		return nil, err
	}

	expense := &entity.Expense{
		ExpenseId:     expenseId,
		Description:   desc,
		ExpenseAmount: expenseAmount,
		PaidBy:        paidBy,
		SplitType:     splitType,
		SplitDetails:  splitDetails,
	}
	err = es.ExpenseRepo.CreateExpense(*expense)
	if err != nil {
		return nil, err
	}

	return expense, nil
}
