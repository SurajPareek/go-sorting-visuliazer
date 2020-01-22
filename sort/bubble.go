package sort

// BubbleSort - Exported Function
func BubbleSort(limit int, arr []int) []int {
	for i := 0; i < limit-1; i++ {
		for j := limit - 1; j > i; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
	return arr
}
