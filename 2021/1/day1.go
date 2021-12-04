package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := getInput()
	fmt.Println("[*] Part 1: ", part1(input))
	fmt.Println("[*] Part 2: ", part2(input))
}

func part1(input []int) int {
	count := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			count++
		}
	}
	return count
}

func part2(input []int) int {
	count := 0
	for i := 2; i < len(input)-1; i++ {
		if input[i-2]+input[i-1]+input[i] < input[i-1]+input[i]+input[i+1] {
			count++
		}
	}
	return count
}

func getInput() []int {
	var input []int

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input,
			func() int {
				num, err := strconv.Atoi(scanner.Text())
				if err != nil {
					log.Fatal(err)
				}
				return num
			}(),
		)
	}
	return input
}
