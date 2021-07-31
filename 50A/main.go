package main

//https://codeforces.com/problemset/problem/50/A

import "fmt"

func main() {
	var m, n int

	fmt.Scan(&m, &n)

	board := m * n

	number := board / 2

	fmt.Println(number)
}
