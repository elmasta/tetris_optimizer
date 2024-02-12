package tetris

var tetroslist []Tetro
var solution [][]string
var found bool

type Tetro struct {
	tetro1 []int
	tetro2 []int
	tetro3 []int
	tetro4 []int
	letter string
}
