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
	input := getInput()
	fmt.Println("[*] Part 1: ", part1(input))
	fmt.Println("[*] Part 2: ", part2(input))
}

func part1(input []string) int {
	horizontal := 0
	depth := 0

	for _, line := range input {
		instruction := strings.Split(line, " ")
		num, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		if instruction[0] == "up" {
			depth -= num
		} else if instruction[0] == "forward" {
			horizontal += num
		} else { // down
			depth += num
		}
	}
	return horizontal * depth
}

func part2(input []string) int {
	horizontal := 0
	depth := 0
	aim := 0

	for _, line := range input {
		instruction := strings.Split(line, " ")
		num, err := strconv.Atoi(instruction[1])
		if err != nil {
			log.Fatal(err)
		}

		if instruction[0] == "up" {
			aim -= num
		} else if instruction[0] == "forward" {
			horizontal += num
			depth += (aim * num)
		} else { // down
			aim += num
		}
	}
	return horizontal * depth
}

func getInput() []string {
	var input []string

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
