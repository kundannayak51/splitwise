package entity

type Group struct {
	GroupId      string
	GroupName    string
	GroupMembers []User
	ExpenseList  []Expense
}
