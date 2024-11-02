package main

import (
	"ch18_interface/ex18.2/fedex"
	"ch18_interface/ex18.2/post"
)

// Sender와 Send메소드를 인터페이스로 정의하면
// Send메소드를 갖는 모든 구조체들은 Sender형태로 객체를 만들 수 있다.
type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	var postSender Sender = &post.PostSender{}
	SendBook("어린왕자", postSender)
	SendBook("그리스인 조르바", postSender)

	var FedexSender Sender = &fedex.FedexSender{}
	SendBook("야야야야", FedexSender)
	SendBook("임마임마", FedexSender)
}
