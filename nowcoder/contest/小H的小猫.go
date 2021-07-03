package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)
	var x, y float64
	xmin, ymin := big.NewFloat(math.MaxFloat32), big.NewFloat(math.MaxFloat32)
	for n > 0 {
		fmt.Scan(&x, &y)
		bigx := new(big.Float).SetFloat64(x).GobEncode()
		yminf, _ := ymin.Float64()
		xminf, _ := xmin.Float64()
		if x == 0 && y < yminf {
			ymin.SetFloat64(y)
		}
		if y == 0 && x < xminf {
			xmin.SetFloat64(x)
		}
		n--
	}

	yminf, _ := ymin.Float64()
	xminf, _ := xmin.Float64()
	if xminf == math.MaxFloat32 || yminf == math.MaxFloat32 {
		fmt.Println("Poor Little H!")
		return
	}
	zz := new(big.Float)
	xx := xmin.Mul(xmin, xmin)
	yy := ymin.Mul(ymin, ymin)
	zz.Add(xx, yy)
	ans, _ := zz.Sqrt(zz).Float64()
	fmt.Printf("%.10f\n", ans)
}
