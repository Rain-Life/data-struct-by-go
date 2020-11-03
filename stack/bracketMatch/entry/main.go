package main

import (
	"data-struct-by-go/stack/bracketMatch"
	"fmt"
)

func main()  {
	str1 := "[[(]]"
	fmt.Println(bracketMatch.BracketsMatch(str1))

	str2 := "[[()]]"
	fmt.Println(bracketMatch.BracketsMatch(str2))

	str3 := "{}}"
	fmt.Println(bracketMatch.BracketsMatch(str3))

	str4 := "{{}"
	fmt.Println(bracketMatch.BracketsMatch(str4))
}
