package main

import (
	// /dev/random에서 바이너리를 읽어 정수 값으로 변환하기 때문에 encoding/binary가 필요하다.

	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/dev/random")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var seed int64
	binary.Read(f, binary.BigEndian, &seed)
	fmt.Println("Seed :", seed)
}
