package main

//https://codeforces.com/problemset/problem/263/A

import (
	"fmt"
	"math"
)

func main() {
	var in, p int

	for i := 0; i < 25; i++ {
		fmt.Scan(&in)
		if in == 1 {
			p = i
		}
	}

	pos, cen := p/5, p%5

	fmt.Println(int(math.Abs(float64(pos-2)) + math.Abs(float64(cen-2))))
}
