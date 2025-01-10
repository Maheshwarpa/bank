package cust

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DB() {

	// Establish a connection to the MySQL database
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/sample")
	if err != nil {
		panic(err.Error()) // Exit if there is an error in the connection
	}

	// Ping the database to ensure the connection is successful
	if err := db.Ping(); err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to the MySQL database successfully!")

	ins, err := db.Exec(`INSERT INTO employee (emp_id, emp_name, email, emp_sal, emp_location) 
VALUES (6, "Maheshwar", "maheshwar@gmail.com", 98000, "Tampa")`)
	if err != nil {
		panic(err.Error()) // Handle the error appropriately
	}
	fmt.Println("Record inserted successfully.", ins)

	// Query to fetch all records from the 'employee' table
	sel, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err.Error()) // Exit if the query fails
	}
	defer sel.Close() // Ensure the query is closed after execution

	// Process the results
	fmt.Println("Employee Details:")
	//fmt.Println(sel)
	for sel.Next() {
		var empID int
		var empName, empMail, empLocation string
		var empSal int

		// Scan each row into variables
		err := sel.Scan(&empID, &empName, &empMail, &empSal, &empLocation)
		if err != nil {
			panic(err.Error())
		}

		// Print the retrieved data
		fmt.Printf("ID: %d, Name: %s, Email: %s, Salary: %d, Location: %s\n",
			empID, empName, empMail, empSal, empLocation)
	}

	// Close the database connection
	defer db.Close()

}
