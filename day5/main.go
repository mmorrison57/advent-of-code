package main

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

func parseMapping(lines []string) map[int]int {
    mapping := make(map[int]int)
    for _, line := range lines {
        parts := strings.Split(line, " ")
        destStart, _ := strconv.Atoi(parts[0])
        srcStart, _ := strconv.Atoi(parts[1])
        length, _ := strconv.Atoi(parts[2])
        for i := 0; i < length; i++ {
            mapping[srcStart+i] = destStart + i
        }
    }
    return mapping
}

func findLocation(seed int, mappings []map[int]int) int {
    for _, mapping := range mappings {
        if val, ok := mapping[seed]; ok {
            seed = val
        }
    }
    return seed
}

func main() {
    file, _ := os.Open("input.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    seeds := strings.Split(scanner.Text(), " ")
    mappings := []map[int]int{}
    for scanner.Scan() {
        lines := []string{}
        for i := 0; i < 3; i++ {
            lines = append(lines, scanner.Text())
            scanner.Scan()
        }
        mappings = append(mappings, parseMapping(lines))
    }

	v, _ := strconv.Atoi(seeds[0])
    minLocation := findLocation(v, mappings)
    for _, seed := range seeds[1:] {
		z, _ := strconv.Atoi(seed)
        location := findLocation(z, mappings)
        if location < minLocation {
            minLocation = location
        }
    }

    println(minLocation)
}