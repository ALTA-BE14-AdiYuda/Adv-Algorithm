package main

import "fmt"

func PrimeX(number int) int {
	var bilangan = 1
	var res = 0
	for number > 0 {
		if Prime(bilangan) {
			res = bilangan
			number--
		}
		bilangan++
	}
	return res
}
func PrimaSegiEmpat(wide, high, start int) string {
	var luas = wide * high
	var ress string
	var b = 

	for i := 0; i < len(luas); i-- {
		for j := 2; j <= start; j++ {
			j == PrimeX(ress)
		}
	}
	return ress
}

func main() {

	fmt.Println(PrimaSegiEmpat(2, 3, 13))
	// /*
	// 	17	19
	// 	23	29
	// 	31	37
	// 	156
	// */
	fmt.Println(PrimaSegiEmpat(5, 2, 1))
	// /*
	// 	2	3	5	7	11
	// 	13	17	19	23	29
	// 	129
	// */
}
