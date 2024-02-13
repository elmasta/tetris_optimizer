# Tetris Optimizer

### Description
_______
Program that oraganise tetrominos pieces taken fro a text file in the smallest rectangle possible.

### Example
_______
for a file containing:
```
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

....
.##.
.##.
....

....
....
##..
.##.

##..
.#..
.#..
....

....
###.
.#..
....
```

the programme should print:
```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

### Usage
_______
```go
go run . <text file>
```
