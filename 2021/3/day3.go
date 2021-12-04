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

	oxygenRating, err := strconv.ParseInt(extractList(input, true, 0)[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	co2Rating, err := strconv.ParseInt(extractList(input, false, 0)[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return (oxygenRating * co2Rating)
}

func extractList(input []string, mode bool, bit int) []string {

	// Base case
	if len(input) == 1 {
		return input
	}

	var inputONES []string
	var inputZEROS []string
	var bitsOne = 0
	for _, line := range input {
		if line[bit] == 49 {
			bitsOne++
			inputONES = append(inputONES, line)
		} else {
			inputZEROS = append(inputZEROS, line)
		}
	}

	/*
		Input	Output
		A	B	A XNOR B
		0	0	1
		0	1	0
		1	0	0
		1	1	1
	*/

	/*
		Lo que buscamos es una XNOR:
		Si coincide que hay mas unos y que buscamos la mayor cantidad de lineas (mode = true)
		o si hay menos unos, pero buscamos la menor cantidad de lineas (mode = false). Para
		hacer una XNOR simplemente buscamos si las variables son distintas -> A != B.
		A = (bitsOne >= (len(input) - bitsOne)
		B = mode
	*/
	bit++
	if !((bitsOne >= (len(input) - bitsOne)) != mode) {
		return extractList(inputONES, mode, bit)
	}

	return extractList(inputZEROS, mode, bit)
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
