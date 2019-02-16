package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func getResultingFrequency(input []int) int64 {
	var sum int64

	for _, val := range input {
		sum = sum + int64(val)
	}

	return sum
}

func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {

	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("file opening error:", err)
	}

	reader := bufio.NewReader(file)

	input, err := readInts(reader)
	if err != nil {
		fmt.Println(err)
	}

	sum := getResultingFrequency(input)

	fmt.Printf("Sum = %d\n", sum)
}
