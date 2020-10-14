package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := NewReader()
	pb := NewPonyByer()
	input := reader.Read()
	if input.numOfFigures == 0{
		fmt.Println(input.actual)
		os.Exit(0)
	}
	res, _ := pb.Buy(input.desires, input.actual, input.actualToysCount)
	fmt.Println(toString(res))
}

func toString(arr []int) string {
	var sb strings.Builder
	for _, item := range arr {
		sb.WriteString(strconv.Itoa(item))
	}
	return sb.String()
}

type PonyBuyer struct {
}

func NewPonyByer() *PonyBuyer {
	return &PonyBuyer{}
}

func (pb PonyBuyer) Buy(desires map[string][]int, actual []int, numOfToys int) ([]int, int) {
	min := math.MaxInt32
	result := make([]int, 0)
	for k, v := range desires {
		newOne, num := pb.newActual(actual, v, k, numOfToys)
		if newOne != nil {
			delete(desires, k)
			curr, newNum := pb.Buy(desires, newOne, num)
			desires[k] = v
			if newNum < min {
				min = newNum
				result = curr
			}
		}
	}
	if min == math.MaxInt32 {
		return actual, numOfToys
	} else {
		return result, min
	}
}

func (pb PonyBuyer) newActual(actual, desired []int, pattern string, oldCounter int) ([]int, int) {

	for _, r := range pattern {
		num, err := strconv.Atoi(string(r))
		if err == nil {
			if actual[num] != 1 {
				return nil, 0
			}
		}
	}
	newOne := make([]int, 0)
	newOne = append(newOne, actual...)
	counter := oldCounter
	for _, num := range desired {
		newOne[num] = 1
		if actual[num] != 1 {
			counter++
		}
	}
	return newOne, counter
}

type Reader struct {
}

func NewReader() *Reader {
	return &Reader{}
}

func (rdr Reader) Read() Input {
	scanner := bufio.NewScanner(os.Stdin)

	input := rdr.scan(scanner)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	return input
}

func (rdr Reader) scan(scanner *bufio.Scanner) Input {
	state := 0
	input := Input{desires: make(map[string][]int, 0), actual: make([]int, 0)}
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			switch state {
			case noneRead:
				parseLine(scanner.Text(), state, &input)
				state++
				if input.numOfDesires == 0 {
					state++
				}
			case numbersRead:
				state = readDesires(scanner, state, &input, input.numOfDesires)
			case desiresRead:
				parseLine(scanner.Text(), state, &input)
				return input
			default:
				return input
			}
		}
	}
	return input
}

func readDesires(scanner *bufio.Scanner, state int, input *Input, total int) int {

	for i := 0; i < total; i++ {
		if i != 0 {
			scanner.Scan()
		}
		parseLine(scanner.Text(), state, input)
	}
	return state + 1
}

func parseLine(line string, state int, input *Input) {
	switch state {
	case noneRead:
		arr := strings.Split(line, " ")
		numOfFigures, err := strconv.Atoi(strings.TrimSpace(arr[0]))
		if err != nil {
			log.Fatal("can't read input", err)
		}
		numOfDesires, err := strconv.Atoi(strings.TrimSpace(arr[1]))
		if err != nil {
			log.Fatal("can't read input", err)
		}
		input.numOfFigures = numOfFigures
		input.numOfDesires = numOfDesires
	case numbersRead:
		arr := strings.Split(line, " ")
		k, v := convertToMapEntry(arr[0], arr[1])
		input.desires[k] = v
	case desiresRead:
		input.actual, input.actualToysCount = convertToIntSlice(line)

	}
}

func convertToIntSlice(str string) ([]int, int) {
	arr := make([]int, len(str))
	counter := 0
	for i, r := range str {
		if r == '1' {
			arr[i] = 1
			counter++
		} else {
			arr[i] = 0
		}
	}
	return arr, counter
}
func convertToMapEntry(actual, desired string) (string, []int) {
	var sb strings.Builder
	arr := make([]int, 0)
	for i, r := range actual {
		if r == '1' {
			sb.WriteString(strconv.Itoa(i))
		}
		des := desired[i]
		if des == '1' {
			arr = append(arr, i)
		}
	}
	return sb.String(), arr
}

const (
	noneRead = iota
	numbersRead
	desiresRead
)

type Input struct {
	numOfFigures    int
	numOfDesires    int
	desires         map[string][]int
	actual          []int
	actualToysCount int
}

type FileReader interface {
	ReadFile(string) Input
}

type ConsoleReader interface {
	Read() Input
}

type AnswerReader interface {
	ReadTheAnswer(string) Input
}
