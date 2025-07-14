package sorter

func PartitionSortAscending(list []int, low, high int) int {
	pivot := list[high]
	i := low - 1

	for j := low; j < high; j++ {
		if list[j] <= pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[high] = list[high], list[i+1]
	return i + 1
}

func PartitionSortDescending(list []int, low, high int) int {
	pivot := list[high]
	i := low - 1

	for j := low; j < high; j++ {
		if list[j] >= pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[high] = list[high], list[i+1]
	return i + 1
}

func QuickSort(list []int, low, high int, sorterFunction func([]int, int, int) int) {
	if low < high {
		pi := sorterFunction(list, low, high)

		QuickSort(list, low, pi-1, sorterFunction)
		QuickSort(list, pi+1, high, sorterFunction)
	}
}
