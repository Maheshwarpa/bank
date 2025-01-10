package cust

import (
	"fmt"
	//"./cust"
)

/*
type Customer struct {
	Custid   int
	Fname    string
	Lname    string
	Email    string
	Age      int
	Pnum     string
	Bal			int
}
*/

func AddCustomer() {
	var cust Customer
	var cid int
	ag, bl := 0, 0
	fn, ln, em, pn := "", "", "", ""
	fmt.Println("Please enter the requested details of the Customer:")
	fmt.Println("Please enter the Customer Id:")
	fmt.Scan(&cid)
	fmt.Println(cid)
	if !CheckAcc(cid) {
		fmt.Println("Please enter the Customer Fname:")
		fmt.Scan(&fn)
		fmt.Println("Please enter the Customer Lname:")
		fmt.Scan(&ln)
		fmt.Println("Please enter the Customer Email:")
		fmt.Scan(&em)
		fmt.Println("Please enter the Customer Age:")
		fmt.Scan(&ag)
		fmt.Println("Please enter the Customer Pnum:")
		fmt.Scan(&pn)
		fmt.Println("Please enter the Customer Bal:")
		fmt.Scan(&bl)

		cust = Customer{Custid: cid, Fname: fn, Lname: ln, Email: em, Age: ag, Pnum: pn, Bal: bl}
		CustomerList = append(CustomerList, cust)
	} else {
		fmt.Println("Already Customer have an account")
	}

}

func CustCount() {
	fmt.Println("Please find the number of customers present in the list :")
	fmt.Println(len(CustomerList))
}

func DisplayCust() {
	fmt.Println("Please find the list of customers :")
	for _, cus := range CustomerList {
		fmt.Println(cus)
	}
}
