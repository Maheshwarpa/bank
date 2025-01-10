package main

import (

	//"fmt"

	"fmt"
	"king/cust"
	//_ "github.com/go-sql-driver/mysql"
)

// func init() {
// 	sql.Register("mysql", &MySQLDriver{})
// }

// func main() {
// 	// Create Database
// 	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/sample")

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	db.Ping()
// 	// defer the close till after the main function has finished
// 	// executing

// 	// err = db.Ping()
// 	sel, err := db.Query("select * from employee")

// 	// if there is an error inserting, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// be careful deferring Queries if you are using transactions
// 	defer sel.Close()

// 	// fmt.Println("Please enter the number of Customers you want to add: ")
// 	// var cnt int
// 	// fmt.Scanln(&cnt)
// 	// var fn, ln, em, pn string
// 	// var status bool
// 	// var age, id int
// 	// var cl []cust.Customer
// 	// for i := 0; i < cnt; i++ {
// 	// 	fmt.Scanln(&id)
// 	// 	fmt.Scanln(&fn)
// 	// 	fmt.Scanln(&ln)
// 	// 	fmt.Scanln(&em)
// 	// 	fmt.Scanln(&pn)
// 	// 	fmt.Scanln(&age)
// 	// 	fmt.Scanln(&status)
// 	// 	cst := cust.Customer{Custid: id, Fname: fn, Lname: ln, Email: em, Age: age, Pnum: pn, M_status: status}
// 	// 	cl = append(cl, cst)
// 	// 	cust.CustomerList = append(cust.CustomerList, cst)

// 	// }
// 	// fmt.Println(len(cust.CustomerList))
// 	// fmt.Println(cust.CustomerList)

// 	// fmt.Println(len(cl))
// 	// fmt.Println(cl)
// 	defer db.Close()
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	_ "github.com/go-sql-driver/mysql" // Import MySQL driver
// )

func main() {

	for flag := false; flag != true; {
		fmt.Println("Please check what operation you want to perform :\n1. Add\n2. Display Customers\n3. Customer Count\n4. Deposit\n5. Withdraw\n6. Close Account\nPlease enter any other option to exit")
		k := 0
		fmt.Scan(&k)
		switch k {
		case 1:
			cust.AddCustomer()
			fmt.Println("------------")

		case 2:
			cust.DisplayCust()
			fmt.Println("------------")

		case 3:
			cust.CustCount()
			fmt.Println("------------")
		case 4:
			cid := 0
			amt := 0
			fmt.Println("Please enter the customer id of the customer :")
			fmt.Scan(&cid)
			fmt.Println("Please enter the Amount to Deposit :")
			fmt.Scan(&amt)
			cust.DepositAmt(cid, amt)
			fmt.Println("------------")
		case 5:
			cid := 0
			amt := 0
			fmt.Println("Please enter the customer id of the customer :")
			fmt.Scan(&cid)
			fmt.Println("Please enter the Amount to Withdraw :")
			fmt.Scan(&amt)
			cust.WithdrawlAct(cid, amt)
			fmt.Println("------------")
		case 6:
			cid := 0
			fmt.Println("Please enter the customer id of the customer :")
			fmt.Scan(&cid)
			cust.DelAct(cid)
			fmt.Println("------------")
		default:
			flag = true
		}
	}
	// fmt.Println("*****Cust count Check********")
	// cust.CustCount()
	// fmt.Println("***Display**********")
	// cust.DisplayCust()
	// fmt.Println("*****Cust Dup Check********")
	// cust.AddCustomer()
	// fmt.Println("*****Deposit********")
	// cust.DepositAmt(1, 10000)
	// fmt.Println("****Display*********")
	// cust.DisplayCust()
	// fmt.Println("***Withdrawl**********")
	// cust.WithdrawlAct(1, 5000)
	// fmt.Println("***Display**********")
	// cust.DisplayCust()
	// fmt.Println("***Delete**********")
	// cust.DelAct(5)
	// fmt.Println("*****Display********")
	// cust.DisplayCust()
}
