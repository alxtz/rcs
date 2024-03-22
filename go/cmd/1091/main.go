package main

import (
	"rcs/pkg/ds"
	"strconv"
)

type XY []int // x: horizontal, y: vertical

func shortestPathBinaryMatrix(grid [][]int) (shortestCount int) {
	if grid[0][0] == 1 {
		return -1
	}

	gridWidth := len(grid[0])
	gridHeight := len(grid)
	bfsMap := make(map[string]int)
	bfsMap["0.0"] = 0

	queue := ds.Queue[XY]{}
	queue.Enque([]int{0, 0})

	currDepth := 0

	maxSearch := 0

	for maxSearch < 2000000 {

		newLevelQueue := ds.Queue[XY]{}

		workQueueDedupMap := make(map[string]bool)

		for !queue.Empty() {
			maxSearch = maxSearch + 1

			traversed := queue.Deque()
			x, y := traversed[0], traversed[1]
			nextTraverse := [][]int{
				{x - 1, y - 1}, {x, y - 1}, {x + 1, y - 1},
				{x - 1, y} /* curr */, {x + 1, y},
				{x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1},
			}

			for _, pair := range nextTraverse {
				pairX, pairY := pair[0], pair[1]

				inRange := (0 <= pairX && pairX < gridWidth) && (0 <= pairY && pairY < gridHeight)
				_, alreadyFoundPath := bfsMap[strconv.Itoa(pairX)+"."+strconv.Itoa(pairY)]
				_, isDup := workQueueDedupMap[strconv.Itoa(pairX)+"."+strconv.Itoa(pairY)]

				if inRange && !alreadyFoundPath {
					notObstacle := grid[pairX][pairY] == 0
					if notObstacle && !isDup {
						newLevelQueue.Enque([]int{pairX, pairY})
						workQueueDedupMap[strconv.Itoa(pairX)+"."+strconv.Itoa(pairY)] = true
					}
				}
			}

			key := strconv.Itoa(x) + "." + strconv.Itoa(y)
			if currDistance, exists := bfsMap[key]; exists {
				if currDistance > currDepth {
					bfsMap[key] = currDepth
				}
			} else {
				bfsMap[key] = currDepth
			}
		}

		if newLevelQueue.Empty() {
			break
		}

		queue = newLevelQueue
		currDepth += 1

		// fmt.Println("queue", queue, bfsMap)

		// fmt.Println("search Depth", currDepth, "try expand searching", newLevelQueue.Len())
	}

	// fmt.Println("map", bfsMap, maxSearch)

	distance, distanceExists := bfsMap[strconv.Itoa(gridWidth-1)+"."+strconv.Itoa(gridHeight-1)]

	if distanceExists {
		return distance + 1
	}

	return -1
}

// 0, 1, 0, 1, 1, 0, 1, 0, 1, 0, 1,
// 1, 0, 0, 0, 1, 0, 0, 1, 0
// 0, 0, 1, 0, 1, 1, 1, 1, 1
// 0, 0, 1, 1, 0, 1, 1, 1, 1,
// 0, 0, 1, 1, 1, 0, 1, 1, 1, 1
// 1, 0, 1, 1, 1, 1, 0, 0, 0,
