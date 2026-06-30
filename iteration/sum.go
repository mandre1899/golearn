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
		ret[i] = Sum(nbrs[1:])
	}
	return ret
}
