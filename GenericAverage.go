package main

import "fmt"

type NumberType int

func average(numbers ...NumberType) float64 {
	total := NumberType(0)
	for _, number := range numbers {
		total += number
	}
	return float64(total) / float64(len(numbers))
}

func main() {
	fmt.Println(average(1, 2, 3))
	fmt.Println(average(1, 2, 3, 4, 5))
}
