package mypackage

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// 사용할 DB의 두개의 테이블의 데이터를 하나로 합쳤다.
type Userdata struct {
	ID         int
	Username   string
	Name       string
	Surname    string
	Descrption string
}

var (
	Hostname = ""
	Port     = 3036
	Username = ""
	Password = ""
	Database = ""
)

func OpenConnection() (*sql.DB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", Username, Password, Hostname, Port, Database)

	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}

	return db, err
}

// 이 함수는 사용자 이름을 받아 ID를 반환한다.
// 사용자가 존재하지 않으면 -1을 반환한다.
func exist(name string) int {
	username := strings.ToLower(name)

	db, err := OpenConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userId := -1

	row, err := db.Query(`select "id" from "users" where username='%s'`, username)
	if err != nil {
		return -1
	}
	for row.Next() {
		var id int
		err = row.Scan(&id)
		if err != nil {
			fmt.Println("name_scan", err)
			return -1
		}
		userId = id
	}
	defer row.Close()
	return userId
}

// 데이터베이스에 사용자를 추가하는 함수이다.
// 성공하면 새로운 사용자의 id를 반환하고 실패하면 -1을 반환한다.
func AddUser(data Userdata) int {

	db, err := OpenConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	userName := data.Username
	userId := exist(userName)

	if userId != -1 {
		fmt.Println("이미 존재하는 유저이다.")
		return -1
	}

	// 매개변수를 이용한 쿼리만들기
	statement := `insert into users(username) values ($1)`
	_, err = db.Exec(statement, userName)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	userId = exist(userName)
	if userId == -1 {
		return userId
	}

	// userdata 테이블에 입력
	statement = `insert into userdata(userid,name,surname,dexcription) values($1,$2,$3,$4)`

	// Exec() 리턴되는 row 값이 없이 실행!
	_, err = db.Exec(statement, userId, data.Name, data.Surname, data.Descrption)
	if err != nil {
		fmt.Println("db.Exec()", err)
		return -1
	}

	return userId
}

func DeleteUser(id int) error {
	db, err := OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	statement := `select username from users where id = $1`
	row, err := db.Query(statement, id)
	if err != nil {
		return err
	}
	var username string
	if row.Next() {
		err = row.Scan(&username)
		if err != nil {
			return err
		}
	}
	defer row.Close()

	// id로 반환된 값이랑 id가 같다... 다를 수 있나?ㅡ,.ㅡ
	if exist(username) != id {
		return fmt.Errorf("user with id %d does not exist", id)
	}

	// Userdata에서 지운다.
	deleteStatement := `delete from userdata where userid=$1`
	_, err = db.Exec(deleteStatement, id)
	if err != nil {
		return err
	}

	deleteStatement = `delete form users where id = $1`
	_, err = db.Exec(deleteStatement, id)
	if err != nil {
		return err
	}
	return nil
}

func ListUsers() ([]Userdata, error) {
	records := []Userdata{}

	db, err := OpenConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	statement := `select * from userdata`
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var userdata Userdata
		err = rows.Scan(&userdata.ID, &userdata.Name, &userdata.Surname, &userdata.Descrption)
		if err != nil {
			return nil, err
		}
		records = append(records, userdata)
	}
	defer rows.Close()

	return records, nil
}

func UpdateUser(data Userdata) error {
	db, err := OpenConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	userId := exist(data.Username)
	if userId == -1 {
		return errors.New("존재하지 않는 사용자이다")
	}

	data.ID = userId

	statement := `update userdata set name=$1, surname=$2, description=$3 where UserId = $4`

	_, err = db.Exec(statement, data.Name, data.Surname, data.Descrption, data.ID)
	if err != nil {
		return err
	}

	return nil
}
