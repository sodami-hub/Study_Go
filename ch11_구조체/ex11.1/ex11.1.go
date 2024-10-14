package main

import "fmt"

type House struct {
	Address string  // 초기값 : ""
	Size    int     // 0
	Price   float64 // 0.0
	Type    string  // ""
}

func main() {
	//===== 초기화 방식 2 구조체 초기값 없이 초기화만
	var houseInit House = House{}
	fmt.Println(houseInit)

	// ====== 초기화 방식 1 - 직접 대입?
	var house House
	house.Address = "경기도 성남시..."
	house.Size = 28
	house.Price = 6.2
	house.Type = "APT"

	fmt.Println("주소 :", house.Address)
	fmt.Printf("크기 : %d평\n", house.Size)
	fmt.Printf("가격 : %0.2f억원\n", house.Price)
	fmt.Println("타입 :", house.Type)

	// 초기화 방식 3. 모든 변수 초기화
	// 구조체 변수의 초기화는 아래와 같은 방법도 가능하다. 아래는 모든 변수를 초기화하는 경우이다.
	var house02 House = House{
		"서울시 서초구",
		30,
		25,
		"APT", // 여러줄로 초기화하는 경우 제일 뒤에 ,를 붙이도록 한다.(한 줄로 나열할때는 상관없다.)
	}
	fmt.Println(house02.Address, house02.Price, house02.Size, house02.Type)

	// 초기화 방식 4. 일부 필드값 초기화.
	// 구조체 변수의 일부 필드값만 초기화하는 경우
	var house03 House = House{
		Size: 34,
		Type: "단독주택",
	}

	fmt.Println(house03.Address, house03.Price, house03.Size, house03.Type)
}
