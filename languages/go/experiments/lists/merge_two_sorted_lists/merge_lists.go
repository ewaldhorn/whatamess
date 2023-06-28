package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(left *ListNode, right *ListNode) *ListNode {
	// first see if we actually have anything to do or not
	if left == nil && right == nil {
		return nil
	} else if left == nil {
		return right
	} else if right == nil {
		return left
	}

	if left.Val <= right.Val {
		left.Next = mergeTwoLists(left.Next, right)
		return left
	} else {
		right.Next = mergeTwoLists(left, right.Next)
		return right
	}
}

func (l *ListNode) addVal(val int) {
	newNode := &ListNode{Val: val}

	// first addition, just add and move along
	if l.Next == nil {
		l.Next = newNode
		return
	}

	// mmm, gotta walk till we find the nil
	// start at the front and continue until the NEXT node is nil
	currentPosition := l

	for currentPosition.Next != nil {
		currentPosition = currentPosition.Next
	}

	currentPosition.Next = newNode
}

func main() {
	// list1 := [1,2,4]
	// list2 := [1,3,4]
	// expectedOutput := [1,1,2,3,4,4]
	left := &ListNode{Val: 1}
	left.addVal(2)
	left.addVal(4)

	right := &ListNode{Val: 1}
	right.addVal(3)
	right.addVal(4)

	fmt.Println()
	fmt.Print("LEFT: ")
	printList(left)

	fmt.Print("RIGHT: ")
	printList(right)

	fmt.Print("MERGED: ")
	merged := mergeTwoLists(left, right)
	printList(merged)

	fmt.Println()
}

func printList(l *ListNode) {
	curr := l
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}
