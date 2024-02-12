package tetris

import (
	"fmt"
	"sync"
)

func Solver(gridSize int) {
	var wg sync.WaitGroup
	wg.Add(2)

	//first routine that try to place the pieces in the order they appear in the file
	grid := gridBuilder(gridSize)
	gridCopy := make([]string, len(grid))
	copy(gridCopy, grid)
	go RoutineOrder(grid, gridSize, gridCopy, &wg)

	//second routine that try to place the pieces in reverse order
	gridR := make([]string, len(grid))
	gridCopyR := make([]string, len(grid))
	copy(gridR, grid)
	gridSizeR := gridSize
	copy(gridCopyR, gridCopy)
	go RoutineReverseOrder(gridR, gridSizeR, gridCopyR, &wg)

	//we wait until we find a solution
	wg.Wait()
	for _, subV := range solution[0] {
		fmt.Println(subV)
	}
	fmt.Println()
}

func RoutineOrder(grid []string, gridSize int, gridCopy []string, wg *sync.WaitGroup) {
	defer wg.Done()
	tetrosCopy := make([]Tetro, len(tetroslist))
	copy(tetrosCopy, tetroslist)
	solved := false
	for !solved && !found {
		solved, grid = Recurse(grid, tetrosCopy, gridCopy, solved)
		if !solved && !found {
			gridSize++
			grid = gridBuilder(gridSize)
		}
	}
}

func RoutineReverseOrder(grid []string, gridSize int, gridCopy []string, wg *sync.WaitGroup) {
	defer wg.Done()
	var tetrosCopy []Tetro
	for i := len(tetroslist)-1; i >= 0; i-- {
		tetrosCopy = append(tetrosCopy, tetroslist[i])
	}
	solved := false
	for !solved && !found {
		solved, grid = Recurse(grid, tetrosCopy, gridCopy, solved)
		if !solved && !found {
			gridSize++
			grid = gridBuilder(gridSize)
		}
	}
}

func gridBuilder(gridSize int) []string {
	/* build a grid based on the formula square root of tetromis * 4
	is also called when we want to expend the grid */
	strTemp := ""
	grid := []string{}
	for i := gridSize; i > 0; i-- { //build the x
		strTemp += "."
	}
	for i := gridSize; i > 0; i-- { //build the y
		grid = append(grid, strTemp)
	}
	return grid
}

func Recurse(grid []string, tetcopy []Tetro, gridCopy []string, solved bool) (bool, []string) {

	problem := false
	if len(tetcopy) == 0 {
		solved = true
		return solved, grid
	}
	beforeChangeGrid := make([]string, len(grid))
	copy(beforeChangeGrid, grid)
	for i, v := range grid { // for y
		for subI, subV := range v {	//for x
			if subV == '.' { // if square is empty, try to put the pieces
				for subsubI := 0; subsubI < 4; subsubI++ {
					problem = false
					switch subsubI { // one case per part of the tetromino
					case 0:
						if grid[i][subI] == '.' {
							temp := ""
							for cursor, value := range grid[i] {
								if cursor == subI {
									temp += tetcopy[0].letter
								} else {
									temp += string(value)
								}
							}
							grid[i] = temp
						} else {
							problem = true
						}
					case 1:
						if i+tetcopy[0].tetro2[1] < len(grid) && subI+tetcopy[0].tetro2[0] < len(grid[0]) && i+tetcopy[0].tetro2[1] >= 0 && subI+tetcopy[0].tetro2[0] >= 0 {
							if grid[i+tetcopy[0].tetro2[1]][subI+tetcopy[0].tetro2[0]] == '.' {
								temp := ""
								for cursor, value := range grid[i+tetcopy[0].tetro2[1]] {
									if cursor == subI+tetcopy[0].tetro2[0] {
										temp += tetcopy[0].letter
									} else {
										temp += string(value)
									}
								}
								grid[i+tetcopy[0].tetro2[1]] = temp
							} else {
								problem = true
							}
						} else {
							problem = true
						}
					case 2:
						if i+tetcopy[0].tetro3[1] < len(grid) && subI+tetcopy[0].tetro3[0] < len(grid[0]) && i+tetcopy[0].tetro3[1] >= 0 && subI+tetcopy[0].tetro3[0] >= 0 {
							if grid[i+tetcopy[0].tetro3[1]][subI+tetcopy[0].tetro3[0]] == '.' {
								temp := ""
								for cursor, value := range grid[i+tetcopy[0].tetro3[1]] {
									if cursor == subI+tetcopy[0].tetro3[0] {
										temp += tetcopy[0].letter
									} else {
										temp += string(value)
									}
								}
								grid[i+tetcopy[0].tetro3[1]] = temp
							} else {
								problem = true
							}
						} else {
							problem = true
						}
					case 3:
						if i+tetcopy[0].tetro4[1] < len(grid) && subI+tetcopy[0].tetro4[0] < len(grid[0]) && i+tetcopy[0].tetro4[1] >= 0 && subI+tetcopy[0].tetro4[0] >= 0 {
							if grid[i+tetcopy[0].tetro4[1]][subI+tetcopy[0].tetro4[0]] == '.' {
								temp := ""
								for cursor, value := range grid[i+tetcopy[0].tetro4[1]] {
									if cursor == subI+tetcopy[0].tetro4[0] {
										temp += tetcopy[0].letter
									} else {
										temp += string(value)
									}
								}
								grid[i+tetcopy[0].tetro4[1]] = temp
							} else {
								problem = true
							}
						} else {
							problem = true
						}
					}
					if problem {
						break
					}
				}
				if !problem && len(tetcopy) > 1 { // if we manged to place the piece and there's more to place
					solved, grid = Recurse(grid, tetcopy[1:], gridCopy, solved)
					if !solved {
						copy(grid, beforeChangeGrid)
					}
				} else if len(tetcopy) == 1 && !problem { // if we manged to place the piece and if it was the last
					solved = true
					found = true
					solution = append(solution, grid)
				} else {
					copy(grid, beforeChangeGrid)
				}
			}
			if solved || found {
				break
			}
		}
		if solved || found {
			break
		}
	}
	return solved, grid
}
