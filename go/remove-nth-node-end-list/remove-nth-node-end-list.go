package main

import(
	"fmt"
) 

type ListNode struct {
    Val int
    Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    slow, fast := head, head
    for i := 0; i < n; i++ {
        fast = fast.Next
    }
    if fast == nil {
        return head.Next
    }
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return head
}

func main(){
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

    res := removeNthFromEnd(node1, 2)


	fmt.Printf("answer to the first example is: [%v, %v, %v, %v]", res.Val, res.Next.Val, res.Next.Next.Val, res.Next.Next.Next.Val)
}