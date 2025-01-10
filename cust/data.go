package cust

type Customer struct {
	Custid int
	Fname  string
	Lname  string
	Email  string
	Age    int
	Pnum   string
	Bal    int
}

var CustomerList []Customer
