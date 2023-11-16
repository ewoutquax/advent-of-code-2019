package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-01-the-tyranny-of-the-rocket-equation"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-02-1202-program-alarm"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-03-crossed-wires"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-04-secure-container"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-05-sunny-with-a-change-of-astroids"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-06-universal-orbit-map"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-07-amplification-circuit"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-08-space-image-format"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-09-sensor-boost"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-10-monitoring-station"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-13-care-package"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-14-space-stoichiometry"
	_ "github.com/ewoutquax/advent-of-code-2019/internal/day-15-oxygen-system"
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
	if len(os.Args) == 2 {
		return os.Args[1]
	}

	fmt.Printf("%s", question)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	return strings.Trim(text, "\n")
}
