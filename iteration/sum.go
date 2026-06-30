package iteration

func Sum(nbrs []int) (sum int) {
	for _, nbr := range nbrs {
		sum += nbr
	}
	return
}
