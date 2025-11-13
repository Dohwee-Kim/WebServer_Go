package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//lottery filename count as command line arguments
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: lotto <filename> <count>")
	}

	filename := os.Args[1]
	count, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Can not convert count:", count)
		return
	}

	candidates, err := readCandidates(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read candidates", err)
		return
	}

	winners := make([]string, count)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		n := rand.Intn(len(candidates))
		winners[i] = candidates[n]

		candidates = append(candidates[:n], candidates[n+1:]...)
	}

	fmt.Println("Winners are:")
	for _, winner := range winners {
		fmt.Println(winner)
	}
}

// 입력타입 = 스트링, 파일이름 , 출력타입 = 리스트 오브 스트링, 에러
func readCandidates(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close() //바로 종료되는것이 아니라, 함수 종료시점에 닫아짐
	scanner := bufio.NewScanner(file)
	var candidates []string
	for scanner.Scan() {
		candidates = append(candidates, scanner.Text())
	}

	return candidates, nil
}
