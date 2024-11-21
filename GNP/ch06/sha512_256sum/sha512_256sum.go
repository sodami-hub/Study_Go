// 서버의 파일과 tftp로 다운받은 파일의 체크섬을 비교한다.

package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s file...\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	for _, file := range flag.Args() {
		fmt.Printf("%s %s\n", checksum(file), file)
	}
}

func checksum(file string) string {
	b, err := os.ReadFile(file)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%x", sha512.Sum512_256(b))
}
