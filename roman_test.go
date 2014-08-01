package main

import "testing"

type testpair struct {
	roman string
	value int
}

var tests = []testpair {
	{"M", 1000 },
	{ "CM", 900 },
	{"D", 500 },
	{ "CD", 400 },
	{"C", 100 },
	{ "XC", 90 },
	{ "L", 50 },
	{ "XL", 40 },
	{ "X", 10 },
	{ "IX", 9 },
	{ "V", 5 },
	{ "IV", 4 },
	{ "I", 1 },
	{ "MMMCMXCIX", 3999 },
	{ "MMMD", 3500 },

}

func TestRoman(t *testing.T) {
	for _, pair := range tests {
		v := Tointeger(pair.roman)
			if v != pair.value {
				t.Error(
					"For", pair.roman,
					"expected", pair.value,
					"got", v,
				)
			}
		}
	}
