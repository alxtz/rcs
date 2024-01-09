package main

func main() {
}

var presdMap = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

func ConvertToRPN(expression []string) []string {
	var oprStack = Stack{}
	var numQueue = Queue{}

	var input = expression

	for {
		if len(input) == 0 {
			if len(oprStack.Slice) != 0 {
				var toPop = oprStack.Pop()
				numQueue.Enque(toPop)
				continue
			} else {
				break
			}
		}

		var head = input[0]

		if head == "+" || head == "-" || head == "*" || head == "/" {
			var oprToPush = head

			if len(oprStack.Slice) > 0 {
				var toPushPresd = presdMap[head]

				for {
					if len(oprStack.Slice) == 0 {
						break
					}

					var topEle = oprStack.Slice[len(oprStack.Slice)-1]
					if topEle == "(" {
						break
					}

					var topElePresd = presdMap[topEle]

					if toPushPresd <= topElePresd {
						numQueue.Enque(oprStack.Pop())
						continue
					} else {
						break
					}
				}
			}

			oprStack.Push(oprToPush)
		} else if head == "(" {
			oprStack.Push(head)
		} else if head == ")" {
			for {
				var topEle = oprStack.Slice[len(oprStack.Slice)-1]
				if topEle == "(" {
					oprStack.Pop()
					break
				} else {
					numQueue.Enque(oprStack.Pop())
				}
			}
		} else {
			numQueue.Enque(head)
		}

		input = input[1:]
	}

	if numQueue.Slice == nil {
		return []string{}
	}

	return numQueue.Slice
}

type Stack struct {
	Slice []string
}

func (s *Stack) Pop() string {
	var itemToPop = s.Slice[len(s.Slice)-1]
	s.Slice = s.Slice[:len(s.Slice)-1]

	return itemToPop
}

func (s *Stack) Push(val string) {
	s.Slice = append(s.Slice, val)
}

type Queue struct {
	Slice []string
}

func (q *Queue) Deque() string {
	var itemToPop = q.Slice[0]
	q.Slice = q.Slice[1:len(q.Slice)]

	return itemToPop
}

func (q *Queue) Enque(val string) {
	q.Slice = append(q.Slice, val)
}
