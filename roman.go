package main

import "strings"



var romans = map[string]int{
    "M" : 1000,
     "CM" : 900,
     "D" : 500,
    "CD" : 400,
    "C" : 100,
    "XC" : 90,
     "L" : 50,
		"XL" : 40,
     "X" : 10,
     "IX" : 9,
     "V" : 5,
    "IV" : 4,
     "I" : 1,
}



func Tointeger (roman string) int {
	var val int
	s := strings.Split(roman, "")
	for n := 0; n <= (len(s) - 1); n++ {
		 if n == (len(s) - 1) {
			val = val + romans[s[n]]
		} else if romans[s[n]] < romans[s[n+1]] {
				val = val + romans[(s[n] + s[n+1])]
				n += 1
			} else {
				val = val + romans[s[n]]
			}
		}
	return val
}
