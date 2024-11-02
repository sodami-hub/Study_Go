// 메모리 쓰레기의 재활용. 자주 할당되는 객체를 객체 풀에 넣었다가 다시 꺼내쓰면 된다.
// -> 플라이웨이트(flyweight) 패턴 방식

// 어떤 객체에 메모리가 얼마나 할당되는지 알아내는 방법
// go test -cpuprofile cpu.prof -memprofile mem.prof -bench .

package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fac := NewFlyweightFactory(1000) // 객체공장
	for i := 0; i < 1000; i++ {
		obj := fac.Create()
		obj.Somedata = "Somedata"
		fac.Dispose(obj)
	}

	fmt.Println("AlloCnt", fac.AlloCnt)
}

type FlyweightFactory struct {
	pool    []*Flyweight
	AlloCnt int
}

func (fac *FlyweightFactory) Create() *Flyweight {
	var obj *Flyweight
	if len(fac.pool) > 0 { // 재활용
		obj, fac.pool = fac.pool[len(fac.pool)-1], fac.pool[:len(fac.pool)-1]
		obj.Reuse()
	} else {
		obj = &Flyweight{} // 새로 만든다.
		fac.AlloCnt++
	}
	return obj
}

func (fac *FlyweightFactory) Dispose(obj *Flyweight) { // 반환
	obj.Dispose()
	fac.pool = append(fac.pool, obj)
}

func NewFlyweightFactory(initSize int) *FlyweightFactory {
	return &FlyweightFactory{pool: make([]*Flyweight, 0, initSize)}
}

type Flyweight struct {
	Somedata  string
	isDispose bool
}

func (f *Flyweight) Reuse() {
	f.isDispose = false
}

func (f *Flyweight) Dispose() {
	f.isDispose = true
}

func (f *Flyweight) IsDisposed() bool {
	return f.isDispose
}
