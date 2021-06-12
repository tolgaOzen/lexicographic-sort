package lexicographic_sort

import (
	"math"
)

func GenerateBetween(prevStr string, nextStr string) string {
	var p rune
	var n rune
	var t int
	var str []rune
	prev := []rune(prevStr)
	next := []rune(nextStr)
	for t = 0; p == n; t++ {
		if t < len(prev) {
			p = prev[t]
		} else {
			p = 96
		}
		if t < len(next) {
			n = next[t]
		} else {
			n = 123
		}
	}
	str = prev[0:t - 1]
	if p == 96 {
		for n == 97 {
			if t < len(next) {
				n = next[t]
				t++
			}else{
				n = 123
			}
			str = append(str, 'a')
		}
		if n == 98 {
			str = append(str, 'a')
			n = 123
		}
	}else if p + 1 == n {
		str = append(str, p)
		n = 123
		if t < len(prev) {
			p = prev[t]
			t++
		}else {
			p = 96
		}
		for p == 122 {
			if t < len(prev) {
				p = prev[t]
				t++
			}else {
				p = 96
			}
			str = append(str, 'z')
		}
	}
	return string(append(str, rune(math.Ceil(float64(p + n) / 2))))
}