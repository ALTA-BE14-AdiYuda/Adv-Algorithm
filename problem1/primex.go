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
func Prime(number int) bool {
	if number < 2 {
		return false
	}
	for j := 2; j < number; j++ {
		if number%j == 0 {
			return false
		}

	}
	return true
}

// var hitungan = 0
// var res = 0
// for i := 1; i <= number*number; i++ {
// 	var habisBagi = 0
// 	for j := 1; j <= i; j++ {
// 		if i%j == 0 {
// 			habisBagi = habisBagi + 1
// 		}
// 	}
// 	if habisBagi == 2 {
// 		hitungan = hitungan + 1
// 		if hitungan == number {
// 			res = i
// 			break
// 		}
// 	}
// }
// return res

func main() {
	fmt.Println(PrimeX(1))  //2
	fmt.Println(PrimeX(5))  //11
	fmt.Println(PrimeX(8))  //19
	fmt.Println(PrimeX(9))  //23
	fmt.Println(PrimeX(10)) //29

}
