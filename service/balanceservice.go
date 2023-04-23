package service

import (
	"fmt"
	"splitwise/entity"
	"splitwise/repository"
)

type BalanceService struct {
	UserRepo repository.UserRepo
}

func NewBalanceService(userRepo repository.UserRepo) *BalanceService {
	return &BalanceService{
		UserRepo: userRepo,
	}
}

func (bs *BalanceService) UpdateUserExpenseBalanceSheet(paidBy entity.User, splits []entity.Split, totalExpenseAmount float64) {
	//Updating totatl amount paid of the expense paid by user
	paidByUserExpenseSheet := paidBy.UserBalanceSheet
	paidByUserExpenseSheet.TotalPayment += totalExpenseAmount

	for _, split := range splits {
		userOwe := split.User
		oweUserBalanceSheet := userOwe.UserBalanceSheet
		oweAmount := split.Amount

		if userOwe.UserId == paidBy.UserId {
			paidByUserExpenseSheet.TotalYourExpense += oweAmount
		} else {
			paidByUserExpenseSheet.TotalYouGetBack += oweAmount

			var userOweBalance entity.Balance
			if _, exists := paidByUserExpenseSheet.UserVsBalance[userOwe.UserId]; exists {
				userOweBalance = paidByUserExpenseSheet.UserVsBalance[userOwe.UserId]
			} else {
				paidByUserExpenseSheet.UserVsBalance[userOwe.UserId] = entity.Balance{}
			}

			userOweBalance.AmountGetBack += oweAmount

			//Update Balance sheet of owe user
			oweUserBalanceSheet.TotalYouOwe += oweAmount
			oweUserBalanceSheet.TotalYourExpense += oweAmount

			var userPaidBalance entity.Balance
			if _, exists := oweUserBalanceSheet.UserVsBalance[paidBy.UserId]; exists {
				userPaidBalance = oweUserBalanceSheet.UserVsBalance[paidBy.UserId]
			} else {
				oweUserBalanceSheet.UserVsBalance[paidBy.UserId] = entity.Balance{}
			}
			userPaidBalance.AmountOwe += oweAmount

			oweUserBalanceSheet.UserVsBalance[paidBy.UserId] = userPaidBalance
			userOwe.UserBalanceSheet = oweUserBalanceSheet

			bs.UserRepo.UpdateUser(userOwe)

		}
		paidBy.UserBalanceSheet = paidByUserExpenseSheet
		bs.UserRepo.UpdateUser(paidBy)
	}
}

func (bs *BalanceService) ShowBalanceSheetOfUser(u entity.User) {
	user, err := bs.UserRepo.GetUser(u.UserId)
	if err != nil {
		return
	}

	fmt.Println("Balance Sheet of User: " + user.Username)

	balanceSheet := user.UserBalanceSheet
	fmt.Println(fmt.Sprintf("TotalYourExpense: %s", balanceSheet.TotalYourExpense))
	fmt.Println(fmt.Sprintf("TotalGetBack: %s", balanceSheet.TotalYouGetBack))
	fmt.Println(fmt.Sprintf("TotalYouOwe: %s", balanceSheet.TotalYouOwe))
	fmt.Println(fmt.Sprintf("TotalPaymentMade: %s", balanceSheet.TotalPayment))

	for key, val := range balanceSheet.UserVsBalance {
		userId := key
		balance := val

		fmt.Println(fmt.Sprintf("UserId: %s, YouGetBack: %s, YouOwe: %s", userId, balance.AmountGetBack, balance.AmountOwe))
	}

	fmt.Println("___________________________________________________")
}
