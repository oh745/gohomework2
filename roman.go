package main

import "fmt"

func main() {
	fmt.Println(convertToRoman(43))
}

var Roman = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
var Numeric = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}

func convertToRoman(number int) string {
	var roman string
	for {
		if number >= 100 {
			roman += Roman[8]
			number = number - Numeric[8]
		} else if number >= 90 {
			roman += Roman[7]
			number = number - Numeric[7]
		} else if number >= 50 {
			roman += Roman[6]
			number = number - Numeric[6]
		} else if number >= 40 {
			roman += Roman[5]
			number = number - Numeric[5]
		} else if number >= 10 {
			roman += Roman[4]
			number = number - Numeric[4]
		} else if number >= 9 {
			roman += Roman[3]
			number = number - Numeric[3]
		} else if number >= 5 {
			roman += Roman[2]
			number = number - Numeric[2]
		} else if number >= 4 {
			roman += Roman[1]
			number = number - Numeric[1]
		} else {
			roman += Roman[0]
			number = number - Numeric[0]
		}

		if number == 0 {
			break
		}
	}

	return roman
}
