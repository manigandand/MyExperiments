// Function values
/*
Functions are values too. They can be passed around just like other values.

Function values may be used as function arguments and return values.
*/

package main

import (
	"fmt"
	"math"
)

//get func as argument
func FunGetCall(funRef func(int, int) int) int {
	return funRef(5, 4)
}
func FunGetCall1(funRef func(float64, float64) float64) float64 {
	return funRef(5, 4)
}
func main() {
	funPass := func(x, y int) int {
		return x * y * 2
	}
	// get the value from outside function, by passing
	// funPass as argument
	result := FunGetCall(funPass)
	fmt.Println(result)
	fmt.Println(FunGetCall1(math.Pow))
}
