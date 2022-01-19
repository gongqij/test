package array

import "fmt"

type List struct {
	size uint64 // 节点数
	head *Node  // 头
	tail *Node  // 尾
}

type Node struct {
	Val  int
	Next *Node
}

// init
func (list *List) init(slice []int) {
	var temp *Node
	for i := 0; i < len(slice); i++ {
		newNode := &Node{
			Val:  slice[i],
			Next: nil,
		}
		if i == 0 {
			list.head = newNode
			temp = list.head
		} else {
			temp.Next = newNode
			temp = temp.Next
		}
	}
}

// print
func (list *List) print() {
	temp := list.head
	for temp != nil {
		fmt.Printf("Val: %d, current:%p, Next:%p\n", temp.Val, temp, temp.Next)
		temp = temp.Next
	}
}

//迭代
func reverseList(head *Node) *Node {
	var pre *Node
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

//递归
func reverseList2(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func mergeTwoLists(l1 *Node, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	for {
		if l1.Val < l2.Val {

		}
	}
}
