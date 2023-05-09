// MineSweeper

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
	easyMines  = 18

	middleLength = 17
	middleMines  = 72

	hardLength = 26
	hardMines  = 169
)

type Field struct {
	fld       [][]Cell
	mines     int
	minesOpen int
}

type Cell struct {
	val  int
	open bool
	flag bool
}

func (c *Cell) printColor() {
	white := color.New(color.FgHiWhite, color.BgHiBlack)
	blue := color.New(color.FgBlue, color.BgHiBlack)
	back := color.New(color.BgHiBlack)
	red := color.New(color.FgRed, color.BgHiBlack)
	green := color.New(color.FgGreen, color.BgHiBlack)
	magenta := color.New(color.FgMagenta, color.BgHiBlack)
	cyan := color.New(color.FgCyan, color.BgHiBlack)
	yellow := color.New(color.FgYellow, color.BgHiBlack)
	if c.flag {
		cyan.Print("⚑ ")
		return
	}
	if !c.open {
		white.Print("▮ ")
		return
	}
	switch c.val {
	case -1:
		red.Printf("%c", '*')
	case 0:
		back.Printf("%c", ' ')
	case 1:
		blue.Print(c.val)
	case 2:
		green.Print(c.val)
	case 3:
		red.Print(c.val)
	case 4:
		yellow.Print(c.val)
	case 5:
		magenta.Print(c.val)
	case 6:
		blue.Print(c.val)
	case 7:
		green.Print(c.val)
	case 8:
		magenta.Print(c.val)
	}
	back.Print(" ")
}

func main() {
	var lvl int
	green := color.New(color.FgGreen)
	fmt.Printf("Rules:\n")
	green.Printf("open <number> <letter>")
	fmt.Printf(" - open a cell with coords (number, letter)\n")
	green.Printf("flag <number> <letter>")
	fmt.Printf(" - set a flag on a cell with coords (number, letter)\n")
	green.Printf("exit")
	fmt.Printf(" - exit the game\n\n")
CHOISE:
	fmt.Printf("Choose the level of difficulty:\n1. Easy\n2. Middle\n3. Hard\n")
	fmt.Scan(&lvl)
	var field Field = Field{nil, 0, 0}
	switch lvl {
	case 1:
		field.generate(easyLength, easyMines)
	case 2:
		field.generate(middleLength, middleMines)
	case 3:
		field.generate(hardLength, hardMines)
	default:
		fmt.Printf("There is no such option\n")
		goto CHOISE
	}
	ex := true
	for ex {
		field.showField()
		if field.win() {
			color.Green("WIN")
			break
		}
		ex = field.cmd()
	}
}

func (f *Field) generate(length, mines int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	f.mines = mines
	f.fld = make([][]Cell, length)
	for idx := range f.fld {
		f.fld[idx] = make([]Cell, length)
	}

	for i := 0; i < mines; i++ {
		for {
			crd := r.Intn(length * length)
			if f.fld[crd/length][crd%length].val != -1 {
				f.fld[crd/length][crd%length].val = -1
				for k := -1; k <= 1; k++ {
					if crd/length+k == -1 || crd/length+k == length {
						continue
					}
					for l := -1; l <= 1; l++ {
						if crd%length+l == -1 || crd%length+l == length || (k == 0 && l == 0) {
							continue
						}
						if f.fld[crd/length+k][crd%length+l].val != -1 {
							f.fld[crd/length+k][crd%length+l].val++
						}
					}
				}
				break
			}
		}
	}

	// for i := 0; i < length; i++ {
	// 	for j := 0; j < length; j++ {
	// 		if f.fld[i][j].val == -1 {
	// 			continue
	// 		}
	// 		for k := -1; k <= 1; k++ {
	// 			if i+k == -1 || i+k == length {
	// 				continue
	// 			}
	// 			for l := -1; l <= 1; l++ {
	// 				if j+l == -1 || j+l == length || (k == 0 && l == 0) {
	// 					continue
	// 				}
	// 				if f.fld[i+k][j+l].val == -1 {
	// 					f.fld[i][j].val++
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}

func (f *Field) showField() {
	screen.Clear()
	screen.MoveTopLeft()
	grey := color.New(color.BgHiBlack, color.Bold)
	grey.Print("   ")
	var letter int = 'a'
	for i := 0; i < len(f.fld); i++ {
		grey.Printf("%-2c", letter)
		letter++
		if i != len(f.fld)-1 {
			grey.Print(" ")
		}
	}
	fmt.Print("\n")
	for i := 0; i < len(f.fld); i++ {
		grey.Printf("%2d ", i+1)
		for j := 0; j < len(f.fld); j++ {
			f.fld[i][j].printColor()
			if j != len(f.fld)-1 {
				grey.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Printf("Mines: %d\n", f.mines-f.minesOpen)
}

func (f *Field) cmd() bool {
	var op string
	fmt.Scan(&op)
	switch op {
	case "open":
		var i int
		var j string
		fmt.Scan(&i)
		fmt.Scan(&j)
		mine := openCell(f.fld, i-1, int(j[0]-97))
		if mine {
			f.openCells()
			color.Red("DEFEAT")
			return false
		}
	case "flag":
		var i int
		var j string
		fmt.Scan(&i)
		fmt.Scan(&j)
		if f.fld[i-1][int(j[0]-97)].open {
			return true
		}
		f.fld[i-1][int(j[0]-97)].flag = !f.fld[i-1][int(j[0]-97)].flag
		f.minesOpen++
	case "exit":
		f.openCells()
		return false
	}
	return true
}

func openCell(field [][]Cell, i, j int) bool {
	field[i][j].open = true
	if field[i][j].val == 0 {
		autoOpen(field, i, j, 0, -1) // left
		autoOpen(field, i, j, 0, 1)  // right
		autoOpen(field, i, j, -1, 0) // top
		autoOpen(field, i, j, 1, 0)  // bottom
	}
	return field[i][j].val == -1
}

func autoOpen(field [][]Cell, i, j, di, dj int) {
	if di == 0 {
		if j+dj == -1 || j+dj == len(field) {
			return
		}
		for k := -1; k <= 1; k++ {
			if i+k == -1 || i+k == len(field) {
				continue
			}
			if field[i+k][j+dj].open {
				continue
			}
			field[i+k][j+dj].open = true
			if field[i+k][j+dj].val == 0 {
				autoOpen(field, i+k, j+dj, di, dj)
				autoOpen(field, i+k, j+dj, k, di)
			}
		}
	} else {
		if i+di == -1 || i+di == len(field) {
			return
		}
		for k := -1; k <= 1; k++ {
			if j+k == -1 || j+k == len(field) {
				continue
			}
			if field[i+di][j+k].open {
				continue
			}
			field[i+di][j+k].open = true
			if field[i+di][j+k].val == 0 {
				autoOpen(field, i+di, j+k, di, dj)
				autoOpen(field, i+di, j+k, dj, k)
			}
		}
	}
}

func (f *Field) openCells() {
	for i := range f.fld {
		for j := range f.fld[i] {
			f.fld[i][j].open = true
		}
	}
	f.showField()
}

func (f *Field) win() bool {
	for i := range f.fld {
		for j := range f.fld[i] {
			if !f.fld[i][j].open && f.fld[i][j].val != -1 {
				return false
			}
		}
	}
	return true
}
