// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// // CardValueMap maps card labels to their numerical values
// var CardValueMap = map[string]int{
// 	"A": 14, // Ace is the highest
// 	"K": 13,
// 	"Q": 12,
// 	"J": 11,
// 	"T": 10,
// 	"9": 9,
// 	"8": 8,
// 	"7": 7,
// 	"6": 6,
// 	"5": 5,
// 	"4": 4,
// 	"3": 3,
// 	"2": 2,
// }

// // Hand represents a poker hand
// type Hand struct {
// 	Cards       []string // Cards are stored as strings
// 	Type        int
// 	OriginalStr string
// 	Bid         int
// }

// // ByType sorts hands by their type and then by card values
// type ByType []Hand

// func (h ByType) Len() int      { return len(h) }
// func (h ByType) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
// func (h ByType) Less(i, j int) bool {
// 	if h[i].Type != h[j].Type {
// 		return h[i].Type > h[j].Type
// 	}
// 	for k := 0; k < 5; k++ {
// 		if CardValueMap[h[i].Cards[k]] != CardValueMap[h[j].Cards[k]] {
// 			return CardValueMap[h[i].Cards[k]] > CardValueMap[h[j].Cards[k]]
// 		}
// 	}
// 	return false
// }

// // AnalyzeHand analyzes and categorizes a hand
// func AnalyzeHand(handStr string) Hand {
// 	cardMap := make(map[string]int)
// 	hand := Hand{OriginalStr: handStr, Cards: strings.Split(handStr, "")}

// 	for _, card := range hand.Cards {
// 		cardMap[card]++
// 	}

// 	// Determine the type of hand
// 	switch len(cardMap) {
// 	case 5:
// 		hand.Type = 1 // High card or straight
// 	case 4:
// 		hand.Type = 2 // One pair
// 	case 3:
// 		for _, count := range cardMap {
// 			if count == 3 {
// 				hand.Type = 4 // Three of a kind
// 				break
// 			}
// 		}
// 		if hand.Type == 0 {
// 			hand.Type = 3 // Two pair
// 		}
// 	case 2:
// 		for _, count := range cardMap {
// 			if count == 4 {
// 				hand.Type = 6 // Four of a kind
// 				break
// 			}
// 		}
// 		if hand.Type == 0 {
// 			hand.Type = 5 // Full house
// 		}
// 	default:
// 		hand.Type = 7 // Five of a kind
// 	}

// 	return hand
// }

// // CalculateTotalWinnings calculates the total winnings from sorted hands
// func CalculateTotalWinnings(hands []Hand) int {
// 	total := 0
// 	for i, hand := range hands {
// 		rank := len(hands) - i
// 		total += hand.Bid * rank
// 	}
// 	return total
// }

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println("Error opening file:", err)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	var hands []Hand
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		parts := strings.Split(line, " ")
// 		handStr, bidStr := parts[0], parts[1]
// 		bid, _ := strconv.Atoi(bidStr)
// 		hand := AnalyzeHand(handStr)
// 		hand.Bid = bid
// 		hands = append(hands, hand)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println("Error reading from file:", err)
// 		return
// 	}

// 	sort.Sort(ByType(hands))
// 	totalWinnings := CalculateTotalWinnings(hands)
// 	fmt.Println("Total Winnings:", totalWinnings)
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// CardValueMap maps card labels to their numerical values, with 'J' as the weakest
var CardValueMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

// Hand represents a poker hand
type Hand struct {
	Cards       []string
	Type        int
	OriginalStr string
	Bid         int
}

// ByType sorts hands by their type and then by card values
type ByType []Hand

func (h ByType) Len() int      { return len(h) }
func (h ByType) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByType) Less(i, j int) bool {
	if h[i].Type != h[j].Type {
		return h[i].Type > h[j].Type
	}
	for k := 0; k < 5; k++ {
		if CardValueMap[h[i].Cards[k]] != CardValueMap[h[j].Cards[k]] {
			return CardValueMap[h[i].Cards[k]] > CardValueMap[h[j].Cards[k]]
		}
	}
	return false
}

// AnalyzeHand analyzes and categorizes a hand with jokers
func AnalyzeHand(handStr string) Hand {
	cardMap := make(map[string]int)
	jokerCount := 0
	hand := Hand{OriginalStr: handStr, Cards: strings.Split(handStr, "")}

	for _, card := range hand.Cards {
		if card == "J" {
			jokerCount++
		} else {
			cardMap[card]++
		}
	}

	// Use jokers to form the best possible hand type
	maxCount := 0
	for _, count := range cardMap {
		if count > maxCount {
			maxCount = count
		}
	}
	maxCountWithJokers := maxCount + jokerCount

	switch {
	case maxCountWithJokers == 5:
		hand.Type = 7 // Five of a kind
	case maxCountWithJokers == 4:
		hand.Type = 6 // Four of a kind
	case maxCountWithJokers == 3 && len(cardMap) <= 2:
		hand.Type = 5 // Full house
	case maxCount == 3:
		hand.Type = 4 // Three of a kind
	case len(cardMap) == 2 || (jokerCount > 0 && len(cardMap)+jokerCount == 3):
		hand.Type = 3 // Two pair
	case len(cardMap) == 3 || (jokerCount > 0 && len(cardMap)+jokerCount == 4):
		hand.Type = 2 // One pair
	default:
		hand.Type = 1 // High card
	}

	return hand
}

// CalculateTotalWinnings calculates the total winnings from sorted hands
func CalculateTotalWinnings(hands []Hand) int {
	total := 0
	for i, hand := range hands {
		rank := len(hands) - i
		total += hand.Bid * rank
	}
	return total
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hands []Hand
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		handStr, bidStr := parts[0], parts[1]
		bid, _ := strconv.Atoi(bidStr)
		hand := AnalyzeHand(handStr)
		hand.Bid = bid
		hands = append(hands, hand)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	sort.Sort(ByType(hands))
	totalWinnings := CalculateTotalWinnings(hands)
	fmt.Println("Total Winnings:", totalWinnings)
}
