package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "admin"
	dbname = "db-go-sql"
)

var (
	db *sql.DB
	err error
)

func main(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
	CreateEmployee()
}

type Employee struct {
	ID 			int		`json:"id" db:"id`
	FullName 	string	`json:"full_name" db:"full_name"`
	Email 		string	`json:"email" db:"email"`
	Age 		int		`json:"age" db:"age"`
	Division 	string	`json:"division" db:"division"`
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	Returning *
	`

	err = db.QueryRow(sqlStatement, "Budi", "budi@mail.com", 23, "IT").
		Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func Getemployee() {
	var results = []Employee{}

	sqlStatement := `SELECT full_name, email, age, division from employee`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)

	}

	fmt.Println("Employee datas:", results)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employee
	SET full_name = $2, email = $3, division = $4, age = $5
	WHERE id = $1;`
	res, err := db.Exec(sqlStatement, 12, "Airell Jordan Hidayat", "airellhidayat@gmail.com", "CurDevs", 24)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Updated data amount:", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE from employee
	WHERE id = $1;
	`
	res, err := db.Exec(sqlStatement, 12)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted data amount:", count)
}