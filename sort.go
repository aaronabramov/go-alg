package main

import "fmt"
import "math/rand"

func main() {
	a := make([]int, 101)

	for i, _ := range a {
		a[i] = rand.Intn(100)
	}

	fmt.Println(sort(a))
}

func sort(a []int) []int {
	l := make([]int, 0)
	r := make([]int, 0)
	n := len(a) / 2

	if len(a) <= 1 {
		return a
	}

	for i, num := range(a) {
		if i < n {
			l = append(l, num)
		} else {
			r = append(r, num)
		}
	}

	l = sort(l);
	r = sort(r);

	return merge(l, r)
}

func merge(l, r []int) []int {
	result := make([]int, 0)

	i := 0
	j := 0
	var num int

	for (i + j) <= (len(l) + len(r) - 1) {
		if i > len(l) - 1 {
			num = r[j]
			j++
		} else if j > len(r) - 1 {
			num = l[i]
			i++
		} else if l[i] < r[j] {
			num = l[i]
			i++
		} else {
			num = r[j]
			j++
		}
		result = append(result, num)
	}
	return result
}
