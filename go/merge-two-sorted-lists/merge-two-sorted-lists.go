package main

import(
	"fmt"
)

type ListNode struct{
	Val int
	Next *ListNode
} 

func mergeTwoLists(list1, list2 *ListNode) *ListNode{
	dummyNode := &ListNode{}
	tail := dummyNode

	for list1 != nil && list2 !=nil {
		if list1.Val <= list2.Val{
			tail.Next = list1
			list1 = list1.Next
		}else{
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil{
		tail.Next = list2
	}
	return dummyNode.Next
}

func main(){

	list1 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{}}}}

	list2 := ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{}}}}

	list1Pointer := &list1
	list2Pointer := &list2

	res := mergeTwoLists(list1Pointer, list2Pointer)


	fmt.Println("answer to the first example is:", res.Val, res.Next.Val, res.Next.Next.Val, res.Next.Next.Next.Val, res.Next.Next.Next.Next.Val, res.Next.Next.Next.Next.Next.Val, res.Next.Next.Next.Next.Next.Next.Val)
}