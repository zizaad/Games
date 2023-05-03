// Быки и коровы
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	colorReset  = "\033[0m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"

	YELLOW = 1
	GREEN  = 2
)

func main() {
	try := 0
	var digits int
	fmt.Println("Введите количество цифр")
	fmt.Scan(&digits)
	rand.Seed(time.Now().UnixNano())
	src := make([]int, digits)
	for i := 0; i < digits; i++ {
		src[i] = rand.Intn(10)
	}
	mp := make(map[int]int, 10)
	usrWrd := make([]int, digits)
	for {
		fillMap(mp, src)
		fmt.Printf("Введите число\n")
		var usr string
		fmt.Scan(&usr)
		if len(usr) != digits {
			fmt.Printf("Неправильное количество цифр!\n")
			continue
		}
		parse(usrWrd, usr)
		try++
		printWrd(src, mp, usrWrd)
		if win(src, usrWrd) {
			fmt.Printf("Вы отгадали число за %d попыток!\n", try)
			break
		}
	}
}

func fillMap(mp map[int]int, sl []int) {
	for idx := range mp {
		mp[idx] = 0
	}
	for _, val := range sl {
		mp[val]++
	}
}

func parse(sl []int, wrd string) {
	for i := 0; i < len(sl); i++ {
		sl[i] = int(wrd[i] - '0')
	}
}

func printWrd(src []int, mp map[int]int, usr []int) {
	color := make([]int, len(src))
	for idx, val := range src {
		if val == usr[idx] {
			color[idx] = GREEN
			mp[val]--
		}
	}
	for idx, val := range usr {
		if color[idx] != 0 {
			continue
		}
		if mp[val] != 0 {
			color[idx] = YELLOW
			mp[val]--
		}
	}
	for idx, val := range usr {
		if color[idx] == YELLOW {
			fmt.Print(colorYellow)
		} else if color[idx] == GREEN {
			fmt.Print(colorGreen)
		} else {
			fmt.Print(colorReset)
		}
		fmt.Print(val, colorReset)
	}
	fmt.Printf("\n")
}

func win(src []int, usr []int) bool {
	for idx, val := range src {
		if val != usr[idx] {
			return false
		}
	}
	return true
}
