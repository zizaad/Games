// MineSweeper

// ! уровни сложности - easy, middle (?), hard
// ! генерация поля

// ! двумерный слайс структур cell: -1 - бомба, 0 - пустая ячейка, 1 - 8 - цифры, 9 - флаг; false - закрыто, true - открыто

// ! вывод поля (посмотреть, как чистить экран перед этим, чтобы поле всегда было в одном месте)
// ! поле с буквами и цифрами, как в шахматах, чтобы пользователь вводил определённую координату

// ! ввод пользователя: <оператор> <буква> <цифра> - операторы: open (открыть ячейку) и flag (поставить флаг); exit
// по координате открываю ячейку - если бомба, поле открывается и конец игры; если 0 - открывать соседние клетки
// ! цветные цифры, бомбы - ?, флаги - красные

// добавить цвета - escape последовательности (после [ поставить 90+)

// написать правила: open - открыть клетку, flag - поставить флаг
// показать открытое поле после exit
// проверка победы

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

const (
	easyLength = 10
	easyMines  = 25

	middleLength = 17
	middleMines  = 72

	hardLength = 26
	hardMines  = 169

	colorReset        = "\033[0m"
	colorRed          = "\033[31m"  // -1	// 3
	colorBrightRed    = "\033[91m]" // 9
	colorGreen        = "\033[32m"  // 2	// 7
	colorBrightGreen  = "\033[92m]"
	colorYellow       = "\033[33m" // 4	// 9
	colorBrightYellow = "\033[93m]"
	colorBlue         = "\033[34m" // 1	// 6
	colorBrightBlue   = "\033[94m]"
	colorPurple       = "\033[35m" // 5	// 8
	colorBrightPurple = "\033[95m]"
	colorCyan         = "\033[36m" // 9 - flag
	colorBrightCyan   = "\033[96m]"
	colorWhite        = "\033[37m"
	colorBrightBlack  = "\033[90m"
)

type Cell struct {
	val  int
	open bool
}

func (c *Cell) String() string {
	if !c.open && c.val != 9 {
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
		res = colorCyan + "⚑"
	}
	res += colorReset
	return res
}

func main() {
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
	ex := true
	for ex {
		showField(field)
		ex = cmd(field)
	}
}

func generation(field [][]Cell, length, mines int) [][]Cell {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	field = make([][]Cell, length)
	for idx := range field {
		field[idx] = make([]Cell, length)
	}

	// генерируем mines различных чисел - координаты мин
	for i := 0; i < mines; i++ {
		for {
			crd := r.Intn(length * length)
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
	fmt.Print("   ")
	grey := color.New(color.FgHiBlack, color.Bold)
	var letter int = 'a'
	for i := 0; i < len(field); i++ {
		// fmt.Printf("%s%-2c%s", colorBrightBlack, letter, colorReset)
		grey.Printf("%-2c", letter)
		letter++
		if i != len(field)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
	for i := 0; i < len(field); i++ {
		fmt.Printf("%s%-2d%s ", colorBrightBlack, i+1, colorReset)
		for j := 0; j < len(field); j++ {
			fmt.Print(field[i][j].String())
			if j != len(field)-1 {
				fmt.Print("  ")
			}
		}
		fmt.Print("\n")
	}
}

func cmd(field [][]Cell) bool {
	var op string
	fmt.Scan(&op)
	switch op {
	case "open":
		var i int
		var j string
		fmt.Scan(&i)
		fmt.Scan(&j)
		openCell(field, i-1, int(j[0]-97))
	case "flag":
		var i int
		var j string
		fmt.Scan(&i)
		fmt.Scan(&j)
		field[i-1][int(j[0]-97)].val = 9
	case "exit":
		// показать открытое поле
		return false
	}
	return true
}

func openCell(field [][]Cell, icur, jcur int) {
	field[icur][jcur].open = true
}
