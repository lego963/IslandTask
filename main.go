package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var currentIsland = 0

func main() {
	file, err := os.Open("island_data.txt")
	if err != nil {
		log.Fatalf("can't open file %s", err)
	}

	scanner := bufio.NewScanner(file)

	var island [][]int

	for scanner.Scan() {
		line := scanner.Text()
		islandLine := strings.Split(line, " ")
		islandData := make([]int, len(islandLine))

		for i, v := range islandLine {
			value, err := strconv.Atoi(v)
			if err != nil {
				log.Fatalf("can't convert character to int %s", err)
			}
			islandData[i] = value
		}
		island = append(island, islandData)
	}

	printIsland(island)

	var islands []int

	for i := 0; i < len(island); i++ {
		for j := 0; j < len(island[i]); j++ {
			if island[i][j] != 0 {
				currentIsland = 0
				makeStep(island, i, j)
				islands = append(islands, currentIsland)
			}
		}
	}

	maxIsland := math.MinInt
	for _, v := range islands {
		if v > maxIsland {
			maxIsland = v
		}
	}
	fmt.Printf("Max island is %d\n", maxIsland)
}

func printIsland(island [][]int) {
	for _, line := range island {
		for _, elm := range line {
			fmt.Print(fmt.Sprintf("%d ", elm))
		}
		fmt.Println()
	}
}

// i [0..N]
// j [0..M]
func makeStep(island [][]int, i, j int) {
	if island[i][j] == 0 {
		return
	}
	currentIsland++
	island[i][j] = 0
	if i-1 >= 0 {
		makeStep(island, i-1, j)
	}
	if j+1 < len(island[i]) {
		makeStep(island, i, j+1)
	}
	if i+1 < len(island) {
		makeStep(island, i+1, j)
	}
	if j-1 >= 0 {
		makeStep(island, i, j-1)
	}
}
