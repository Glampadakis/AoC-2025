package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		repeatedNumber, dupNumber := 0, 0
		idRanges := strings.Split(input, ",")
		for _, r := range idRanges {
			startInt, endInt, err := parseRange(r)
			if err != nil {
				log.Fatal(err)
			}
			for i := startInt; i <= endInt; i++ {
				if isRepeatedNumber(i) {
					repeatedNumber += i
				}
				if isRepeatedMultipleTimes(i) {
					dupNumber += i
				}

			}
		}
		fmt.Printf("repeated number total: %d\n", repeatedNumber)
		fmt.Printf("duplicated number total: %d\n", dupNumber)
	}
}

func parseRange(r string) (int, int, error) {
	points := strings.Split(r, "-")
	if len(points) != 2 {
		return 0, 0, fmt.Errorf("invalid range")
	}
	start, err := strconv.Atoi(points[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start point")
	}
	end, err := strconv.Atoi(points[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid end point")
	}
	return start, end, nil
}

func isRepeatedNumber(d int) bool {
	s := fmt.Sprintf("%d", d)
	return s[:len(s)/2] == s[len(s)/2:]
}

func isRepeatedMultipleTimes(d int) bool {
	s := fmt.Sprintf("%d", d)
	// first of all traverse the string
initial:
	for i := 0; i < len(s); i++ {
		left := s[0:i]
		right := s[i:]
		if len(left) == 0 || len(s)%len(left) != 0 {
			continue initial
		}
		if compareString(left, right) {
			return true
		}
	}
	return false
}

func compareString(v, s string) bool {
	// if left is bigger than right, definitely false.
	if len(v) > len(s) {
		return false
	}

	// if same length, we have reached the end.
	if len(v) == len(s) {
		return v == s
	}
	// check if right starts with left.
	if !strings.HasPrefix(s, v) {
		return false
	}

	return compareString(v, s[len(v):])
}
