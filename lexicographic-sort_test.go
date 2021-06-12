package lexicographic_sort

import (
	"fmt"
	"testing"
)

func TestGenerateBetween(t *testing.T) {
	fmt.Println(GenerateBetween("abcde", "abchi"))
	fmt.Println(GenerateBetween("abc", "abchi"))
	fmt.Println("-")
	fmt.Println(GenerateBetween("abhs", "abit"))
	fmt.Println(GenerateBetween("abh", "abit"))
	fmt.Println("-")
	fmt.Println(GenerateBetween("abhz", "abit"))
	fmt.Println(GenerateBetween("abhzs", "abit"))
	fmt.Println(GenerateBetween("abhzz", "abit"))
	fmt.Println("-")
	fmt.Println(GenerateBetween("abc", "abcah"))
	fmt.Println(GenerateBetween("abc", "abcab"))
	fmt.Println(GenerateBetween("abc", "abcaah"))
	fmt.Println(GenerateBetween("abc", "abcb"))
}
