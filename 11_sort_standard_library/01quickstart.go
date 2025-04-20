package main

import (
	"fmt"
	"sort"
)

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	// 返回false的时候调用swap()方法, 返回True不调用swap()方法
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	n := []uint{1, 3, 2}
	sort.Sort(NewInts(n))
	fmt.Println(n)
}
