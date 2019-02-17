package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func getResultingFrequency(frequenciesChange []int) int64 {
	var sum int64
	for _, f := range frequenciesChange {
		sum = sum + int64(f)
	}
	return sum
}

func getRepeatedFrequency(frequenciesChange []int) int64 {
	hash := make(map[int64]bool)
	var sum int64
	for {
		for _, f := range frequenciesChange {
			sum = sum + int64(f)
			if _, ok := hash[sum]; ok {
				return sum
			}
			hash[sum] = true
		}
	}
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

	frequenciesChange, err := readInts(reader)
	if err != nil {
		fmt.Println(err)
	}

	sum := getResultingFrequency(frequenciesChange)

	fmt.Printf("Sum = %d\n", sum)

	repeated := getRepeatedFrequency(frequenciesChange)

	fmt.Printf("Repeated = %d\n", repeated)
}
