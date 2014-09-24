package main

import (
	"fmt"
	"math/rand"
)

func main() {
	a := make([]int, 101)

	for i, _ := range a {
		a[i] = rand.Intn(100)
	}

	fmt.Println(a)
	fmt.Println(sort(a))
}

func sort(a []int) []int {
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				swap(a, i, i+1)
				swapped = true
			}
		}
	}
	return a
}

func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}
