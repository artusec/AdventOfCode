package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	crabs := getInput()
	log.Println("[*] Part 1: ", part1(crabs))
	log.Println("[*] Part 2: ", part2(crabs))
}

func part1(crabs []int) int {

	max := 0
	for num := range crabs {
		if num > max {
			max = num
		}
	}

	min := math.MaxInt64
	sum := 0
	for position := 0; position < max; position++ {
		for _, crab := range crabs {
			sum += abs(crab - position)
		}
		if sum < min {
			min = sum
		}
		sum = 0
	}

	return min
}

func part2(crabs []int) int {

	max := 0
	for num := range crabs {
		if num > max {
			max = num
		}
	}

	min := math.MaxInt64
	sum := 0
	for position := 0; position <= max; position++ {
		for _, crab := range crabs {
			sum += summation(abs(crab - position))
		}
		if sum < min {
			min = sum
		}
		sum = 0
	}

	return min
}

func summation(num int) int {
	return (num * (num + 1) / 2)
}

func abs(num int) int {
	if num < 0 {
		return (num * -1)
	}
	return num
}

func getInput() []int {

	var crabs []int
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	temp := strings.Split(scanner.Text(), ",")
	for _, i := range temp {
		crabs = append(crabs,
			func() int {
				num, err := strconv.Atoi(i)
				if err != nil {
					log.Fatal(err)
				}
				return num
			}(),
		)
	}
	return crabs
}
