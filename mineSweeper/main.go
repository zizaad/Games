// MineSweeper

// ! уровни сложности - easy, middle (?), hard
// ! генерация поля

// ! двумерный слайс структур cell: -1 - бомба, 0 - пустая ячейка, 1 - 8 - цифры, 9 - флаг; false - закрыто, true - открыто

// вывод поля (посмотреть, как чистить экран перед этим, чтобы поле всегда было в одном месте)
// ! поле с буквами и цифрами, как в шахматах, чтобы пользователь вводил определённую координату

// ввод пользователя: <оператор> <буква> <цифра> - операторы: show (открыть ячейку) и flag (поставить флаг)
// по координате открываю ячейку - если бомба, поле открывается и конец игры; если 0 - открывать соседние клетки
// ! цветные цифры, бомбы - ?, флаги - красные

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/inancgumus/screen"
)

const (
	easyLength = 10
	easyMines  = 25

	middleLength = 17
	middleMines  = 72

	hardLength = 26
	hardMines  = 169

	colorReset  = "\033[0m"
	colorRed    = "\033[31m" // -1	// 3
	colorGreen  = "\033[32m" // 2	// 7
	colorYellow = "\033[33m" // 4	// 9
	colorBlue   = "\033[34m" // 1	// 6
	colorPurple = "\033[35m" // 5	// 8
	colorCyan   = "\033[36m" // 9 - flag
	colorWhite  = "\033[37m"
)

type Cell struct {
	val  int
	open bool
}

func (c *Cell) String() string {
	if !c.open {
		return "▯"
	}
	var res string
	switch c.val {
	case -1:
		res = colorRed + "*"
	case 0:
		res = " "
	case 1:
		res = colorBlue + "1"
	case 2:
		res = colorGreen + "2"
	case 3:
		res = colorRed + "3"
	case 4:
		res = colorYellow + "4"
	case 5:
		res = colorPurple + "5"
	case 6:
		res = colorBlue + "6"
	case 7:
		res = colorGreen + "7"
	case 8:
		res = colorPurple + "8"
	case 9:
		res = colorCyan + "9"
	}
	res += colorReset
	return res
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var lvl int
CHOISE:
	fmt.Printf("Выберите уровень сложности:\n1. Easy\n2. Middle\n3. Hard\n")
	fmt.Scan(&lvl)
	var field [][]Cell = nil
	switch lvl {
	case 1:
		field = generation(field, easyLength, easyMines)
	case 2:
		field = generation(field, middleLength, middleMines)
	case 3:
		field = generation(field, hardLength, hardMines)
	default:
		fmt.Printf("Нет такого варианта\n")
		goto CHOISE
	}
	showField(field)
}

func generation(field [][]Cell, length, mines int) [][]Cell {
	field = make([][]Cell, length)
	for idx := range field {
		field[idx] = make([]Cell, length)
	}

	// генерируем mines различных чисел - координаты мин
	for i := 0; i < mines; i++ {
		for {
			crd := rand.Intn(length * length)
			if field[crd/length][crd%length].val != -1 {
				field[crd/length][crd%length].val = -1
				break
			}
		}
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if field[i][j].val == -1 {
				continue
			}
			for k := -1; k <= 1; k++ {
				if i+k == -1 || i+k == length {
					continue
				}
				for l := -1; l <= 1; l++ {
					if j+l == -1 || j+l == length || (k == 0 && l == 0) {
						continue
					}
					if field[i+k][j+l].val == -1 {
						field[i][j].val++
					}
				}
			}
		}
	}

	return field
}

func showField(field [][]Cell) {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Print("  ")
	for i := 0; i < len(field); i++ {
		fmt.Printf("%-2d", i+1)
		if i != len(field)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
	var letter int = 'a'
	for i := 0; i < len(field); i++ {
		fmt.Printf("%c ", letter)
		letter++
		for j := 0; j < len(field); j++ {
			fmt.Print(field[i][j].String())
			if j != len(field)-1 {
				fmt.Print("  ")
			}
		}
		if i != len(field)-1 {
			fmt.Print("\n")
		}
	}
}
