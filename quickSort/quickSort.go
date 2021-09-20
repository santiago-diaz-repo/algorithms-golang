package quickSort

func QuickSort(arr *[]int,lowerBound int, upperBound int ) {
	if lowerBound < upperBound{
		location := partition(arr,lowerBound,upperBound)
		QuickSort(arr,lowerBound, location-1)
		QuickSort(arr,location+1,upperBound)
	}
}

func partition(arr *[]int, lowerBound int, upperBound int) int {
	pivot := (*arr)[lowerBound]
	start := lowerBound
	end := upperBound

	for start < end {
		for start < len(*arr) && (*arr)[start] <= pivot{
			start++
		}
		for end >= 0 && (*arr)[end] > pivot{
			end--
		}
		if start < end {
			swap(arr,start,end)
		}
	}
	swap(arr,lowerBound,end)
	return end
}

func swap(arr *[]int, pos1 int, pos2 int) {
	tmp := (*arr)[pos1]
	(*arr)[pos1] = (*arr)[pos2]
	(*arr)[pos2] = tmp
}