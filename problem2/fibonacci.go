package main

import "fmt"

func Fibonacci(number int) int {
	var (
		listFibo = map[int]int{}
	)
	if number < 2 {
		listFibo[number] = number
		return number
	}
	if val, found := listFibo[number]; found {
		return val
	} else {
		listFibo[number] = Fibonacci(number-2) + Fibonacci(number-1)
	}
	return listFibo[number]
}

func main() {

	fmt.Println(Fibonacci(0))  // 0
	fmt.Println(Fibonacci(2))  // 1
	fmt.Println(Fibonacci(9))  // 34
	fmt.Println(Fibonacci(10)) // 55
	fmt.Println(Fibonacci(12)) // 144

}
