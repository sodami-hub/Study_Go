package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"housework"
	/*
		이 애플리케이션의 목적은 데이터 직렬화이기 때문에 여러 데이터 직렬화 포맷을 사용하여 데이터를
		저장해보겠다. 이를 통해 다양한 직렬화 포맷으로 전환하는 것이 얼마나 간단한 일인지 알수있다.
		이를 위해서 아래의 직렬화 포맷들을 임포트한다.
	*/
	// storage "json"
	storage "gob"
	// storage "protobuf"
)

var dataFile string

// =============== 집안일 어플리케이션의 초기화 코드
func init() {
	// StringVar 함수는 dataFile 변수를 커맨드 라인 플래그 -file에 연결한다. 기본값은 housework.db 이며
	// 사용법(flag.Usage)은 data file 이다.
	// go run housework.go -file=mydata.db 로 하면 연결되는 파일이름이 바뀐다.
	flag.StringVar(&dataFile, "file", "housework.db", "data file")
	/*
		이 코드는 커맨드 라인의 매개변수와 사용법을 나타낸다. 매개변수로 add와 함께 쉼표로 구분된 집안일을 받아 목록에 더하거나
		또는 매개변수로 complete와 함께 완료로 마킹할 집안일의 개수를 숫자로 받는다.
		현재는 별도로 커맨드 라인의 옵션을 받지 않기 때문에 현 상태의 집안일의 목록을 보여준다.
	*/
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			`Usage: %s [flags] [add chore, ... |complete #]
		add add comma-separated chores
		complete complete designated chore
		Flags :
		`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

// ======= 파일에서 집안일 데이터를 역직렬화하기
// 스토리지(파일)를 불러와서 housework.Chore 구조체 슬라이스에 데이터를 저장한다.
func load() ([]*housework.Chore, error) {
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		return make([]*housework.Chore, 0), nil
	}

	df, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := df.Close(); err != nil {
			fmt.Printf("closing data file: %v", err)
		}
	}()

	return storage.Load(df)
}

// ======== 메모리상의 집안일을 스토리지(파일)에 저장한다.
func flush(chores []*housework.Chore) error {
	df, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer func() {
		if err := df.Close(); err != nil {
			fmt.Printf("closeing data file: %v", err)
		}
	}()

	return storage.Flush(df, chores)
}

// ============= 표준 출력으로 집안일 목록 출력하기
func list() error {
	chores, err := load()
	if err != nil {
		return err
	}

	if len(chores) == 0 {
		fmt.Println("You're all caught up!")
		return nil
	}

	fmt.Println("#\t[X]\tDescription")
	for i, chore := range chores {
		c := " "
		if chore.Complete {
			c = "X"
		}
		fmt.Printf("%d\t[%s]\t%s\n", i+1, c, chore.Descrption)
	}
	return nil
}

// ================= 목록에 집안일 추가하기
func add(s string) error {
	chores, err := load()
	if err != nil {
		return err
	}

	for _, chore := range strings.Split(s, ",") {
		if desc := strings.TrimSpace(chore); desc != "" {
			chores = append(chores, &housework.Chore{
				Descrption: desc,
			})
		}
	}
	return flush(chores)
}

// ========================== 완료한 집안일 마킹하기
func complete(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}

	chores, err := load()
	if err != nil {
		return err
	}

	if i < 1 || i > len(chores) {
		return fmt.Errorf("chore %d not found", i)
	}

	chores[i-1].Complete = true

	return flush(chores)
}

// ======================= 집안일 애플리케이션의 메인 로직
func main() {
	flag.Parse()

	var err error
	/*
		flag.Arg(0) 과 os.Args[0]은 서로 다른 값을 가진다.
		os.Args[0]은 프로그램의 실행 파일 이름(경로)을 값으로 가진다.
		flag.Arg(0)은 커맨드 라인 플래그를 제외한 첫 번째 인수를 반환한다. 이는 flag.Parse() 함수가 호출된 후에 사용가능하다.
	*/
	switch strings.ToLower(flag.Arg(0)) {
	case "add":
		err = add(strings.Join(flag.Args()[1:], " "))
	case "complete":
		err = complete(flag.Arg(1))
	}

	if err != nil {
		log.Fatal(err)
	}

	err = list()
	if err != nil {
		log.Fatal(err)
	}

}
