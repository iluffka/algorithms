package binary

/*func main()  {
	arr := []int{
		0,2,333,444,567,1234,
	}
	fmt.Println(BinarySearch(arr,567))
	fmt.Println(BinarySearchRec(arr,567))
}*/

func BinarySearch(arr []int, element int) int {
	first := 0
	last := len(arr)
	if last == 0 || element > arr[last - 1] {
		return -1
	}
	for first < last {
		mid := (first + last)/2
		if element <= arr[mid] {
			last = mid
		} else {
			first = mid +1
		}
	}

	if arr[first] != element {
		return -1
	}

	return first
}

func BinarySearchRec(data []int, needle int) int {
	left := 0
	right := len(data)
	if right == 0 || needle < data[left] || needle > data[right - 1] {
		return -1
	}

	return intsRecursion(data, needle, left, right - 1)
}

func intsRecursion(data []int, needle, left, right int) int {
	if left >= right{
		if data[left] == needle {
			return left
		}
		return -1
	}

	mid := left + (right - left)/2
	if v := data[mid]; v > needle {
		return intsRecursion(data, needle, left, mid)
	} else if v < needle {
		return intsRecursion(data, needle, mid + 1, right)
	}

	return mid
}