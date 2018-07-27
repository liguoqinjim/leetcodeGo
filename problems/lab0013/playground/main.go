package main

import "log"

func main() {
	log.Println(romanToInt("III"))
}

func romanToInt(s string) int {
	//Symbol       Value
	//I             1
	//V             5
	//X             10
	//L             50
	//C             100
	//D             500
	//M             1000

	//m := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	for n, v := range s {
		log.Println(n, v)
	}

	return 0
}
