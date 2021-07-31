package main

//https://codeforces.com/problemset/problem/158/A

import "fmt"

func main() {
	var n, k, input, a, p int
	a = 0
	p = 0
	fmt.Scan(&n, &k)
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		if i == k-1 {
			p = input
		}
		if input >= p && input != 0 {
			a++
		}
	}
	fmt.Println(a)
}
