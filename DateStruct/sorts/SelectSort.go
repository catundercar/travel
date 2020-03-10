package sorts_v1

func SelectSort(n []int) {
	for i := 0; i < len(n); i++ {
		j := len(n) - 1
		min := n[j]
		x := j
		for ; j >= i; j-- {
			if n[j] < min {
				min = n[j]
				x = j
			}
		}
		temp := n[i]
		n[i] = n[x]
		n[x] = temp
	}
}
