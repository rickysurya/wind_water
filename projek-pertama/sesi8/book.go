type Book struct {
	BookID int    `json:"id" db:"bookid"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
	Desc   string `json:"desc" db:"description"`
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