package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.Println("[*] Part 1: ", part1(getInput()))
	log.Println("[*] Part 2: ", part2(getInput()))
}

func part1(lanternfish []int) int {
	news := 0
	for i := 0; i < 80; i++ {
		for iFish, fish := range lanternfish {
			if fish == 0 {
				news++
				lanternfish[iFish] = 6
			} else {
				lanternfish[iFish]--
			}
		}
		for j := 0; j < news; j++ {
			lanternfish = append(lanternfish, 8)
		}
		news = 0
	}

	return len(lanternfish)
}

func part2(lanternfish []int) int {

	ocean := make([]int, 9)
	for _, fish := range lanternfish {
		ocean[fish]++
	}

	for i := 0; i < 256; i++ {
		temp := ocean[0]
		for j := 0; j < 8; j++ {
			ocean[j] = ocean[j+1]
		}
		ocean[6] = ocean[6] + temp
		ocean[8] = temp
	}

	sum := 0
	for _, value := range ocean {
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
