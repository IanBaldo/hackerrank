// https://www.hackerrank.com/contests/hourrank-23/challenges/selective-additions
package main

import "fmt"

func main() {
	nmk, err := intScanln(3)
	checkError(err)

	values, err := intScanln(nmk[0])
	checkError(err)

	favNumbers, err := intScanln(nmk[2])
	checkError(err)

	oper := make(map[int][]int)
	for i:=0; i< nmk[1]; i++ {

		oper[i], err = intScanln(3)
		checkError(err)
	}

	for _, op := range oper {

		start := op[0]-1
		end := op[1]
		a := values[start:end]
		addValue2Slice(a, op[2], favNumbers)
		fmt.Println(sumArray(values))
	}

}

// Function to add the specified value to all non-favourite numbers in a slice
func addValue2Slice (w []int, val int, favs []int) []int {
	for i,num := range w {
		if !sliceContains(favs, num) {
			w[i] += val
		}
	}

	return w
}

// Simple function to search for a value in a slice
func sliceContains (inputSlice []int, val int) bool {
	
	for _, n := range inputSlice {
		if n == val {
			return true
		}
	}

	return false
}

// Returns the total sum of the content an integer slice
func sumArray (in []int) int {
	var sum int
	for _,val := range in {
		sum += val
	}
	return sum
}

// Scans n integers in a line. Returns an integer slice with those.
func intScanln (n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))

	for i := range x {
		y[i] = &x[i]
	}

	n, err := fmt.Scanln(y...)
	x = x[:n]

	return x, err
}

func checkError (err error) {
	if err != nil {
		panic(err)
	}
}