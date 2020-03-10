package sorts_v1

func Bubblesort(nums []int) {
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j+1] < nums[j] {
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
