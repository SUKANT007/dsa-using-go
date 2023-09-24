package main

import (
	"dsa-using-go/sorting/algorithms"
	"fmt"
)

func main() {
	fmt.Printf("Enter size of your array: ")
	var n int
	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	algorithms.BubbleSort(arr)

	// sortedArr := algorithms.MergeSort(arr)

	fmt.Println("Your sorted array is: ", arr)
}
