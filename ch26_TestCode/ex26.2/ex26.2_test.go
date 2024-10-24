// 테스트 코드 + 벤치마크

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci1(-1), "fibo1(-1) should be 0")
	assert.Equal(0, fibonacci1(0), "fibo1(0) should be 0")
	assert.Equal(1, fibonacci1(1), "fibo1(1) should be 1")
	assert.Equal(2, fibonacci1(3), "fibo1(3) should be 2")
	assert.Equal(233, fibonacci1(13), "fibo1(13) should be 233")
}

func TestFibonacci2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, fibonacci2(-1), "fibo2(-1) should be 0")
	assert.Equal(0, fibonacci2(0), "fibo2(0) should be 0")
	assert.Equal(1, fibonacci2(1), "fibo2(1) should be 1")
	assert.Equal(2, fibonacci2(3), "fibo2(3) should be 2")
	assert.Equal(233, fibonacci2(13), "fibo2(13) should be 233")
}

// go test -bench .  한번에 모든 벤치마크 해서 비교할 수 있다.

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N 만큼 반복
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ { // b.N 만큼 반복
		fibonacci2(20)
	}
}
