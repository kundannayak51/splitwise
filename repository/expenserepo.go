package repository

import "splitwise/entity"

type ExpenseRepo struct {
	Expenses map[string]entity.Expense
}

func NewExpenseRepo(expenses map[string]entity.Expense) *ExpenseRepo {
	return &ExpenseRepo{
		Expenses: expenses,
	}
}

func (er *ExpenseRepo) CreateExpense(expense entity.Expense) error {
	er.Expenses[expense.ExpenseId] = expense
	return nil
}
