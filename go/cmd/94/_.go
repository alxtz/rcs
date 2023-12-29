// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Hello, World 0")
// }

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// // func inorderTraversal(root *TreeNode) []int {
// // 	var traversalResults = []int{}
// // 	var currNode *TreeNode = root
// // 	var currNodeParent *TreeNode = nil

// // 	for {
// // 		if currNode.Left != nil {
// // 			currNodeParent = currNode
// // 			currNode = currNode.Left
// // 			continue
// // 		}

// // 		if (currNode.Left == nil) {
// // 			traversalResults = append(traversalResults, currNode.Val)

// // 			currNodeParent.Left = nil
// // 			currNode = currNodeParent
// // 		}
// // 	}

// // 	return []int{}
// // }
// // 感覺需要一個歷史的 variable
// // 歷史的 variable, 又是 last in first out
// // 那就是一個 stack 啊！

// type TreeNodeWithUsed struct {
// 	Node *TreeNode
// 	Used bool
// }

// func inorderTraversal(root *TreeNode) []int {
// 	var traversalResults = []int{}

// 	var stack []TreeNodeWithUsed = []TreeNodeWithUsed{
// 		TreeNodeWithUsed{
// 			Node: root,
// 			Used: false,
// 		},
// 	}

// 	for {
// 		if len(stack) == 0 {
// 			break
// 		}

// 		var lastOfStack = stack[len(stack)-1]

// 		fmt.Println("debug", lastOfStack.Node.Val, stack, lastOfStack.Node)

// 		// 左下探訪
// 		if lastOfStack.Node.Left != nil {
// 			if !lastOfStack.Used {
// 				fmt.Println("cond-a")
// 				stack = append(stack, TreeNodeWithUsed{Node: lastOfStack.Node.Left, Used: false})
// 				fmt.Println("---end---")
// 				continue
// 			} else {
// 				fmt.Println("cond-a.1", lastOfStack.Node)
// 			}
// 		}

// 		if lastOfStack.Node.Left == nil {
// 			if lastOfStack.Node.Right != nil {
// 				fmt.Println("cond-b")
// 				stack = append(stack, TreeNodeWithUsed{Node: lastOfStack.Node.Right, Used: false})
// 			} else {
// 				fmt.Println("cond-c")
// 				// 到絕對底了，寫入並回溯
// 				traversalResults = append(traversalResults, lastOfStack.Node.Val)
// 				if len(stack) > 1 {
// 					stack[len(stack)-2].Used = true
// 					// stack[len(stack)-2].Node.Left = nil
// 					// stack[len(stack)-2].Node.Right = nil
// 				}
// 				stack = stack[:len(stack)-1] // pop()
// 			}

// 			fmt.Println("---end---")
// 		}

// 		// if lastOfStack.Node.Left == nil && lastOfStack.Node.Right == nil {
// 		// 	fmt.Println("cond-b")
// 		// 	traversalResults = append(traversalResults, lastOfStack.Node.Val)
// 		// 	stack = stack[:len(stack)-1] // pop()

// 		// 	var latestLastOfStack = stack[len(stack)-1]
// 		// 	latestLastOfStack.Node.Left = nil
// 		// }

// 		// if lastOfStack.Node.Left == nil && lastOfStack.Node.Right != nil {
// 		// 	if !lastOfStack.Used {
// 		// 		fmt.Println("cond-c")
// 		// 		traversalResults = append(traversalResults, lastOfStack.Node.Val)
// 		// 		// lastOfStack.Used = true
// 		// 		stack = append(stack, TreeNodeWithUsed{
// 		// 			Node: lastOfStack.Node.Right,
// 		// 			Used: false,
// 		// 		})
// 		// 	} else {
// 		// 		fmt.Println("cond-d")
// 		// 		stack = stack[:len(stack)-1] // pop()
// 		// 	}
// 		// }
// 	}

// 	return traversalResults
// }
