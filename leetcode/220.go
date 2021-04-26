package main

import "fmt"

type BSTNode struct {
	val   int
	left  *BSTNode
	right *BSTNode
}

//     root
//  left   right
// left < root < right

func search(v int, head *BSTNode) *BSTNode {
	t := head
	for t != nil {
		if t.val == v {
			return t
		} else if t.val > v {
			t = t.left
		} else {
			t = t.right
		}
	}

	return nil
}

func insert(v int, head *BSTNode) *BSTNode {
	if head == nil {
		head = &BSTNode{val: v}
		return head
	}
	for t := head; t != nil; {
		if t.val < v {
			if t.right != nil {
				t = t.right
			} else {
				t.right = &BSTNode{
					val: v,
				}
				return t.right
			}
		} else {
			// t.val >= v
			if t.left != nil {
				t = t.left
			} else {
				t.left = &BSTNode{
					val: v,
				}
				return t.left
			}
		}
	}

	return nil
}

func successor(head *BSTNode) int {
	fmt.Println(head.val)
	t := head.right
	if t == nil {
		return 0x3f3f3f
	}
	for t.left != nil {
		t = t.left
	}
	return t.val
}

func predecessor(head *BSTNode) int {
	t := head.left
	if t == nil {
		return -0x3f3f3f
	}
	for t.right != nil {
		t = t.right
	}
	return t.val
}

func del(v int, head *BSTNode) *BSTNode {
	if head == nil {
		return nil
	}

	if head.val < v {
		head.right = del(v, head.right)
	} else if head.val > v {
		head.left = del(v, head.left)
	} else {
		if head.left == nil && head.right == nil {
			return nil
		}
		if head.right != nil {
			head.val = successor(head)
			head.right = del(head.val, head.right)
		} else {
			head.val = predecessor(head)
			head.left = del(head.val, head.left)
		}
	}

	return head
}

var result []int

func dfs(root *BSTNode) {
	if root == nil {
		return
	}
	result = append(result, root.val)
	if root.left != nil {
		dfs(root.left)
	}
	if root.right != nil {
		dfs(root.right)
	}
	//result = append(result, root.Val)
	return
}

// 寻找 大于等于 v 的最小值
func searchLower(v int, head *BSTNode) *BSTNode {
	var res *BSTNode
	t := head
	fmt.Println(v)
	for t != nil {
		if t.val == v {
			return t
		} else if t.val > v {
			res = t
			t = t.left
		} else {
			t = t.right
		}
	}

	return res
}

func preorderTraversal(root *BSTNode) []int {
	result = []int{}
	if root == nil {
		return result
	}
	dfs(root)
	return result
}

// 滑动窗口
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k == 0 {
		return false
	}
	var root *BSTNode
	lenn := len(nums)
	for i := 0; i < lenn; i ++ {
		n := searchLower(nums[i] - t, root)
		if n != nil {
			// fmt.Println(preorderTraversal(root), n.val, i, nums[i])
			if n.val <= nums[i] + t {
				return true
			}
		}
		now := insert(nums[i], root)
		if root == nil {
			root = now
		}
		// fmt.Println(preorderTraversal(root))
		if i >= k {
			del(nums[i-k], root)
		}
		fmt.Println(preorderTraversal(root))
	}
	// fmt.Println(preorderTraversal(root))
	return false
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1, 2}, 0, 1))
}
