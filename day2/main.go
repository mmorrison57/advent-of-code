package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
//     file, err := os.Open("input.txt")
//     if err != nil {
//         panic(err)
//     }
//     defer file.Close()

//     scanner := bufio.NewScanner(file)
//     sum := 0
//     for scanner.Scan() {
//         line := scanner.Text()
//         parts := strings.Split(line, ": ")
//         gameID, _ := strconv.Atoi(parts[0][5:])
//         sets := strings.Split(parts[1], "; ")
//         gamePossible := true
//         for _, set := range sets {
//             cubes := strings.Split(set, ", ")
//             for _, cube := range cubes {
//                 colorCount := strings.Split(cube, " ")
//                 count, _ := strconv.Atoi(colorCount[0])
//                 color := colorCount[1]
//                 if (color == "red" && count > 12) || (color == "green" && count > 13) || (color == "blue" && count > 14) {
//                     gamePossible = false
//                     break
//                 }
//             }
//             if !gamePossible {
//                 break
//             }
//         }
//         if gamePossible {
//             sum += gameID
//         }
//     }

//     if err := scanner.Err(); err != nil {
//         panic(err)
//     }

//     // The sum of the IDs of the possible games is now in the variable sum
// 	fmt.Println(sum)
// }

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

    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ": ")
        sets := strings.Split(parts[1], "; ")
        minCubes := map[string]int{"red": 0, "green": 0, "blue": 0}
        for _, set := range sets {
            cubes := strings.Split(set, ", ")
            for _, cube := range cubes {
                colorCount := strings.Split(cube, " ")
                count, _ := strconv.Atoi(colorCount[0])
                color := colorCount[1]
                if count > minCubes[color] {
                    minCubes[color] = count
                }
            }
        }
        sum += minCubes["red"] * minCubes["green"] * minCubes["blue"]
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    // The sum of the powers of the minimum sets of cubes is now in the variable sum
    fmt.Println(sum)
}