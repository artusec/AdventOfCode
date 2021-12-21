package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type octopus struct {
	energy int
	flash  bool
}

func main() {
	octopuses := getInput()
	fmt.Println(day11(octopuses))
}

func day11(octopuses [10][10]octopus) (int, int) {
	sum := 0
	prevSum := 0
	firstTime := -1
	result := -1
	for step := 1; step < math.MaxInt && (firstTime == -1 || result == -1); step++ {
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if !octopuses[i][j].flash {
					if octopuses[i][j].energy+1 >= 10 {
						octopuses, sum = flash(octopuses, i, j, sum)
					} else {
						octopuses[i][j].energy++
					}
				}
			}
		}
		if sum-prevSum == 100 && firstTime == -1 {
			firstTime = step
		}
		prevSum = sum
		if step == 100 {
			result = sum
		}
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopuses[i][j].flash = false
			}
		}
	}
	return result, firstTime
}

func flash(octopuses [10][10]octopus, i, j int, sum int) ([10][10]octopus, int) {
	octopuses[i][j].flash = true
	octopuses[i][j].energy = 0
	sum++
	if i-1 >= 0 && !octopuses[i-1][j].flash {
		octopuses[i-1][j].energy++
		if octopuses[i-1][j].energy >= 10 {
			octopuses, sum = flash(octopuses, i-1, j, sum)
		}
	}
	if i-1 >= 0 && j+1 < 10 && !octopuses[i-1][j+1].flash {
		octopuses[i-1][j+1].energy++
		if octopuses[i-1][j+1].energy >= 10 {
			octopuses, sum = flash(octopuses, i-1, j+1, sum)
		}
	}
	if j+1 < 10 && !octopuses[i][j+1].flash {
		octopuses[i][j+1].energy++
		if octopuses[i][j+1].energy >= 10 {
			octopuses, sum = flash(octopuses, i, j+1, sum)
		}
	}
	if i+1 < 10 && j+1 < 10 && !octopuses[i+1][j+1].flash {
		octopuses[i+1][j+1].energy++
		if octopuses[i+1][j+1].energy >= 10 {
			octopuses, sum = flash(octopuses, i+1, j+1, sum)
		}
	}
	if i+1 < 10 && !octopuses[i+1][j].flash {
		octopuses[i+1][j].energy++
		if octopuses[i+1][j].energy >= 10 {
			octopuses, sum = flash(octopuses, i+1, j, sum)
		}
	}
	if i+1 < 10 && j-1 >= 0 && !octopuses[i+1][j-1].flash {
		octopuses[i+1][j-1].energy++
		if octopuses[i+1][j-1].energy >= 10 {
			octopuses, sum = flash(octopuses, i+1, j-1, sum)
		}
	}
	if j-1 >= 0 && !octopuses[i][j-1].flash {
		octopuses[i][j-1].energy++
		if octopuses[i][j-1].energy >= 10 {
			octopuses, sum = flash(octopuses, i, j-1, sum)
		}
	}
	if i-1 >= 0 && j-1 >= 0 && !octopuses[i-1][j-1].flash {
		octopuses[i-1][j-1].energy++
		if octopuses[i-1][j-1].energy >= 10 {
			octopuses, sum = flash(octopuses, i-1, j-1, sum)
		}
	}
	return octopuses, sum
}

func getInput() [10][10]octopus {

	var octopuses [10][10]octopus

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		for j, octo := range line {
			temp, err := strconv.Atoi(string(octo))
			if err != nil {
				log.Fatal(err)
			}
			octopuses[i][j].energy = temp
			octopuses[i][j].flash = false
		}
		i++
	}

	return octopuses
}
