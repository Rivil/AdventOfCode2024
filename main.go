package main

import (
	puzzle1 "AdventOfCode2024/puzzle1"
	"fmt"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	fmt.Fprintln(os.Stdout, "Day 1 Puzzle 1: %i", puzzle1.Puzzle1())
	fmt.Fprintln(os.Stdout, "Day 1 Puzzle 2: %i", puzzle1.Puzzle2())
}
