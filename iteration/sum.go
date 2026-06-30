package iteration

func Sum(nbrs []int) (sum int) {
	for _, nbr := range nbrs {
		sum += nbr
	}
	return
}

func SumAllTails(arr ...[]int) []int {
	var ret []int
	ret = make([]int, len(arr))
	for i, nbrs := range arr {
		if len(nbrs) == 0 {
			ret[i] = 0
		} else {
			ret[i] = Sum(nbrs[1:])
		}
	}
	return ret
}
