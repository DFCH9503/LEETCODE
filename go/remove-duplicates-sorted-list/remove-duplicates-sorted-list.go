package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil{
		if curr.Val == curr.Next.Val{
			curr.Next = curr.Next.Next
		}else{
			curr = curr.Next
		}
	}
	return head
}

func main() {
	res := deleteDuplicates(
		&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{}}}},
		
	)

	fmt.Println("answer to the first example is:", res.Val, res.Next.Val)
}