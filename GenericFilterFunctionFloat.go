package main

import (
	"fmt"
)

func Filter(slice []interface{}, f func(interface{}) bool) []interface{} {
	filtered := make([]interface{}, 0)
	for _, v := range slice {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	numbers := []float64{5.12, 15.0, 3.39, 8.0, 9.0, 14.0, 6.63}
	numbersInterface := make([]interface{}, len(numbers))
	for i, v := range numbers {
		numbersInterface[i] = v
	}

	noDecimals := Filter(numbersInterface, func(v interface{}) bool {
		return v.(float64) == float64(int(v.(float64)))
	})
	fmt.Println("Numbers which do not have any digits after the decimal:")
	fmt.Println(noDecimals)

	divisibleByThree := Filter(numbersInterface, func(v interface{}) bool {
		return v.(float64)/3.0 == float64(int(v.(float64)/3.0))
	})
	fmt.Println("Numbers which are divisible by 3.0:")
	fmt.Println(divisibleByThree)
}
