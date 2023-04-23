package entity

type UserExpenseBalanceSheet struct {
	UserVsBalance    map[string]Balance
	TotalYourExpense float64
	TotalPayment     float64
	TotalYouOwe      float64
	TotalYouGetBack  float64
}
