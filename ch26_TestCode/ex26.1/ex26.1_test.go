/*
go mod init goproject/ch26_TestCode/ex26.1
go get github.com/stretchr/testif

코드를 다 작성하고
해당 디렉터리의 go.mod 파일에서 디펜던시를 업데이트해줘야 됨.(코드에 필요한 디펜던시 불러옴)
그러고 run test 하고 실행됨

*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestSquare1(t *testing.T) {
// 	rst := square(9)
// 	if rst != 81 {
// 		t.Errorf("square(9) should be 81 but square(9) returns %d", rst)
// 	}
// }

// func TestSquare2(t *testing.T) {
// 	res := square(3)
// 	if res != 3*3 {
// 		t.Errorf("square(3) should be 9 but square(3) returns %d", res)
// 	}
// }

// "github.com/stretchr/testify/assert" 패키지 사용
func TestSquare3(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(81, square(9), "square(9) should be 81")
}

func TestSquare4(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(9, square(3), "square(9) should be 81")
}
