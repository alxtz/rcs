package main

func QuickSelect_lomuto(input []int, kth int) int {
	// lomuto always uses the last val from input as pivot/partition val

	// when this func ends, will return len=1 list, containing the answer

	var recurFunc func(recurInput []int, recurKth int) []int

	recurFunc = func(recurInput []int, recurKth int) []int {
		// fmt.Println("received", recurInput, "recurKth", recurKth)
		pivotValue := recurInput[len(recurInput)-1]
		listToPartition := recurInput[:len(recurInput)-1]

		// start partitioning
		lowHighPartition := 0
		knownUnknownPartition := 0

		for knownUnknownPartition != len(listToPartition) {
			selectedUnknown := listToPartition[knownUnknownPartition]

			if selectedUnknown < pivotValue {
				listToPartition[lowHighPartition], listToPartition[knownUnknownPartition] = listToPartition[knownUnknownPartition], listToPartition[lowHighPartition]

				knownUnknownPartition += 1
				lowHighPartition += 1
				// fmt.Println("incre lh", lowHighPartition-1, "->", lowHighPartition)
			} else if pivotValue < selectedUnknown {
				knownUnknownPartition += 1
			} else if pivotValue == selectedUnknown {
				knownUnknownPartition += 1
			}

			// fmt.Println("listToP", listToPartition)
		}

		// fmt.Println(
		// 	"end",
		// 	"pivotVal",
		// 	pivotValue,
		// 	"listToPartition",
		// 	listToPartition,
		// 	"lowHighPartition",
		// 	lowHighPartition,
		// 	"knownUnknownPartition",
		// 	knownUnknownPartition,
		// )

		// fmt.Println("recurInput", recurInput, "recurKth", recurKth, "\n---")
		// fmt.Println("listToPartition", listToPartition, "recurKth", recurKth, "\n---")

		higherRankCount := len(listToPartition) - lowHighPartition // items count that had higher rank that pivot val
		// fmt.Println("higherRankCount", higherRankCount)

		if (higherRankCount + 1) == recurKth {
			return []int{pivotValue}
		}

		toPassLeft, toPassRight := recurInput[:lowHighPartition], recurInput[lowHighPartition:len(recurInput)-1]

		// fmt.Println("toPassLeft", toPassLeft)
		// fmt.Println("toPassRight", toPassRight)

		if recurKth < (higherRankCount + 1) {
			// fmt.Println("chose right", recurKth)
			return recurFunc(toPassRight, recurKth)
		} else if (higherRankCount + 1) < recurKth {
			// fmt.Println("chose left", recurKth-1-higherRankCount)
			return recurFunc(toPassLeft, recurKth-1-higherRankCount)
		}

		return []int{-10000}
	}

	ans := recurFunc(input, kth)

	return ans[0]
}
