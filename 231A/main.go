package main

//https://codeforces.com/problemset/problem/231/A

import "fmt"

func main() {
	var n int
	var t int
	var agree int
	agree = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if a == 1 {
			agree += 1
		}

		if b == 1 {
			agree += 1
		}

		if c == 1 {
			agree += 1
		}

		if agree >= 2 {
			t++
		}

		agree = 0

	}

	fmt.Println(t)
}
