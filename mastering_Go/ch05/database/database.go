package main

import (
	"database/sql"
	"fmt"
	"time"

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

	conn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, pass, database)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("Open():", err)
		return
	}
	db.SetConnMaxIdleTime(3 * time.Second) // timeout 세팅 5분 이하로 하도록 하자.
	db.SetMaxOpenConns(10)                 // 애플리케이션으로 인한 연결 제한 설정
	db.SetMaxIdleConns(10)                 // SetMaxOpenConns보다 크거나 같게 설정하도록 한다.

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

	rows, err = db.Query(`select * from article`)
	if err != nil {
		fmt.Println("query2", err)
		return
	}

	for rows.Next() {
		var id, title, content, c_at, u_at string
		err = rows.Scan(&id, &title, &content, &c_at, &u_at)
		if err != nil {
			fmt.Println("row2", err)
			return
		}
		fmt.Println("article data : ", id, title, content, c_at, u_at)
	}
	defer rows.Close()
}
