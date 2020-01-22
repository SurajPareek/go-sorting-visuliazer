package sort

//SelectionSort - Exporting Function
func SelectionSort(limit int, arr []int) []int {
	for i := 0; i < limit-1; i++ {
		minIndex := i
		for j := i + 1; j < limit; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
	return arr
}
