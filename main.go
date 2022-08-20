package main

import "fmt"

func main() {
	var r float32
	fmt.Println("enter radius")
	fmt.Scanln(&r)
	fmt.Println(diameter(r))
	fmt.Println(circumference(r))
	fmt.Println(area(r))
}

func diameter(r float32) float32 {
	var result float32
	result = (2 * r)
	return result
}
func circumference(r float32) float32 {
	var result1 float32
	result1 = (2 * 3.14 * r)
	return result1
}
func area(r float32) float32 {
	var result2 float32
	result2 = (3.14 * r * r)
	return result2
}
