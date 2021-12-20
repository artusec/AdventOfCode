package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var (
	open  = []rune{'(', '[', '{', '<'}
	close = []rune{')', ']', '}', '>'}
)

type chunk struct {
	char        string
	kind        int
	openOrClose bool
}

func main() {
	lines := getInput()
	fmt.Println(day10(lines))
}

func day10(lines []string) (int, int) {

	var pila []chunk
	var completed []string
	var newElement chunk
	part1 := 0
	corrupt := false

	for _, line := range lines {
		for _, char := range line {
			if close, typeChunk := getType(char); !close {
				// ( [ { <
				newElement = chunk{string(char), typeChunk, close}
				pila = push(pila, newElement)
			} else {
				// ) ] } >
				if len(pila) > 0 && !pila[len(pila)-1].openOrClose && pila[len(pila)-1].kind == typeChunk {
					pila = pop(pila)
				} else {
					part1 += getValueCorrupts(typeChunk)
					corrupt = true
					break
				}
			}
		}
		if !corrupt {
			if len(pila) > 0 {
				news := ""
				for _, char := range pila {
					news = string(close[char.kind]) + news
				}
				completed = append(completed, news)
			}
		} else {
			corrupt = false
		}
		pila = nil
	}
	return part1, calcScore(completed)
}

// false -> open, true -> close
func getType(char rune) (bool, int) {
	for i, element := range open {
		if char == element {
			return false, i
		}
	}
	for i, element := range close {
		if char == element {
			return true, i
		}
	}
	return false, -1
}

func calcScore(completes []string) int {
	var scores []int
	result := 0
	for _, e := range completes {
		for _, char := range e {
			_, kind := getType(char)
			result *= 5
			result += (kind + 1)
		}
		scores = append(scores, result)
		result = 0
	}
	sort.Ints(scores)
	return scores[(len(scores) / 2)]
}

func getValueCormpleted(char rune) int {
	return int(char) + 1
}

// 3 57 1197 25137 ?? TODO: find a formula
func getValueCorrupts(element int) int {
	switch element {
	case 0:
		return 3
	case 1:
		return 57
	case 2:
		return 1197
	case 3:
		return 25137
	}
	return -1
}

func push(pila []chunk, newElement chunk) []chunk {
	return append(pila, newElement)
}

func pop(pila []chunk) []chunk {
	return pila[0 : len(pila)-1]
}

func getInput() []string {

	var lines []string
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
