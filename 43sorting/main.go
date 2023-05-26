package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strings", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("integers", ints)

	floats := []float64{10.3, 10.1, 10.2}
	sort.Float64s(floats)
	fmt.Println("floats", floats)

	s := sort.IntsAreSorted(ints)
	fmt.Println("isIntsSorted", s)
}