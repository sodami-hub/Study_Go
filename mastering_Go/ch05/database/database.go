package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// arguments := os.Args
	// if len(arguments) != 6 {
	// 	fmt.Println("Please provide: hostname port username password db")
	// 	return
	// }

	// host := arguments[1]	// localhost
	// port := arguments[2]		// 3306
	// user := arguments[3]	// root
	// pass := arguments[4]	// admin
	// database := arguments[5]	// test_db

	// // 포트 번호는 무조건 정수여야 한다.
	// port, err := strconv.Atoi(port)
	// if err != nil {
	// 	fmt.Println("Not a valid port number", err)
	// 	return
	// }

	// host := "localhost"
	// port := 3306
	user := "root"
	pass := "admin"
	database := "test_db"

	conn := fmt.Sprintf("%s:%s@/%s", user, pass, database)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("Open():", err)
		return
	}
	defer db.Close()
	rows, err := db.Query(`show tables`)
	if err != nil {
		fmt.Println("Query :", err)
		return
	}

	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("scan", err)
			return
		}
		fmt.Println("*", name)
	}
	defer rows.Close()
}
