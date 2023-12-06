package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var Time []int
    var Distance []int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words := strings.Fields(scanner.Text())
        for i, word := range words {
            if i == 0 {
                continue
            }
            num, err := strconv.Atoi(word)
            if err != nil {
                panic(err)
            }
            if words[0] == "Time:" {
                Time = append(Time, num)
            } else if words[0] == "Distance:" {
                Distance = append(Distance, num)
            }
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

	fmt.Printf("Time: %v\n", Time)
	fmt.Printf("Distance: %v\n", Distance)

	// make an array the same length as Time and Distance
	vals := driver(Time, Distance)
	fmt.Printf("vals: %v\n", vals)
	ans := 1
	for _, val := range vals {
		ans *= val
	}

	print(ans)
}

func driver(Time []int, Distance []int) []int {
	waysToBeat := make([]int, len(Time))
	for i := 0; i < len(Time); i++ {
		v := calc(Time[i], Distance[i])
		waysToBeat[i] = v
	}
	return waysToBeat
}

func calc(Time int, Distance int) int {
	fmt.Printf("Time: %d\n", Time)
	fmt.Printf("Distance: %d\n", Distance)

	i1 := 0
	i2 := 0

	for i := 0; i < Time; i++ {
		distCanTravel := i * (Time - i)
		if distCanTravel > Distance {
			i1 = i
			fmt.Println(distCanTravel)
			fmt.Println(i1)
			break
		}
	}
	for i := Time; i > 0; i-- {
		distCanTravel := i * (Time - i)
		if distCanTravel > Distance {
			i2 = i
			fmt.Println(distCanTravel)
			fmt.Println(i2)
			break
		}
	}

	return i2 - i1 + 1
}