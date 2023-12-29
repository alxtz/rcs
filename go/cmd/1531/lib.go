package main

type Pair struct {
	Char  string
	Count int
}

type ValuePair struct {
	Char     string `json:"char"`
	Count    int    `json:"count"`
	AvgValue int    `json:"avg_value"`
}

func toValueSlice(pList []Pair, key string) []ValuePair {
	var result []ValuePair

	for _, newestVal := range pList {
		// fmt.Println("debug-a")

		if len(result) == 0 {
			// fmt.Println("debug-a.init")
			var charToUse = newestVal.Char
			if charToUse != key {
				charToUse = "*"
			}
			result = append(result, ValuePair{
				Char:     charToUse,
				Count:    newestVal.Count,
				AvgValue: -1,
			})
			continue
		}

		if newestVal.Char != key {
			// fmt.Println("debug-b", result[len(result)-1])
			if lastItemOfResult := result[len(result)-1]; lastItemOfResult.Char != key {
				// fmt.Println("debug-b.add")
				result[len(result)-1].Count += newestVal.Count
			} else {
				result = append(result, ValuePair{
					Char:     "*",
					Count:    newestVal.Count,
					AvgValue: -1,
				})
			}
		} else {
			// 預期一定是 * -> a
			result = append(result, ValuePair{
				Char:     newestVal.Char,
				Count:    newestVal.Count,
				AvgValue: -1,
			})
		}
	}

	return result
}
