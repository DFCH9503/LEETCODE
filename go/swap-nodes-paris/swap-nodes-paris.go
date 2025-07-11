package main

import(
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode{
	if head == nil || head.Next == nil {
		return head
	}

	head, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head

	return head
}

func main(){
	head := &ListNode{Val: 1}
	nodeA := &ListNode{Val: 2}
	nodeB := &ListNode{Val: 3}
	nodeC := &ListNode{Val: 4}

	head.Next = nodeA
	nodeA.Next = nodeB
	nodeB.Next = nodeC

	res := swapPairs(head)

	fmt.Printf("Swaping the nodes in %v the result is %v", head, res)

}