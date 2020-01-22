package main

import (
	"fmt"
	"go-sorting-visuliazer/sort"
)

func getSortingAlgoName() string {
	var name string
	fmt.Scanf("%v", &name)
	return name
}

func getNumbers(n int) []int {
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	return arr
}

func getLimit() int {
	var limit int
	fmt.Scanf("%d", &limit)
	return limit
}

func main() {
	fmt.Println("Enter the size of an array")
	limit := getLimit()
	fmt.Printf("Enter the %d numbers of an element \n", limit)
	arr := getNumbers(limit)
	sort.BubbleSort(limit, arr)
}
