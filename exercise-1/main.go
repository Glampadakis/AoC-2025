package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type direction int

const (
	left direction = iota
	right
	maxNo        = 99
	leastNo      = 0
	totalNumbers = maxNo + 1
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var res, rowCo int64
	curr := int64(50)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dir, steps, err := parseLine(scanner.Text())
		if err != nil {
			log.Fatalf("parsing line:%v", err)
		}
		rowCo++
		res += countInBetweenClicks(dir, curr, steps)
		curr = rotateDial(dir, curr, steps)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result is %d, rowCo:%d", res, rowCo)
}

func parseLine(l string) (direction, int64, error) {
	d := l[0]
	var dir direction
	switch d {
	case 'l', 'L':
		dir = left
	case 'r', 'R':
		dir = right
	default:
		return 0, 0, fmt.Errorf("invalid direction")
	}
	steps, err := strconv.ParseInt(l[1:], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid step couns")
	}
	return dir, steps, nil
}

func rotateDial(dir direction, curr, steps int64) int64 {
	if dir == left {
		curr = (curr - steps) % totalNumbers
		if curr < 0 {
			curr += totalNumbers
		}
		return curr
	}
	curr = (curr + steps) % totalNumbers
	return curr
}

func countInBetweenClicks(dir direction, curr, steps int64) (clicks int64) {
	clicks = 0
	distToZero := curr
	if distToZero == 0 {
		distToZero = 100
	}
	if dir == left {
		if steps >= distToZero {
			clicks = 1
			clicks += (steps - distToZero) / totalNumbers
		}
		return clicks
	}

	clicks += (steps + curr) / totalNumbers
	return clicks
}
