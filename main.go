package main

import (
	"flag"
	"fmt"
	"go-sorting-visuliazer/sort"
	"log"
)

type sorting func(int, []int) []int

var algoName string
var (
	sortingAlgoMatrix = map[string]sorting{
		"bubble":    sort.BubbleSort,
		"selection": sort.SelectionSort}
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

func printNumbers(limit int, arr []int) {
	for i := 0; i < limit; i++ {
		fmt.Printf("%d ", arr[i])
	}
}

func init() {
	const (
		algos = "selection, bubble, all"
	)
	flag.StringVar(&algoName, "sorting", "", algos)
}

func print(name string, limit int, arr []int) {
	fmt.Printf("Here is the list of sorted elements by %v \n", name)
	printNumbers(limit, arr)
}

func main() {
	flag.Parse()
	fmt.Println("Enter the size of an array")
	limit := getLimit()
	fmt.Printf("Enter the %d numbers of an element \n", limit)
	arr := getNumbers(limit)

	sortf, isexist := sortingAlgoMatrix[algoName]
	if isexist == false && algoName != "all" {
		log.Fatalln("sorting is not existed:", algoName)
	}

	if algoName != "all" {
		sortf(limit, arr)
		print(algoName, limit, arr)
	} else {
		// no matter about filename
		for name, sortf := range sortingAlgoMatrix {
			sortf(limit, arr)
			print(name, limit, arr)
		}
	}
	fmt.Println("\nThank You!")
}
