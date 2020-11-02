package main

import (
	"data-struct-by-go/stack/arrayStack"
	"fmt"
)

func main()  {
	stack := arrayStack.ItemStack{N: 10}
	stack.Init()
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")
	stack.Push("d")
	for _, v := range stack.Items {
		fmt.Printf("%v\t", v)
	}
	fmt.Println()
	item := stack.Pop()
	fmt.Printf("%v\n", item)
}