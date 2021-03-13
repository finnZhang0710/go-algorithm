package sort

/**
选择排序，每次选择最大或者最小的，交换即可
*/

func selectSort(arr []int) []int {
	n := len(arr)

	for j := n - 1; j > 0; j-- {
		maxIndex := j
		for i := 0; i < n-1-j; i++ {
			if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}

		arr[j], arr[maxIndex] = arr[maxIndex], arr[j]

	}

	return arr
}

func selectSort2(arr []int) []int {

	n := len(arr)

	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[minIndex] < arr[j] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}
