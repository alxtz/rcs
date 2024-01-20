package main

type StackNode struct {
	Node  *ExpressionTreeNode
	IsOpr bool
}

type Stack struct {
	Slice []StackNode
}

func (s *Stack) Pop() StackNode {
	var itemToPop = s.Slice[len(s.Slice)-1]
	s.Slice = s.Slice[:len(s.Slice)-1]
	return itemToPop
}

func (s *Stack) Push(node *ExpressionTreeNode, isOpr bool) {
	s.Slice = append(s.Slice, StackNode{Node: node, IsOpr: isOpr})
}

func ConvertToRPN_2(infixExpr []string) []string {
	var output []string
	var stack Stack

	for _, ele := range infixExpr {
		if _, isOperator := precMap[ele]; isOperator {
			for len(stack.Slice) > 0 && stack.Slice[len(stack.Slice)-1].IsOpr &&
				precMap[stack.Slice[len(stack.Slice)-1].Node.Symbol] >= precMap[ele] {
				output = append(output, stack.Pop().Node.Symbol)
			}
			stack.Push(&ExpressionTreeNode{Symbol: ele}, true)
		} else if ele == "(" {
			stack.Push(&ExpressionTreeNode{Symbol: ele}, true)
		} else if ele == ")" {
			for len(stack.Slice) > 0 && stack.Slice[len(stack.Slice)-1].Node.Symbol != "(" {
				output = append(output, stack.Pop().Node.Symbol)
			}
			stack.Pop() // Pop the "("
		} else {
			output = append(output, ele)
		}
	}

	for len(stack.Slice) > 0 {
		output = append(output, stack.Pop().Node.Symbol)
	}

	return output
}

func PostfixToBet_2(postfixExpr []string) *ExpressionTreeNode {
	if len(postfixExpr) == 0 {
		return nil
	}

	var stack Stack

	for _, ele := range postfixExpr {
		if _, isOperator := precMap[ele]; isOperator {
			right := stack.Pop().Node
			left := stack.Pop().Node
			stack.Push(&ExpressionTreeNode{Symbol: ele, Left: left, Right: right}, false)
		} else {
			// val, _ := strconv.Atoi(ele)
			stack.Push(&ExpressionTreeNode{Symbol: ele}, false)
		}
	}

	return stack.Pop().Node
}

func Build_2(infixExpr []string) *ExpressionTreeNode {
	rpn := ConvertToRPN_2(infixExpr)
	return PostfixToBet_2(rpn)
}
