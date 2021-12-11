package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type vent struct {
	begin coordinate
	end   coordinate
}

func main() {
	vents := getInput()
	log.Println("[*] Part 1: ", part1(vents))
	log.Println("[*] Part 2: ", part2(vents))
}

func part1(vents []vent) int {
	ocean := make(map[string]int)
	var begin, end int
	var axis, ok bool
	for _, vent := range vents {
		if vent.begin.x == vent.end.x {
			// y changes
			axis = false
			if vent.begin.y < vent.end.y {
				begin = vent.begin.y
				end = vent.end.y
			} else {
				begin = vent.end.y
				end = vent.begin.y
			}
			ok = true
		} else {
			// extra check for diagonals
			if vent.begin.y == vent.end.y {
				// x changes
				axis = true
				if vent.begin.x < vent.end.x {
					begin = vent.begin.x
					end = vent.end.x
				} else {
					begin = vent.end.x
					end = vent.begin.x
				}
				ok = true
			}
		}
		if ok {
			for i := begin; i <= end; i++ {
				if axis {
					ocean[strconv.Itoa(i)+","+strconv.Itoa(vent.begin.y)]++
				} else {
					ocean[strconv.Itoa(vent.begin.x)+","+strconv.Itoa(i)]++
				}
			}
			ok = false
		}
	}
	sum := 0
	for _, value := range ocean {
		if value > 1 {
			sum++
		}
	}
	return sum
}

func part2(vents []vent) int {
	return 0
}

func getInput() []vent {

	var vents []vent
	var vent vent
	var temp []int

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input := append(strings.Split(strings.Split(scanner.Text(), " -> ")[0], ","), strings.Split(strings.Split(scanner.Text(), " -> ")[1], ",")...)
		for _, i := range input {
			num, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			temp = append(temp, num)
		}
		vent.begin.x = temp[0]
		vent.begin.y = temp[1]
		vent.end.x = temp[2]
		vent.end.y = temp[3]
		temp = nil
		vents = append(vents, vent)
	}

	return vents
}
