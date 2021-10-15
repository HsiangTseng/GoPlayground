package main

import (
	"fmt"
	"math"
)

func main() {
	a := 5.0
	b := 3.0
	fmt.Println(calArea(a), calArea(b))
}

func calArea(r float64) float64 {
	// 計算時型別要一致(Pi為float64，r就不能是int)
	return r * r * math.Pi
}
