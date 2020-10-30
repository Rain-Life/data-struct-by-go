package main

import (
	"learnGo/linkedList/doublyLoopLinkedList"
	"testing"
)

func TestAddFromHead(t *testing.T)  {
	tests := []struct {insertValue,rightValue interface{}} {
		{1, 1},
		//{2, 2},
		//{3, 3},
	}

	list := doublyLoopLinkedList.List{}
	for _, testData := range tests {
		list.AddFromHead(testData.insertValue)
		if actual := 1; actual != testData.rightValue {
			t.Errorf("AddFromHead(%d); " + "got %d; excepted %d", testData.insertValue, actual, testData.rightValue)
		}
	}
}
