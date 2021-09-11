package problem

import "fmt"

func main() {
	var w int
	fmt.Scan(&w)
	if w > 3 && w&1 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
