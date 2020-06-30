package sort

//func MergeSort(nums []int) {
//	merge_sort(nums, 0, len(nums)-1)
//}
//
//func merge_sort(nums []int, left int, right int) {
//	// 递归终止条件
//	if left >= right {
//		return
//	}
//	mid := (left + right) / 2
//	// 分治递归
//	merge_sort(nums, left, mid)
//	merge_sort(nums, mid+1, right)
//	merge(nums, left, mid, right)
//}
//
//func merge(arr []int, l int, mid int, r int) {
//
//	// 因为需要直接修改 arr 数据，这里首先复制 [l,r] 的数据到新的数组中，用于赋值操作
//	temp := make([]int, r-l+1)
//	for i := l; i <= r; i++ {
//		temp[i-l] = arr[i]
//	}
//
//	left := l
//	right := mid + 1
//
//	for i := l; i <= r; i++ {
//		// left > mid: 左边数据处理完毕
//		if left > mid {
//			arr[i] = temp[right-l]
//			right++
//		} else if right > r {
//			// 右边的数据处理完毕
//			arr[i] = temp[left-l]
//			left++
//		} else if temp[left-l] <= temp[right-l] {
//			// 左边比右边大
//			arr[i] = temp[left-l]
//			left++
//		} else {
//			arr[i] = temp[right-l]
//			right++
//		}
//	}
//}

func MergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}
