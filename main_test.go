package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func (rdr Reader) ReadTheAnswer(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}
func (rdr Reader) ReadFile(fileName string) Input {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	input := rdr.scan(scanner)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input

}

func toString(arr []int) string {
	var sb strings.Builder
	for _, item := range arr {
		sb.WriteString(strconv.Itoa(item))
	}
	return sb.String()
}

func TestFileRead(t *testing.T) {
	reader := NewReader()
	pb := NewPonyByer()
	for i := 0; i < 6; i++ {
		input := reader.ReadFile("tests/test_input_" + strconv.Itoa(i) + ".txt")
		res, _ := pb.Buy(input.desires, input.actual, input.actualToysCount)
		stringResult := toString(res)
		expectedResult := reader.ReadTheAnswer("tests/ans_" + strconv.Itoa(i) + ".txt")
		if stringResult != expectedResult {
			t.Errorf("expected: %s, got: %s", expectedResult, stringResult)
		}

	}
}