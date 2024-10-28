package main

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score.omitempty"`
}

var students map[int]Student
var lastId int

func SetupHandlers(g *gin.Engine) {
	// 웹 핸들러 셋팅
	g.GET("/students", GetStudentListHandler)
	g.GET("/students/:id", GetStudentHandler)
	g.POST("/students", PostStudentHandler)
	g.DELETE("/students/:id", DeleteStudentHandler)

	students = make(map[int]Student)
	students[1] = Student{1, "aaa", 16, 87}
	students[2] = Student{2, "bbb", 18, 98}
	lastId = 2
}

// Id로 정렬하는 인터페이스
type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentListHandler(c *gin.Context) {
	list := make(Students, 0)
	for _, student := range students {
		list = append(list, student)
	}
	sort.Sort(list)
	c.JSON(http.StatusOK, list)
}
func GetStudentHandler(c *gin.Context) {
	idstr := c.Param("id")
	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	student, ok := students[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, student)
}
func PostStudentHandler(c *gin.Context) {
	var student Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	lastId++
	student.Id = lastId
	students[lastId] = student
	c.String(http.StatusCreated, "Success to add id:%d", lastId)
}
func DeleteStudentHandler(c *gin.Context) {
	idstr := c.Param("id")
	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	_, ok := students[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	delete(students, id)
	c.String(http.StatusOK, "success to delete")
}

func main() {
	r := gin.Default()
	SetupHandlers(r)
	r.Run(":3000")
}
