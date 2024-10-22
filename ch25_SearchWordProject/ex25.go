package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex // 뮤텍스로 하는건 좀 웃긴데.. 무슨 의미가 있지. ㅋㅋㅋ 예제를 보고 풀어봐야될 듯.

func stringMatcher(fileName, target string) {
	mutex.Lock()
	var scan *bufio.Scanner
	var scanLine string

	fmt.Println()
	fmt.Println("====== ", fileName, " ======")
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("file open error")
		return
	}
	scan = bufio.NewScanner(file)
	i := 0
	for scan.Scan() {
		i++
		scanLine = scan.Text()
		if strings.Contains(scanLine, target) {
			fmt.Println(i, " | ", scanLine)

		}
	}
	file.Close()
	mutex.Unlock()
	wg.Done()
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("usage : (실행명령어) (찾을단어) (대상파일)")
		return
	}

	var TargetWord string = os.Args[1]

	fileList, err := filepath.Glob(os.Args[2])
	if err != nil {
		fmt.Println("file path error!!")
		return
	}

	wg.Add(len(fileList))
	for _, fileName := range fileList {
		go stringMatcher(fileName, TargetWord)
	}

	wg.Wait()
}

/*
> .\ex25.exe int te*

======  teab.txt  ======
13  |   "internal/stringslite"
27  |   w          int
29  |   volLen     int
32  |  func (b *lazybuf) index(i int) byte {
131  |          // Turn empty string into "."
272  |  func VolumeNameLen(path string) int {

======  test.txt  ======
14  |  // Scanner provides a convenient interface for reading data such as
19  |  // function breaks the input into lines with line termination stripped. [Scanner.Split]
20  |  // functions are defined in this package for scanning a file into
32  |   maxTokenSize int       // Maximum size of a token; modified by tests.
35  |   start        int       // First non-processed byte in buf.
36  |   end          int       // End of data in buf.
38  |   empties      int       // Count of successive empty tokens.
61  |  // [Scanner] to read more data into the slice and try again with a
62  |  // longer slice starting at the same point in the input.
67  |  type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
106  |  // The underlying array may point to data that will be overwritten
118  |  // ErrFinalToken is a special sentinel error value. It is intended to be
199  |                          const maxInt = int(^uint(0) >> 1)
243  |  func (s *Scanner) advance(n int) bool {
268  |  // By default, [Scanner.Scan] uses an internal buffer and sets the
272  |  func (s *Scanner) Buffer(buf []byte, max int) {
294  |  func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
307  |  // Because of the Scan interface, this makes it impossible for the client to
309  |  func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
355  |  func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
400  |  func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
*/
