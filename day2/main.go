package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func countEveryLetterOfEachBoxID(boxIds []string) []map[rune]int {
	boxIdslettersCount := make([]map[rune]int, len(boxIds))
	for idx, id := range boxIds {
		alph := make(map[rune]int)
		for _, x := range id {
			if val, ok := alph[x]; ok {
				alph[x] = val + 1
			} else {
				alph[x] = 1
			}
		}
		boxIdslettersCount[idx] = alph
	}
	return boxIdslettersCount
}

func calculateCheckSum(boxIdslettersCount []map[rune]int) int {
	exactly2 := 0
	exactly3 := 0
	for _, boxIDLetterCount := range boxIdslettersCount {
		for _, val := range boxIDLetterCount {
			if val == 2 {
				exactly2++
				break
			}
		}
		for _, val := range boxIDLetterCount {
			if val == 3 {
				exactly3++
				break
			}
		}
	}
	return exactly2 * exactly3
}

func readInputs(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func main() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("file opening error:", err)
	}

	reader := bufio.NewReader(file)

	boxIds, err := readInputs(reader)

	if err != nil {
		fmt.Println("file scan error:", err)
	}

	boxIdslettersCount := countEveryLetterOfEachBoxID(boxIds)
	checkSum := calculateCheckSum(boxIdslettersCount)
	fmt.Printf("checksum = %d", checkSum)
}
