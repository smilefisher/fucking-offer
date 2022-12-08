package main

import "fmt"

type LinkedNode struct {
	Data int
	Next *LinkedNode
}

func main() {
	head :=  &LinkedNode{}
	*head = LinkedNode{
		Data: 1,
		Next:  &LinkedNode{
			Next:  &LinkedNode{
				Next: head,
				Data: 3,
			},
			Data: 2,
		},
	}
	fmt.Println(circle(head))
}


func circle(head *LinkedNode) bool {
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false

}