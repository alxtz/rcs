package main

import (
	"fmt"
	"rcs/pkg/ds"
)

func main() {
	fmt.Println("Hello, World")
}

type YXV = []int // x pos, y pos, v cost

func GetBubbleTea(grid [][]byte) (ans int) {
	ans = -1
	gridWidth := len(grid[0])
	gridHeight := len(grid)

	var mem [][]int // 0 = not visited, 1 = visited

	var personYXV YXV

	for i := range grid {
		mem = append(mem, []int{})
		for j := range grid[i] {
			(mem[i]) = append((mem[i]), 0)

			if grid[i][j] == '*' {
				personYXV = YXV{i, j, 0}
			}
		}
	}

	queue := ds.Queue[YXV]{}
	queue.Enque(personYXV)

	for !queue.Empty() {
		removed := queue.Deque()

		removedY, removedX := removed[0], removed[1]

		if grid[removedY][removedX] == '#' {
			ans = removed[2]
			break
		}

		toPush := []YXV{
			{removedY, removedX - 1, removed[2] + 1},
			{removedY, removedX + 1, removed[2] + 1},
			{removedY + 1, removedX, removed[2] + 1},
			{removedY - 1, removedX, removed[2] + 1},
		}

		for _, toPushYx := range toPush {
			toPushY, toPushX := toPushYx[0], toPushYx[1]
			if (0 <= toPushY && toPushY < gridHeight) && (0 <= toPushX && toPushX < gridWidth) {
				if grid[toPushY][toPushX] != 'X' && mem[toPushY][toPushX] == 0 {
					queue.Enque(toPushYx)
					mem[toPushY][toPushX] = 1
				}
			}
		}
	}

	return ans
}
