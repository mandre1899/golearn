package iteration

func Sum(nbrs [5]int) (sum int) {
	for _, nbr := range nbrs {
		sum += nbr
	}
	return
}
