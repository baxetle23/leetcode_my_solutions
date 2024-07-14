package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeNode(array []*int) *TreeNode {
	if len(array) == 0 || array[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *array[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(array) {
		current := queue[0]
		queue = queue[1:]

		if i < len(array) && array[i] != nil {
			current.Left = &TreeNode{Val: *array[i]}
			queue = append(queue, current.Left)
		}
		i++

		if i < len(array) && array[i] != nil {
			current.Right = &TreeNode{Val: *array[i]}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}

func IntPtr(i int) *int {
	return &i
}

func helperInorderTraversalReq(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	helperInorderTraversalReq(root.Left, nums)
	*nums = append(*nums, root.Val)
	helperInorderTraversalReq(root.Right, nums)
}

func InorderTraversalReq(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	nums := make([]int, 0)
	helperInorderTraversalReq(root, &nums)
	return nums
}

// TODO: реализовать прямоц обход дерева с помощью итерации
func InorderTraversalIter(root *TreeNode) []int {
	return []int{}
}

func IsSameTreeHelper(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return IsSameTreeHelper(p.Right, q.Right) && IsSameTreeHelper(p.Left, q.Left)
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	return IsSameTreeHelper(p, q)
}

func IsSymmetricHelper(right *TreeNode, left *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if right == nil || left == nil {
		return false
	}

	if right.Val != left.Val {
		return false
	}

	return IsSymmetricHelper(right.Right, left.Left) && IsSymmetricHelper(right.Left, left.Right)
}

func isSymmetric(root *TreeNode) bool {
	return IsSymmetricHelper(root.Right, root.Left)
}

func maxDepthHelper(root *TreeNode, step int) int {
	if root == nil {
		return step
	}

	return int(math.Max(float64(maxDepthHelper(root.Right, step+1)), float64(maxDepthHelper(root.Left, step+1))))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftStep := maxDepthHelper(root.Left, 1)
	rightStep := maxDepthHelper(root.Right, 1)
	return int(math.Max(float64(leftStep), float64(rightStep)))
}

func SortedArrayToBst(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	root := &TreeNode{Val: nums[middle]}
	root.Left = SortedArrayToBst(nums[:middle])
	root.Right = SortedArrayToBst(nums[middle+1:])
	return root
}
