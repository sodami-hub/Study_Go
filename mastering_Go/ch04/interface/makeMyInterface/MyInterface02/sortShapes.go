package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const min = 1
const max = 5

// rF64()함수는 float64 값을 무작위로 생성
func rF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}

type Cube struct {
	x float64
}

type Cuboid struct {
	x float64
	y float64
	z float64
}

type Sphere struct {
	r float64
}

// Cube 에서 Shape3D의 인터페이스 구현
func (c Cube) Vol() float64 {
	return c.x * c.x * c.x
}

// Cuboid의 Shape3D 인터페이스 구현
func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

// Sphere의 Shape3D 인터페이스 구현
func (c Sphere) Vol() float64 {
	return c.r * c.r * c.r
}

// shapes 데이터 타입에서 sort.Interface 를 사용한다.
type shapes []Shape3D

func (s shapes) Len() int {
	return len(s)
}
func (s shapes) Less(x, y int) bool {
	return s[x].Vol() < s[y].Vol()
}
func (s shapes) Swap(x, y int) {
	s[x], s[y] = s[y], s[x]
}

func PrintShapes(a shapes) {
	for _, v := range a {
		switch v.(type) {
		case Cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case Cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case Sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		default:
			fmt.Println("Unknown data type!")
		}
	}
	fmt.Println()
}

func main() {
	data := shapes{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < 3; i++ {
		cube := Cube{rF64(min, max)}
		cuboid := Cuboid{rF64(min, max), rF64(min, max), rF64(min, max)}
		sphere := Sphere{rF64(min, max)}
		data = append(data, cube)
		data = append(data, cuboid)
		data = append(data, sphere)
	}
	PrintShapes(data)

	// 정렬
	sort.Sort(shapes(data))
	PrintShapes(data)

	// 역순 정렬
	sort.Sort(sort.Reverse(data))
	PrintShapes(data)
}
