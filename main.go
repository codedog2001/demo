package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 将链表每两个节点一组，然后组间反转
func reverseEveryTwoNodes(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var group []*ListNode
	curr := head
	for curr != nil {
		group = append(group, curr)
		next := curr.Next
		//next是当前组的第二个
		if next != nil {
			//curr移动到下一组的第一个
			curr = next.Next
			next.Next = nil //断开与后面的连接
		} else {
			curr = nil //没必要继续了，凑不出一组了
		}
	}
	var newHead *ListNode
	var lastNode *ListNode
	//组间逆序连接
	for i := len(group) - 1; i >= 0; i-- {
		if newHead == nil {
			//第一组的处理
			newHead = group[i]
			lastNode = group[i]
		} else {
			//把上一组尾巴连接到这一组
			lastNode.Next = group[i]
		}
		//走到该组的尾巴
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}
	}
	return newHead
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}
func ReverseEveryNNodes(head *ListNode, N int) *ListNode {
	if head == nil || head.Next == nil || N < 1 {
		return head
	}
	var groups []*ListNode
	curr := head
	for curr != nil {
		//加入组内
		groups = append(groups, curr)
		prev := curr
		//这里需要进行指针判空
		for i := 0; i < N-1 && prev.Next != nil; i++ {
			prev = prev.Next
		}
		curr = prev.Next
		prev.Next = nil
	}
	var newHead, lastNode *ListNode
	for i := len(groups) - 1; i >= 0; i-- {
		if newHead == nil {
			newHead = groups[i]
			lastNode = groups[i]
		} else {
			lastNode.Next = groups[i]
		}
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}
	}
	return newHead
}
func main() {
	// 测试用例1: 1->2->3->4->5
	head1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head1 = reverseEveryTwoNodes(head1)
	printList(head1)
	// 测试用例2: 1->2
	head2 := &ListNode{1, &ListNode{2, nil}}
	head2 = reverseEveryTwoNodes(head2)
	printList(head2)
	//测试反转N
	// 测试用例1: 1->2->3->4->5
	head3 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head3 = ReverseEveryNNodes(head3, 1)
	printList(head3)
	head4 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head4 = ReverseEveryNNodes(head4, 2)
	printList(head4)
	// 测试用例2: 1->2
	head5 := &ListNode{1, &ListNode{2, nil}}
	head5 = ReverseEveryNNodes(head5, 1)
	printList(head5)
}
