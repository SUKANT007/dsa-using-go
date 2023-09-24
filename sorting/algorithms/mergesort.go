package algorithms

import "sync"

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2
	var left, right []int
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		left = MergeSort(arr[:middle])
	}()

	go func() {
		defer wg.Done()
		right = MergeSort(arr[middle:])
	}()

	wg.Wait()

	return merge(left, right)
}

func merge(left, right []int) []int {
	i, j := 0, 0
	var result []int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// func MergeSort(l, r int, arr []int) {
// 	if l >= r {
// 		return
// 	}

// 	mid := l + (r-l)/2

// 	var wg sync.WaitGroup

// 	wg.Add(1)

// 	go func() {
// 		MergeSort(l, mid, arr)
// 		wg.Done()
// 	}()

// 	MergeSort(mid+1, r, arr)

// 	wg.Wait()

// 	merge(l, r, mid, arr)
// }

// func merge(l, r, mid int, arr []int) {
// 	var temp []int

// 	i, j := l, mid+1
// 	idx := 0

// 	for i <= mid && j <= r {
// 		if arr[i] <= arr[j] {
// 			temp = append(temp, arr[i])
// 			i++
// 		} else {
// 			temp = append(temp, arr[j])
// 			j++
// 		}
// 	}

// 	for ; i <= mid; i++ {
// 		temp = append(temp, arr[i])
// 	}

// 	for ; j <= r; j++ {
// 		temp = append(temp, arr[j])
// 	}

// 	for ; l <= r; l++ {
// 		arr[l] = temp[idx]
// 		idx++
// 	}
// }
