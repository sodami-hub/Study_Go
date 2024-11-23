/*
이 패키지는 mysql 서버의 두 테이블을 관리한다

해당 테이블은 다음과 같다
- Users
- Userdata

두 테이블의 정의는 다음과 같다.

create database go;
use go;

create table Users (

		id serial,
	    Username varchar(100) primary key
	    );

create table userdata (

		UserId int not null,
	    name varchar(100),
	    surname varchar(100),
	    description varchar(200)
	    );

		이곳은 코드 형식으로 렌더링된다.

이곳은 코드 형식으로 렌더링되지 않는다.
*/
package document

import (
	"database/sql"
	"fmt"
	"strings"
	// _ "github.com/go-sql-driver/mysql"  - 파일크기를 줄이기 위해서 제거
)

// 사용할 DB의 두개의 테이블의 데이터를 하나로 합쳤다.
type Userdata struct {
	ID          int
	Username    string
	Name        string
	Surname     string
	Description string
}

// mysql 서버의 연결 상세 정보를 담고 있는 전역 변수가 있다.
var (
	Hostname = ""
	Port     = ""
	Username = ""
	Password = ""
	Database = ""
)

// OpenConnection()은 mysql에 연결을 맺을 때 사용하며 패키지의 다른 함수들에서 사용한다.
func OpenConnection() (*sql.DB, error) {
	var db *sql.DB

	return db, nil
}

// 이 함수는 사용자 이름을 갖고 있는 사용자 ID를 반환한다.
// 사용자가 존재하지 않으면 -1을 반환한다.
func exist(name string) int {
	fmt.Println("searching user", name)
	return 0
}

// 데이터베이스에 사용자를 추가하는 함수이다.
// 성공하면 새로운 사용자의 id를 반환하고 실패하면 -1을 반환한다.
func AddUser(data Userdata) int {
	data.Username = strings.ToLower(data.Username)
	return -1
}

// DeleteUser는 사용자가 존재하는 경우 해당 사용자를 삭제한다.
// 삭제하고자 하는 사용자의 ID가 필요하다.
func DeleteUser(id int) error {
	fmt.Println(id)
	return nil
}

// ListUsers는 데이터베이스의 모든 사용자를 찾아 records슬라이스에 저장해서 반환한다.
func ListUsers() ([]Userdata, error) {
	records := []Userdata{}
	return records, nil
}

// UpdateUser는 주어진 Userdata 구조체를 이용해서 해당 사용자의 정보를 업데이트한다.
// 업데이트할 사용자의 ID는 함수 안에서 찾는다.
func UpdateUser(data Userdata) error {
	fmt.Println(data)
	return nil
}
