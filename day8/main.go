package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"math"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	lrpattern := scanner.Text()
	fmt.Println(lrpattern)
	scanner.Scan()
	scanner.Text()

	// make a map[string][]string
	paths := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, " ", "")
		line2 := strings.ReplaceAll(line, "(", "")
		line3 := strings.Trim(line2, ")")
		strs := strings.Split(line3, "=")
		start := strs[0]
		outs := strings.Split(strs[1], ",")
		end := []string{outs[0], outs[1]}
		paths[start] = end
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	startNodes := findANodes(paths)
	fmt.Printf("startNodes: %v\n", startNodes)
	numSteps := make([]int, len(startNodes))
	for idx, node := range startNodes {
		fmt.Println(node)
		numSteps[idx] = findSteps(paths, node, lrpattern)
	}
	fmt.Println(numSteps)

	lcm := findLCM(numSteps)
	fmt.Println(lcm)
}

func findANodes(paths map[string][]string) []string {
	nodes := []string{}
	for k, _ := range paths {
		if k[2] == 'A' {
			nodes = append(nodes, k)
		}
	}
	return nodes
}

func findSteps(paths map[string][]string, start string, lrPattern string) int {
	lrOpts := paths[start]
	lrIndex := 0
	cnt := 0

	for {
		if lrIndex == len(lrPattern) {
			lrIndex = 0
		}
		if lrPattern[lrIndex] == 'L' {
			if lrOpts[0][2] == 'Z' {
				fmt.Println(cnt)
				return cnt + 1
			}
			lrOpts = paths[lrOpts[0]]
		} else {
			if lrOpts[1][2] == 'Z' {
				fmt.Println(cnt)
				return cnt + 1
			}
			lrOpts = paths[lrOpts[1]]
		}
		lrIndex++
		cnt++
	}
}

// copilot wrote me this set of functions to find the LCM of the list of ints
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

func lcm(a, b int) int {
    return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func findLCM(nums []int) int {
    result := nums[0]
    for i := 1; i < len(nums); i++ {
        result = lcm(result, nums[i])
    }
    return result
}