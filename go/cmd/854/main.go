package main

import (
	"rcs/pkg/ds"
)

func main() {
	// fmt.Println("Hello, World")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	  {
			nodeId: [cost, nearestLeafId]
		}
*/
var costToNearestLeafMap = make(map[int]struct {
	Cost          int
	NearestLeafId int
})

var stack = ds.Stack[*TreeNode]{}

var targetNode *TreeNode
var targetNodeParents []*TreeNode

func FindClosestLeaf(root *TreeNode, k int) (ans int) {
	costToNearestLeafMap = make(map[int]struct {
		Cost          int
		NearestLeafId int
	})
	stack = ds.Stack[*TreeNode]{}
	targetNode = &TreeNode{}
	targetNodeParents = []*TreeNode{}

	dfs(root, k)

	// fmt.Println("targetNode", targetNode)
	// fmt.Println("targetNodeParents", targetNodeParents)
	// fmt.Println("map", costToNearestLeafMap)

	ansCost := costToNearestLeafMap[targetNode.Val].Cost
	ans = costToNearestLeafMap[targetNode.Val].NearestLeafId

	// fmt.Println("ans", ans)

	for i := range targetNodeParents {
		// fmt.Println(targetNodeParents[len(targetNodeParents)-1-i])
		parentId := targetNodeParents[len(targetNodeParents)-1-i].Val
		parentCostObj := costToNearestLeafMap[parentId]
		accuCost := i + 1 + parentCostObj.Cost
		// fmt.Println(parentId, parentCostObj, accuCost)
		if accuCost < ansCost {
			ans = parentCostObj.NearestLeafId
			ansCost = parentCostObj.Cost
			// fmt.Println("new ans", ans)
		}
	}

	return ans
}

func dfs(node *TreeNode, given int) {
	if node.Val == given {
		targetNode = node
		targetNodeParents = append(targetNodeParents, stack.Slice...)
	}

	stack.Push(node)

	// fmt.Println("stack", stack.Slice)

	if node.Left != nil {
		dfs(node.Left, given)
	}

	if node.Right != nil {
		dfs(node.Right, given)
	}

	if node.Left == nil && node.Right == nil {
		// is a confirmed leaf, updating the upstream
		for i := stack.Len() - 1; i >= 0; i-- {
			cost := stack.Len() - i - 1
			upstreamItemId := stack.Slice[i].Val
			if existingCostObj, ok := costToNearestLeafMap[upstreamItemId]; !ok {
				costToNearestLeafMap[upstreamItemId] = struct {
					Cost          int
					NearestLeafId int
				}{
					cost, node.Val,
				}
			} else if existingCostObj.Cost > cost {
				costToNearestLeafMap[upstreamItemId] = struct {
					Cost          int
					NearestLeafId int
				}{
					cost, node.Val,
				}
			}
		}

	}

	stack.Pop()
}
