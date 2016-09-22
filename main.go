package main

import "C"

var counter = 10

//export DoubleIt
func DoubleIt(x int) int {
	return x * 2
}

//export Count
func Count() (res int) {
	res = counter
	counter++
	return
}

func main() {
}
