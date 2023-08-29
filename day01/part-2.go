package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var elves []int
	calories := 0

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		// Current elf's inventory is complete on empty line
		if text == "" {
			elves = append(elves, calories)
			calories = 0
			continue
		}

		item_calories, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		calories += item_calories
	}

	// Still have the last elf's calories stored if no newline before EOF
	if calories != 0 {
		elves = append(elves, calories)
	}

	largest := 0
	for _, calories := range elves {
		// fmt.Printf("Elf #%v is carrying %v calories\n", i+1, calories)
		if calories > largest {
			largest = calories
		}
	}

	// Sort in reverse (such that slice[0] is largest value)
	slices.SortFunc(elves, func(x, y int) int {
		return y - x
	})
	// fmt.Println()
	fmt.Println("Top 3 elves are carrying", elves[0] + elves[1] + elves[2], "calories")

	file.Close()
}

