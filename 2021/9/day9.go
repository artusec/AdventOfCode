package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

type cell struct {
	number int
	check  bool
}

func main() {
	lava := getInput()
	log.Println("[*] Part 1: ", part1(lava))
	log.Println("[*] Part 2: ", part2(lava))
}

func part1(lava [][]cell) int {

	sum := 0
	for i := 0; i < len(lava); i++ {
		for j := 0; j < len(lava[i]); j++ {
			if isLowerPoint(lava, i, j) {
				sum += lava[i][j].number + 1
			}
		}
	}
	return sum
}

func isLowerPoint(lava [][]cell, x, y int) bool {
	actual := lava[x][y].number
	if (x + 1) < len(lava) {
		if lava[x+1][y].number <= actual {
			return false
		}
	}
	if (x - 1) >= 0 {
		if lava[x-1][y].number <= actual {
			return false
		}
	}
	if (y + 1) < len(lava[x]) {
		if lava[x][y+1].number <= actual {
			return false
		}
	}
	if (y - 1) >= 0 {
		if lava[x][y-1].number <= actual {
			return false
		}
	}
	return true
}

func part2(lava [][]cell) int {

	var longs []int
	for i := 0; i < len(lava); i++ {
		for j := 0; j < len(lava[i]); j++ {
			if isLowerPoint(lava, i, j) {
				longs = append(longs, getBasin(lava, i, j))
			}
		}
	}
	sort.Slice(longs, func(i, j int) bool {
		return longs[i] > longs[j]
	})
	return longs[0] * longs[1] * longs[2]
}

func getBasin(lava [][]cell, x, y int) int {

	lava[x][y].check = true
	actual := lava[x][y].number

	if actual == 9 {
		return 0
	}

	long := 1

	if (x + 1) < len(lava) {
		if (lava[x+1][y].number > actual) && !lava[x+1][y].check {
			long += getBasin(lava, x+1, y)
		}
	}
	if (x - 1) >= 0 {
		if (lava[x-1][y].number > actual) && !lava[x-1][y].check {
			long += getBasin(lava, x-1, y)
		}
	}
	if (y + 1) < len(lava[x]) {
		if (lava[x][y+1].number > actual) && !lava[x][y+1].check {
			long += getBasin(lava, x, y+1)
		}
	}
	if (y - 1) >= 0 {
		if (lava[x][y-1].number > actual) && !lava[x][y-1].check {
			long += getBasin(lava, x, y-1)
		}
	}
	return long
}

func getInput() [][]cell {

	var lava [][]cell
	var tempCell cell
	tempCell.check = false
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var tempLine []cell
		line := scanner.Text()
		for _, num := range line {
			if err != nil {
				log.Fatal(err)
			}
			tempCell.number = (int(num) - 48)
			tempLine = append(tempLine, tempCell)
		}
		lava = append(lava, tempLine)
		tempLine = nil
	}

	return lava
}
