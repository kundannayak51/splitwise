package entity

type Expense struct {
	ExpenseId     string
	Description   string
	ExpenseAmount float64
	PaidBy        User
	SplitType     SPLITTYPE
	SplitDetails  []Split
}
