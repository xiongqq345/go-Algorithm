package main

import (
	"bufio"
	"fmt"
	"os"
)

//We have two figures S and T on a two-dimensional grid with square cells.
//
//S lies within a grid with N rows and N columns, and consists of the cells where S
//i,j
//​
//  is #.
//T lies within the same grid with N rows and N columns, and consists of the cells where T
//i,j
//​
//  is #.
//
//Determine whether it is possible to exactly match S and T by 90-degree rotations and translations.

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	s := make([][]byte, n)
	t := make([][]byte, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		s[i] = []byte(sc.Text())
	}

	var rpos [][2]int
	first := [2]int{-1, -1}
	for i := 0; i < n; i++ {
		sc.Scan()
		t[i] = []byte(sc.Text())
		for j, ch := range t[i] {
			if ch == '#' {
				if first[0] == -1 {
					first[0], first[1] = i, j
				} else {
					rpos = append(rpos, [2]int{i - first[0], j - first[1]})
				}
			}
		}
	}

	if n == 1 {
		if s[0][0] == t[0][0] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}

	for r := 0; r < 4; r++ {
		sfirst := [2]int{-1, -1}
		idx := 0
		same := true
	p:
		for i := 0; i < n; i++ {
			for j, ch := range s[i] {
				if ch == '#' {
					if sfirst[0] == -1 {
						sfirst[0], sfirst[1] = i, j
					} else {
						if idx >= len(rpos) || rpos[idx][0] != i-sfirst[0] || rpos[idx][1] != j-sfirst[1] {
							same = false
							break p
						}
						idx++
					}
				}
			}
		}
		if same {
			fmt.Println("Yes")
			return
		}
		rotate(s)
	}

	fmt.Println("No")
}

func rotate(matrix [][]byte) {
	n := len(matrix[0])
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-i-1] = matrix[n-i-1], matrix[i]
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
