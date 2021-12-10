package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type number struct {
	data  int
	check bool
}

// La linea x se encarga también de llevar la cuenta de la columna x
type lineNum struct {
	numbers   []number
	winRow    int
	winColumn int
}

type board struct {
	lines []lineNum
}
type bingoData struct {
	boards []board
	input  []int
}

func main() {
	log.Println("[*] Part 1: ", part1(getInput()))
	log.Println("[*] Part 2: ", part2(getInput()))
}

func part1(bingo bingoData) int {
	for _, numOut := range bingo.input {
		for iBoard, board := range bingo.boards {
			for iLineNum, lineNum := range board.lines {
				for iPos, pos := range lineNum.numbers {
					if pos.data == numOut && !pos.check {
						bingo.boards[iBoard].lines[iLineNum].numbers[iPos].check = true
						bingo.boards[iBoard].lines[iLineNum].winRow++
						/* Al haber encontrado el numero en la posición x (columna x),
						nos vamos a la linea x que se encarga de llevar esa cuenta de
						las columnas, con iPos y sumamos 1 */
						bingo.boards[iBoard].lines[iPos].winColumn++
						// check win
						if (bingo.boards[iBoard].lines[iLineNum].winRow == len(lineNum.numbers)) ||
							bingo.boards[iBoard].lines[iPos].winColumn == len(board.lines) {
							sum := 0
							for _, lineWinBoard := range bingo.boards[iBoard].lines {
								for _, numWinLine := range lineWinBoard.numbers {
									if !numWinLine.check {
										sum += numWinLine.data
									}
								}
							}
							return (sum * numOut)
						}
					}
				}
			}
		}
	}
	return -1
}

func part2(bingo bingoData) int {
	var boardsWinner []int
	var lastWinnerNumber = -1
	for _, numOut := range bingo.input {
		for iBoard, board := range bingo.boards {
			if !contains(boardsWinner, iBoard) {
				for iLineNum, lineNum := range board.lines {
					for iPos, pos := range lineNum.numbers {
						if pos.data == numOut && !pos.check {
							bingo.boards[iBoard].lines[iLineNum].numbers[iPos].check = true
							bingo.boards[iBoard].lines[iLineNum].winRow++
							bingo.boards[iBoard].lines[iPos].winColumn++
							if (bingo.boards[iBoard].lines[iLineNum].winRow == len(lineNum.numbers)) ||
								(bingo.boards[iBoard].lines[iPos].winColumn == len(board.lines)) {
								boardsWinner = append(boardsWinner, iBoard)
								lastWinnerNumber = numOut
							}
						}
					}
				}
			}
		}
	}
	sum := 0
	for _, lineWinBoard := range bingo.boards[boardsWinner[(len(boardsWinner)-1)]].lines {
		for _, numWinLine := range lineWinBoard.numbers {
			if !numWinLine.check {
				sum += numWinLine.data
			}
		}
	}
	return (sum * lastWinnerNumber)
}

func contains(array []int, element int) bool {
	for _, a := range array {
		if a == element {
			return true
		}
	}
	return false
}

func getInput() bingoData {
	bingo := bingoData{}

	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// First line, bingo input
	scanner.Scan()
	temp := strings.Split(scanner.Text(), ",")
	for _, j := range temp {
		num, err := strconv.Atoi(j)
		if err != nil {
			log.Fatal(err)
		}
		bingo.input = append(bingo.input, num)
	}

	// Empty line after input, then starts the boards
	scanner.Scan()
	scanner.Text()

	var num number
	num.check = false
	var linea lineNum
	linea.winRow = 0
	linea.winColumn = 0
	var board board
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			temp = strings.Split(line, " ")
			for _, cell := range temp {
				cellNum, err := strconv.Atoi(cell)
				if err != nil {
					log.Fatal(err)
				}
				num.data = cellNum
				linea.numbers = append(linea.numbers, num)
			}
			board.lines = append(board.lines, linea)
			linea.numbers = nil
		} else {
			bingo.boards = append(bingo.boards, board)
			board.lines = nil
			linea.numbers = nil
		}
	}
	// last board because no empty line at EOF
	bingo.boards = append(bingo.boards, board)
	return bingo
}
