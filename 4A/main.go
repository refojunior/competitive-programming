//question link https://codeforces.com/problemset/problem/4/A

package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%v", &a)
	if a%2 == 0 {
		if a == 2 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	} else {
		fmt.Println("NO")
	}
}
