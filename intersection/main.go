package main

import "fmt"

func main() {
	type ListNode struct {
		Data int
		Next *ListNode
	}
	// var com *ListNode
	com := &ListNode{
		Data: 7,
		Next: &ListNode{
			Data: 8,
			Next: nil,
		},
	}
	l1 := &ListNode{
		Data: 1,
		Next: &ListNode{
			Data: 2,
			Next: com,
		},
	}
	L2 := &ListNode{
		Data: 10,
		Next: &ListNode{
			Data: 9,
			Next: com,
		},
	}

	fc := func(l1, l2 *ListNode) *ListNode {
		p1, p2 := l1, l2

		p1cover := true
		p2cover := true

		for p1 != nil && p2 != nil {

			if p1 == p2 {
				return p1
			}
			if p1.Next == nil && p1cover {
				p1cover = false
				p1 = l2
			} else {
				p1cover = false
				p1 = p1.Next
			}
			if p2.Next == nil && p2cover {
				p2cover = false
				p2 = l1
			} else {
				p2 = p2.Next
			}
		}
		return nil
	}

	fmt.Println(fc(l1, L2))

	return
}

//A表：[1, 2, 3, 7, 8, 9]
//B表：[4, 5, 7, 8, 9]

//连接两个链表（表与表之间用 0 隔开）
//AB表：[1, 2, 3, 7, 8, 9, 0, 4, 5, 7, 8, 9, 0]
//BA表：[4, 5, 7, 8, 9, 0, 1, 2, 3, 7, 8, 9, 0]
