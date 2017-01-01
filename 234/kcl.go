package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
    Val int
    Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil{
		return true
	}
	// find middle position
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}

	sec := slow.Next
	slow.Next = nil
	fmt.Println(sec)

	// reverse 
	sec = reverseList(sec)
	for head != nil && sec != nil{
		if head.Val != sec.Val{
			return false
		}
		head = head.Next
		sec = sec.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode{
	if head == nil{
		return head
	}
	var left *ListNode
	for head != nil{
		tmp := head.Next
		head.Next = left
		left = head
		head = tmp
	}
	return left
}

func main() {
	/*
	n5 := &ListNode{5, nil}
	n4 := &ListNode{4, n5}
	n3 := &ListNode{3, n4}
*/
	n2 := &ListNode{0, nil}
	n1 := &ListNode{0, n2}
	fmt.Println(isPalindrome(n1))
}
