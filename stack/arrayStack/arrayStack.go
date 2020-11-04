package arrayStack

import "fmt"

type Item interface {}

type ItemStack struct {
	Items []Item
	N int
}

//init stack
func (stack *ItemStack) Init() *ItemStack {
	stack.Items = []Item{}

	return stack
}

//push stack Item
func (stack *ItemStack) Push(item Item) {
	if len(stack.Items) > stack.N {
		fmt.Println("栈已满")
		return
	}
	stack.Items = append(stack.Items, item)
}

//pop Item from stack
func (stack *ItemStack) Pop() Item {
	if len(stack.Items) == 0 {
		fmt.Println("栈已空")
		return nil
	}

	item := stack.Items[len(stack.Items) - 1]
	stack.Items = stack.Items[0:len(stack.Items) - 1]

	return item
}
