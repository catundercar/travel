package sorts

func MergeSort(nums []int, n int) {
	merge_sort(nums, 0, n-1)
}

func merge_sort(nums []int, left int, right int) {
	// 递归终止条件
	if left >= right {
		return
	}
	mid := (left + right) / 2
	// 分治递归
	merge_sort(nums, left, mid)
	merge_sort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(arr []int, l int, mid int, r int) {

	// 因为需要直接修改 arr 数据，这里首先复制 [l,r] 的数据到新的数组中，用于赋值操作
	temp := make([]int, r-l+1)
	for i := l; i <= r; i++ {
		temp[i-l] = arr[i]
	}

	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		// left > mid: 左边数据处理完毕
		if left > mid {
			arr[i] = temp[right-l]
			right++
		} else if right > r {
			// 右边的数据处理完毕
			arr[i] = temp[left-l]
			left++
		} else if temp[left-l] <= temp[right-l] {
			// 左边比右边大
			arr[i] = temp[left-l]
			left++
		} else {
			arr[i] = temp[right-l]
			right++
		}
	}
	//if i == mid {
	//	for ; j < right; j++ {
	//		temp[k] = nums[j]
	//		k++
	//	}
	//} else {
	//	for ; i < mid; i++ {
	//		temp[k] = nums[i]
	//		k++
	//	}
	//}
	//for ; left < right; left++ {
	//	nums[left] = temp[left]
	//}
}
