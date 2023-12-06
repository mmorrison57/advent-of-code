package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cards := make([][][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		colonPos := strings.Index(line, ":")
		barPos := strings.Index(line, "|")
		if colonPos != -1 && barPos != -1 {
			card := make([][]int, 2)
			for i, part := range []string{line[colonPos+1 : barPos], line[barPos+1:]} {
				numbers := strings.Fields(part)
				card[i] = make([]int, len(numbers))
				for j, number := range numbers {
					card[i][j], _ = strconv.Atoi(number)
				}
			}
			cards = append(cards, card)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// The cards variable now contains the input file data
	// fmt.Println(cards)
	sum := 0
	for card := range cards {
		sum += calculate(cards[card])
	}
	fmt.Println(sum)
}


func calculate(card [][]int) int {
	score := 0
	for i:=0; i < len(card[1]); i++ {
		for j:=0; j < len(card[0]); j++ {
			numIHave := card[1][i]
			winningNum := card[0][j]
			if numIHave == winningNum {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
				break
			}
		}
	}
	fmt.Println(score)
	return score
}

func countMatches(card [][]int) int {
	matches := 0
	for i:=0; i < len(card[1]); i++ {
		for j:=0; j < len(card[0]); j++ {
			numIHave := card[1][i]
			winningNum := card[0][j]
			if numIHave == winningNum {
				matches++
				break
			}
		}
	}
	return matches
}