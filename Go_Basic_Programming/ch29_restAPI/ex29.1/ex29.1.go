package main

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/gorilla/mux"

	"strconv"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter() // gorilla/mux를 만든다.

	// 학생 전체 리스트 불러오기
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	// 해당 id의 학생 정보 불러오기
	mux.HandleFunc("/students/{id:[1-9]+}", GetStudentHandler).Methods("GET")
	// 학생 데이터 추가하기
	mux.HandleFunc("/students", PostStudentHandler).Methods("POST")
	// 학생 데이터 삭제하기
	mux.HandleFunc("/students/{id:[1-9]+}", DeleteStudentHandler).Methods("Delete")

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 17, 94}
	lastId = 2

	return mux
}

type Students []Student //Id로 정렬하는 인터페이스
func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list) // 학생 목록을 Id로 정렬
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}
func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}
func PostStudentHandler(w http.ResponseWriter, r *http.Request) {
	var student Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	w.WriteHeader(http.StatusCreated)
}
func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	_, ok := students[id]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	delete(students, id)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
