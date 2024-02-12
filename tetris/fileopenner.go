package tetris

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

func Openfile(file string) int {
	comming, _ := ioutil.ReadFile("./pieces/" + file)
	data := strings.Split(string(comming), "\n")
	CheckValidity(data)
	count := 0
	for _, v := range data {
		count += strings.Count(v, "#")
	}
	return int(math.Ceil(math.Sqrt(float64(count))))
}

func CheckValidity(entry []string) {
	count := 0
	for i := 0; i < len(entry); i += 5 {
		var tetro Tetro
		//check if each line of the tetro is 4 long
		if len(entry[i]) != 4 || len(entry[i+1]) != 4 || len(entry[i+2]) != 4 || len(entry[i+3]) != 4 {
			fmt.Println("Format error")
			os.Exit(1)
		}
		count = 0
		for subI := 0; subI <= 3; subI++ {
			// we count the number of # and add them in the tetro structure
			for subsubI, v := range entry[i+subI] {
				if v == '#' {
					count++
					if count == 1 {
						tetro.tetro1 = append(tetro.tetro1, subsubI)
						tetro.tetro1 = append(tetro.tetro1, subI)
					} else if count == 2 {
						tetro.tetro2 = append(tetro.tetro2, subsubI)
						tetro.tetro2 = append(tetro.tetro2, subI)
					} else if count == 3 {
						tetro.tetro3 = append(tetro.tetro3, subsubI)
						tetro.tetro3 = append(tetro.tetro3, subI)
					} else if count == 4 {
						tetro.tetro4 = append(tetro.tetro4, subsubI)
						tetro.tetro4 = append(tetro.tetro4, subI)
					}
				}
			}	
		}
		if count != 4 {
			fmt.Println("Format error")
			os.Exit(1)
		} else {
			tetroslist = append(tetroslist, tetro)
		}
	}
	letterAssign := 65
	for teti, v := range tetroslist { // we change the coordinates or the tetrominos so the upper left most one is equal to 0 0
		count = 0
		temp := [][]int{}
		toInsert := [][]int{}
		for i := 0; i <= 3; i++ {
			switch i {
			case 0:
				temp = append(temp, v.tetro1)
			case 1:
				temp = append(temp, v.tetro2)
			case 2:
				temp = append(temp, v.tetro3)
			case 3:
				temp = append(temp, v.tetro4)
			}
		}
		for i, subV := range temp {
			for subI, subsubV := range temp {
				if i != subI {
					if subV[0] == subsubV[0] && (subV[1] == subsubV[1] - 1 || subV[1] == subsubV[1] + 1) {
						count++
					} else if subV[1] == subsubV[1] && (subV[0] == subsubV[0] - 1 || subV[0] == subsubV[0] + 1) {
						count++
					}
				}
			}
			
		}
		if count < 6 { //if there's less than 6 links between the tetrominos pieces then it's not a real tetromino
			fmt.Println("Format error")
			os.Exit(1)
		}
		for i, subV := range temp {
			if i == 0 {
				toInsert = append(toInsert, []int{0, 0})
			} else {
				toInsert = append(toInsert, []int{subV[0]-temp[0][0], subV[1]-temp[0][1]})
			}
		}
		tetroslist[teti].tetro1 = toInsert[0]
		tetroslist[teti].tetro2 = toInsert[1]
		tetroslist[teti].tetro3 = toInsert[2]
		tetroslist[teti].tetro4 = toInsert[3]
		tetroslist[teti].letter = string(letterAssign)
		letterAssign++
	}
}
