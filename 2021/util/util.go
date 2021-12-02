package util

import (
	"bufio"
	"os"
	"strconv"
)

// ReadInts reads newline-separated integers from a file into an []int
func ReadInts(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	return numbers, scanner.Err()
}

// ReadStrings reads lines from a file into a []string
func ReadStrings(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Sum(in []int) int {
	sum := in[0]
	for _, num := range in[1:] {
		sum += num
	}
	return sum
}
