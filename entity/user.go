package entity

type User struct {
	UserId           string                  `json:"userId"`
	Username         string                  `json:"username"`
	UserBalanceSheet UserExpenseBalanceSheet `json:"userBalanceSheet,omitempty"`
}
