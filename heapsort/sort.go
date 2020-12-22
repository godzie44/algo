package heapsort

func buildMaxHeap(arr []int, heapifyFn func(arr []int, ind int)) {
	for i := len(arr) / 2; i >= 0; i-- {
		heapifyFn(arr, i)
	}
}

func MaxHeapifyRecursive(arr []int, ind int) {
	left := ind*2 + 1
	right := ind*2 + 2

	largest := ind
	if left <= len(arr)-1 && arr[left] > arr[ind] {
		largest = left
	}
	if right <= len(arr)-1 && arr[right] > arr[largest] {
		largest = right
	}

	if largest != ind {
		arr[ind], arr[largest] = arr[largest], arr[ind]
		MaxHeapifyRecursive(arr, largest)
	}
}

func MaxHeapifyLinear(arr []int, ind int) {
	for {
		left := ind*2 + 1
		right := ind*2 + 2

		largest := ind
		if left <= len(arr)-1 && arr[left] > arr[ind] {
			largest = left
		}
		if right <= len(arr)-1 && arr[right] > arr[largest] {
			largest = right
		}

		if largest == ind {
			return
		}

		arr[ind], arr[largest] = arr[largest], arr[ind]
		ind = largest
	}
}

func Sort(arr []int, heapifyFn func(arr []int, ind int)) {
	buildMaxHeap(arr, heapifyFn)

	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapifyFn(arr[:i], 0)
	}
}
