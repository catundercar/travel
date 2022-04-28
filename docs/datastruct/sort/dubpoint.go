package sort

func movezore(a []int) {
	i, j := 0, len(a)-1
	for i < j {
		if a[i] != 0 {
			i++
		} else if a[j] == 0 {
			j--
		} else {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
}
