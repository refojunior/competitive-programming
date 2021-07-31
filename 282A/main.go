package main

//https://codeforces.com/problemset/problem/282/A

import "fmt"

func main() {
	var n, m int
	var o string
	m = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&o)
		if o == "X++" || o == "++X" {
			m++
		} else {
			m--
		}
	}
	fmt.Println(m)
}
