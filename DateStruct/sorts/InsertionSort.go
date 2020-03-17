package sorts

func InsertionSort(n []int) {
	for i := 1; i < len(n); i++ {
		value := n[i]
		j := i - 1
		for ; j >= 0; j-- {
			if n[j] > value {
				n[j+1] = n[j]
			} else {
				break
			}
		}
		n[j+1] = value
	}
}
