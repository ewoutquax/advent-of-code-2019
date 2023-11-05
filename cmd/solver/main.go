package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-01-the-tyranny-of-the-rocket-equation"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-02-1202-program-alarm"
	"github.com/ewoutquax/advent-of-code-2019/pkg/register"
)

func main() {
	for _, puzzle := range getPuzzles() {
		register.ExecDay(puzzle)
	}
}

func getPuzzles() (puzzles []string) {
	var allPuzzles []string = register.GetAllDays()
	sort.Strings(allPuzzles)

	selection := readInput(fmt.Sprintf("Which puzzle to run (%s):\n", allPuzzles))
	switch selection {
	case "":
		latestPuzzle := allPuzzles[len(allPuzzles)-1]
		fmt.Printf("Running latest puzzle: %s\n\n", latestPuzzle)
		puzzles = []string{latestPuzzle}
	case "all":
		fmt.Printf("Running all puzzles\n\n")
		puzzles = allPuzzles
	default:
		fmt.Printf("Running selected puzzle: '%s'\n\n", selection)
		puzzles = []string{selection}
	}

	return
}

func readInput(question string) string {
	fmt.Printf("%s", question)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return strings.Trim(text, "\n")
}
