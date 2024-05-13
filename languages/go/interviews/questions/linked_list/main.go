package main

type ListNode struct {
	Data interface{}
	Next *ListNode
}

// calculate the length of the list by walking the list
func (list ListNode) Length() (length int) {
	length = 0

	for pos := &list; pos != nil; pos = pos.Next {
		length++
	}

	return // named return example
}

// find the last element in the list, it does not have a Next node
func (list ListNode) Tail() *ListNode {
	var element *ListNode

	for pos := &list; pos != nil; pos = pos.Next {
		element = pos
	}

	return element
}
