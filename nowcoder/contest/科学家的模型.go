package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var pic [][]byte
	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < 5; i++ {
		sc.Scan()
		pic = append(pic, []byte(sc.Text()))
	}

	var startCol int
	for i := range pic {
		for j := range pic {
			if pic[i][j] == '1' {
				startCol = j
				break
			}
		}
	}

	switch {
	case pic[2][2] == '0':
		fmt.Println(0)
	case pic[3][startCol] == '0':
		fmt.Println(9)
	default:
		fmt.Println(8)
	}
}
