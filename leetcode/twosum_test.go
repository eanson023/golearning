package easy

import "testing"

func twoSum(nums []int, target int) []int {
	map1 := map[int]int{}
	var num int
	for idx := range nums {
		num = target - nums[idx]
		//如果说map1[num]存在 并且 v不等于idx
		if v, ok := map1[num]; ok && v != idx {
			return []int{v, idx}
		}
		map1[nums[idx]] = idx
	}
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestName(t *testing.T) {
	var head ListNode
	head.Val = 1
	t.Log(twoSum([]int{1, 2}, 3))
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{-1, nil}
	rec(l1, l2, &head, 0)
	return head.Next
}
func rec(l1 *ListNode, l2 *ListNode, tail *ListNode, tmp int) {
	if l1 == nil && l2 == nil && tmp == 0 {
		return
	}
	var num = tmp
	var ll1, ll2 *ListNode
	if l1 != nil {
		num += l1.Val
		ll1 = l1.Next
	}
	if l2 != nil {
		num += l2.Val
		ll2 = l2.Next
	}
	tmp = num / 10
	num = num % 10
	tail.Next = &ListNode{num, nil}
	tail = tail.Next
	rec(ll1, ll2, tail, tmp)
}
