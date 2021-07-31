//question link https://codeforces.com/problemset/problem/71/A

package main

import (
	"fmt"
	"strconv"
)

func answer(input string) string {
	var mid string
	var o string
	if len(input) > 10 {
		mid = strconv.Itoa(len(input) - 2)
		o = string(input[0]) + mid + string(input[len(input)-1])
	} else {
		o = input
	}

	return o
}

func main() {

	var limit int
	var t string
	fmt.Scan(&limit)

	for i := 0; i < limit; i++ {
		fmt.Scan(&t)

		fmt.Println(answer(t))
	}

}
