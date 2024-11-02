package main

// Read(), Close() 메서드를 포함한 인터페이스
type Reader interface {
	Read() (n int, err error)
	Close() error
}

// Write(), Close() 메서드를 포함한 인터페이스
type Writer interface {
	Write() (n int, err error)
	Close() error
}

// 인터페이스를 포함하는 인터페이스
// Write(), Read(), Close() 메서드를 포함한 인터페이스
type ReadWriter interface {
	Reader
	Writer
}
