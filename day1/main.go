package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"unicode"
	"strings"
)

var strToInt = map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
	"zero": 0,
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

	// l := []string{"cpcnkvdbrqrxtfnmzbqgffivesix91fivehgrv"}
	finalSum := driver(lines)

	fmt.Printf("The answer is %d\n", finalSum)

}

func driver(lines []string) int {
	sum := 0
	for _, line := range lines {
		// newLine := stringToIntConv(line)
		fmt.Println(line)
		numIdx := getNumIndexes(line)
		strIdx := getSubstrIndexes(line)
		first, last := compareIdxs(numIdx, strIdx)
		sum += convertStrToInt(strconv.Itoa(first), strconv.Itoa(last))
	}

	return sum
}

func stringToIntConv(line string) string {
	og := line
	fmt.Printf("og: %s\n", og)
	for key, value := range strToInt {
		line = strings.Replace(line, key, strconv.Itoa(value), -1)
	}
	fmt.Printf("new: %s\n", line)
	return line
}

type numsIndex struct {
	firstInt int
	firstIntIndex int
	lastInt int
	lastIntIndex int
}

func getNumIndexes(line string) numsIndex {
	firstInt := -1
	firstIntIndex := -1
	lastInt := -1
	lastIntIndex := -1
	for idx, char := range line {
		if unicode.IsDigit(char) {
			if firstInt == -1 {
				firstInt = int(char - '0')
				firstIntIndex = idx
			} else {
				lastInt = int(char - '0')
				lastIntIndex = idx
			}
        } 
	}
	if lastInt == -1 {
		lastInt = firstInt
		lastIntIndex = firstIntIndex
	}

	fmt.Println("Integer indexes")
	fmt.Printf("firstInt: %d\n", firstInt)
	fmt.Printf("firstIntIndex: %d\n", firstIntIndex)
	fmt.Printf("lastInt: %d\n", lastInt)
	fmt.Printf("lastIntIndex: %d\n", lastIntIndex)


	return numsIndex{firstInt, firstIntIndex, lastInt, lastIntIndex}
	
}

type SubstrIndexes struct {
	firstStrIdx int
	firstStrToInt int
	lastStrIdx int
	lastStrToInt int
}

func getSubstrIndexes(line string) SubstrIndexes {
	firstStrIdx := -1
	firstStrToInt := -1
	lastStrIdx := -1
	lastStrToInt := -1
	for key, value := range strToInt {
		if idx := strings.Index(line, key); idx != -1 {
			if idx < firstStrIdx || firstStrIdx == -1 {
				firstStrIdx = idx
				firstStrToInt = value
			}
			if idx > lastStrIdx {
				lastStrIdx = idx
				lastStrToInt = value
			}
		}
		if idx := strings.LastIndex(line, key); idx != -1 {
			if idx < firstStrIdx || firstStrIdx == -1 {
				firstStrIdx = idx
				firstStrToInt = value
			}
			if idx > lastStrIdx {
				lastStrIdx = idx
				lastStrToInt = value
			}
		}
	}

	fmt.Println("String indexes")
	fmt.Printf("firstStrIdx: %d\n", firstStrIdx)
	fmt.Printf("firstStrToInt: %d\n", firstStrToInt)
	fmt.Printf("lastStrIdx: %d\n", lastStrIdx)
	fmt.Printf("lastStrToInt: %d\n", lastStrToInt)

	return SubstrIndexes{firstStrIdx, firstStrToInt, lastStrIdx, lastStrToInt}
}

func convertStrToInt(firstInt string, lastInt string) int {
	// make a string which is %c%c
	numStr := fmt.Sprintf("%s%s", firstInt, lastInt)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	fmt.Println("")
	return num
}

func compareIdxs(numIdx numsIndex, strIdx SubstrIndexes) (int, int) {
	first := -1
	last := -1
	if numIdx.firstIntIndex < strIdx.firstStrIdx {
		first = numIdx.firstInt
	} else {
		first = strIdx.firstStrToInt
	}
	if numIdx.firstIntIndex == -1 {
		first = strIdx.firstStrToInt
	} else if (strIdx.firstStrIdx == -1) {
		first = numIdx.firstInt
	}

	if numIdx.lastIntIndex > strIdx.lastStrIdx {
		last = numIdx.lastInt
	} else {
		last = strIdx.lastStrToInt
	}
	if numIdx.lastIntIndex == -1 {
		last = strIdx.lastStrToInt
	} else if (strIdx.lastStrIdx == -1) {
		last = numIdx.lastInt
	}
	fmt.Printf("first: %d\n", first)
	fmt.Printf("last: %d\n", last)
	return first, last

}