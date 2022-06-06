package main

import (
	"fmt"
)

// purpose:
// test sth about slice
func main() {
	arr := [5]int{0,1,2,3,4}
	sl := arr[0: 3]
	pr5(arr)
	pr(sl)
	arr[1] = 2
	pr5(arr)
	pr(sl)
	sl = append(sl, 5, 6, 7, 8)
	pr5(arr)
	pr(sl)
}

func pr(a []int) {
	n := len(a)
	fmt.Printf("[")
	for i:= 0; i < n; i++ {
		fmt.Printf("%d", a[i])
		if i != n-1 {
			fmt.Printf(",")
		}
	}
	fmt.Printf("]\n")
}

func pr5(a [5]int) {
	n := len(a)
	fmt.Printf("[")
	for i:= 0; i < n; i++ {
		fmt.Printf("%d", a[i])
		if i != n-1 {
			fmt.Printf(",")
		}
	}
	fmt.Printf("]\n")
}