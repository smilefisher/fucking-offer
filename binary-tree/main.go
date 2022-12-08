package main

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func main() {

	var node = &TreeNode{
		data: 1,
		left: &TreeNode{
			data: 2,
			left: &TreeNode{
				data: 4,
			},
			right: &TreeNode{
				data: 5,
				left: &TreeNode{
					data: 7,
				},
				right: &TreeNode{
					data: 8,
				},
			},
		},
		right: &TreeNode{
			data: 3,
			right: &TreeNode{
				data: 6,
			},
		},
	}

	// recursive(node)
	levelTraverse(node)
}

//recursive 递归 ，前序，中序，后续
func recursive(node *TreeNode) {
	if node == nil {
		return
	}
	// fmt.Println(node.data) //前序 根节点在前遍历
	recursive(node.left)
	// fmt.Println(node.data) //中序 根节点在遍历
	recursive(node.right)
	fmt.Println(node.data) //后续遍历
}

//水平遍历 用到堆
func levelTraverse(node *TreeNode) {
	var queen = make(chan *TreeNode, 10)
	queen <- node
	for len(queen) > 0 {
		v := <-queen
		fmt.Println(v.data)
		if v.left != nil {
			queen <- v.left
		}
		if v.right != nil {
			queen <- v.right
		}
	}
}