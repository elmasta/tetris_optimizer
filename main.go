package main

import (
	"os"
	"fmt"
	"time"
	tetris "tetris/tetris"
)

func main() {
	start := time.Now()
	if len(os.Args) != 2 {
		fmt.Println("Not enough or too many arguments.")
		os.Exit(2)
	}
	gridSize := tetris.Openfile(os.Args[1])
	tetris.Solver(gridSize)
	elapsed := time.Since(start)
    fmt.Printf("The program took %s to solve it.\n", elapsed)
}
