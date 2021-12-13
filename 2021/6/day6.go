package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	ocean := getInput()
	log.Println("[*] Part 1: ", lanternfishCount(ocean, 80))
	log.Println("[*] Part 2: ", lanternfishCount(ocean, 256))
}

func lanternfishCount(ocean []int, days int) int {

	fishes := make([]int, 9)
	for _, fish := range ocean {
		fishes[fish]++
	}

	for i := 0; i < days; i++ {
		temp := fishes[0]
		for j := 0; j < 8; j++ {
			fishes[j] = fishes[j+1]
		}
		fishes[6] = fishes[6] + temp
		fishes[8] = temp
	}

	sum := 0
	for _, value := range fishes {
		sum = sum + value
	}
	return sum
}

func getInput() []int {

	var lanternfish []int
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	temp := strings.Split(scanner.Text(), ",")
	for _, i := range temp {
		lanternfish = append(lanternfish,
			func() int {
				num, err := strconv.Atoi(i)
				if err != nil {
					log.Fatal(err)
				}
				return num
			}(),
		)
	}
	return lanternfish
}
