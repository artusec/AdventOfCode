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

func part1(displays [][]string) int {
	for i := range displays {
		displays[i] = displays[i][11:]
	}
	count := 0
	for _, line := range displays {
		for _, display := range line {
			len := len(display)
			if len == 2 || len == 3 || len == 4 || len == 7 {
				count++
			}
		}
	}
	return count
}

func part2(displays [][]string) int {

	codes := make(map[int]string)
	var unknown []string
	sum := 0

	for _, line := range displays {
		newLine := append(line[0:10], line[11:]...)
		for _, display := range newLine {
			if !gotTheValue(codes, display) {
				switch len(display) {
				case 2:
					codes[1] = display
				case 3:
					codes[7] = display
				case 4:
					codes[4] = display
				case 7:
					codes[8] = display
				default:
					unknown = append(unknown, display)
				}
			}
		}

		for _, number := range unknown {
			if !gotTheValue(codes, number) {
				switch len(number) {
				case 5:
					result, _ := containsAndMissings(codes[1], number)
					if result {
						codes[3] = number
					} else {
						_, missings := containsAndMissings(codes[4], number)
						if missings == 1 {
							codes[5] = number
						} else {
							codes[2] = number
						}
					}
				case 6:
					result, _ := containsAndMissings(codes[4], number)
					if result {
						codes[9] = number
					} else {
						result, _ := containsAndMissings(codes[1], number)
						if result {
							codes[0] = number
						} else {
							codes[6] = number
						}

					}
				}
			}
		}
		sum += decode(codes, newLine[10:])
	}
	return sum
}

func decode(codes map[int]string, values []string) int {
	stringNumber := ""
	for _, code := range values {
		for keyDict, valueDict := range codes {
			if found, perfect := containsAndMissings(code, valueDict); found {
				if perfect == 0 {
					stringNumber += strconv.Itoa(keyDict)
					break
				}
			}
		}
	}

	num, err := strconv.Atoi(stringNumber)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func gotTheValue(dict map[int]string, value string) bool {
	for _, valueDict := range dict {
		if valueDict == value {
			return true
		}
	}
	return false
}

func containsAndMissings(minor, major string) (bool, int) {
	num := 0
	for _, letterMinor := range minor {
		for _, letterMajor := range major {
			if letterMinor == letterMajor {
				num++
				break
			}
		}

	}
	if num == len(minor) {
		// perfect match (same digit)
		if num == len(major) {
			return true, 0
		}
		// major contains minor
		return true, -1
	}

	// major does not contains minor by (len(minor) - num) letters
	return false, len(minor) - num
}

func getInput() [][]string {

	var input [][]string
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), " "))
	}

	return input
}
