package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	input := getInput()
	log.Println("[*] Part 1: ", part1(input))
	log.Println("[*] Part 2: ", part2(input))
}

func part1(input []string) int64 {
	gammaRate := ""
	epsilonRate := ""
	bitsOne := 0
	for bit := 0; bit < len(input[0]); bit++ {
		for _, line := range input {
			if line[bit] == 49 {
				bitsOne++
			}
		}
		if bitsOne > (len(input) - bitsOne) {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
		bitsOne = 0
	}

	decGammaRate, err := strconv.ParseInt(gammaRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	decEpsilonRate, err := strconv.ParseInt(epsilonRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return (decGammaRate * decEpsilonRate)
}

func part2(input []string) int64 {

	var inputONES []string = make([]string, 0)
	var inputZEROS []string = make([]string, 0)
	var bitsOne = 0
	var bit int = 0

	temp := input
	for ok := true; ok; ok = (len(temp) != 1) {
		for _, line := range temp {
			if line[bit] == 49 {
				bitsOne++
				inputONES = append(inputONES, line)
			} else {
				inputZEROS = append(inputZEROS, line)
			}
		}

		if bitsOne >= (len(temp) - bitsOne) {
			temp = inputONES
		} else {
			temp = inputZEROS
		}
		inputONES = inputONES[:0]
		inputZEROS = inputZEROS[:0]
		bitsOne = 0
		bit++
	}
	oxygenRating := temp[0]

	bit = 0
	temp = input
	for ok := true; ok; ok = (len(temp) != 1) {
		for _, line := range temp {
			if line[bit] == 49 {
				bitsOne++
				inputONES = append(inputONES, line)
			} else {
				inputZEROS = append(inputZEROS, line)
			}
		}

		if bitsOne >= (len(temp) - bitsOne) {
			temp = inputZEROS
		} else {
			temp = inputONES
		}
		inputONES = inputONES[:0]
		inputZEROS = inputZEROS[:0]
		bitsOne = 0
		bit++
	}

	co2Rating := temp[0]

	oxygenRatingNum, err := strconv.ParseInt(oxygenRating, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	co2RatingNum, err := strconv.ParseInt(co2Rating, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return oxygenRatingNum * co2RatingNum
}

func getInput() []string {
	var input []string

	file, err := os.Open("./test")
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
