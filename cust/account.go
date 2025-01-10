package cust

import "fmt"

func CheckAcc(cid int) bool {
	flag := false

	for _, cus := range CustomerList {
		if cus.Custid == cid {
			flag = true
			return flag
		}
	}

	return false
}

func DelAct(cid int) {
	flag := false
	for ind, cus := range CustomerList {
		if cus.Custid == cid {
			flag = true
			CustomerList = append(CustomerList[:ind], CustomerList[ind+1:]...)
		}
	}
	if !flag {
		fmt.Println("Account Id not found")
	}
}

func DepositAmt(cid int, amt int) {
	flag := false
	for ind, cus := range CustomerList {
		if cus.Custid == cid {
			flag = true
			CustomerList[ind].Bal = cus.Bal + amt
			fmt.Println("Amount Deposit from Customer"+cus.Fname+" "+cus.Lname+" is :", amt)
			CustReport(cus.Custid)
		}
	}
	if !flag {
		fmt.Println("Account Id not found for Deposit")
	}
}

func WithdrawlAct(cid int, amt int) {
	flag := false
	for ind, cus := range CustomerList {
		if cus.Custid == cid {
			flag = true
			if !LowBalChk(cid) {
				if amt > cus.Bal {
					fmt.Println("Insufficient Balance, please try later")
				} else {
					CustomerList[ind].Bal = cus.Bal - amt
					fmt.Println("Amount Withdrawn from Customer"+cus.Fname+" "+cus.Lname+" is :", amt)
					CustReport(cus.Custid)
				}
			} else {
				fmt.Println("You dont have sufficient balance to perform withdraw Operation")
			}
		}
	}
	if !flag {
		fmt.Println("Account Id not found for withdrawl")
	}
}

func CustReport(cid int) {
	fmt.Println("*****   Customer Report     *****")
	for _, cus := range CustomerList {
		if cus.Custid == cid {
			fmt.Printf("Cust Id:\t%d\n", cus.Custid)
			fmt.Printf("Name:\t%s %s\n", cus.Fname, cus.Lname)
			fmt.Printf("Email:\t%s\n", cus.Email)
			fmt.Printf("Balance:\t%d\n", cus.Bal)

		}
	}
	fmt.Println("*********************************")
}

func LowBalChk(cid int) bool {
	flag := false
	for _, cus := range CustomerList {
		if cus.Custid == cid {
			if cus.Bal < 1500 {
				flag = true
				return flag
			}
		}
	}
	return flag
}
