package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := []string{"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892"}
	fmt.Println(saveMobileNumber(s))
}

func saveMobileNumber(phoneNumber []string) map[string]int {
	re := regexp.MustCompile(`[^0-9]`)
	m := make(map[string]int)
	for _, number := range phoneNumber {
		n := re.ReplaceAllString(number, "")
		var count, exist = m[n]
		if exist {
			m[n] = count + 1
		} else {
			m[n] = 1
		}
	}

	return m
}
