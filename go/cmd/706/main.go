package main

import "fmt"

func main() {
	myHM := Constructor()

	myHM.Put(1, 1)
	myHM.Put(2, 2)
	fmt.Println("myHM.Get(1)", myHM.Get(1) == 1)
	fmt.Println("myHM.Get(3)", myHM.Get(3) == -1)
	myHM.Put(2, 1)
	myHM.Get(2)
	fmt.Println("myHM.Get(2)", myHM.Get(2) == 1)
	myHM.Remove(2)
	myHM.Get(2)
	fmt.Println("myHM.Get(2)", myHM.Get(2) == -1)
}

type MyHashMap struct {
	buckets    []*MyListNode
	totalCount int
}

type MyListNode struct {
	Key  int
	Val  int
	Next *MyListNode
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets: []*MyListNode{
			{-1, -1, nil}, {-1, -1, nil},
			{-1, -1, nil}, {-1, -1, nil},
			{-1, -1, nil}, {-1, -1, nil},
			{-1, -1, nil}, {-1, -1, nil},
			{-1, -1, nil}, {-1, -1, nil},
		},
		totalCount: 0,
	}
}

func (myHashMap *MyHashMap) Put(key int, value int) {
	bucketCount := len(myHashMap.buckets)

	remainder := key % bucketCount

	listCorrespondsToBucket := myHashMap.buckets[remainder]

	listToTraverse := listCorrespondsToBucket

	for {
		if listToTraverse.Key != key {
			if listToTraverse.Next == nil {
				listToTraverse.Next = &MyListNode{Key: key, Val: value, Next: nil}
				break
			} else {
				listToTraverse = listToTraverse.Next
			}
		} else {
			listToTraverse.Val = value
			break
		}
	}

}

func (myHashMap *MyHashMap) Get(key int) int {
	bucketCount := len(myHashMap.buckets)
	remainder := key % bucketCount

	listCorrespondsToBucket := myHashMap.buckets[remainder]

	listToTraverse := listCorrespondsToBucket

	for {
		if listToTraverse.Key != key {
			if listToTraverse.Next == nil {
				return -1
			} else {
				listToTraverse = listToTraverse.Next
			}
		} else {
			return listToTraverse.Val
		}

	}
}

func (myHashMap *MyHashMap) Remove(key int) {
	bucketCount := len(myHashMap.buckets)
	remainder := key % bucketCount

	listCorrespondsToBucket := myHashMap.buckets[remainder]

	parent := listCorrespondsToBucket

	for {
		if parent.Next == nil {
			break
		} else {
			candidate := parent.Next
			if candidate.Key == key {
				parent.Next = candidate.Next
			} else {
				parent = candidate
			}
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
