package slicessa

func sliceLen() (int, int) {
	s1 := []int{1, 2, 3}

	l1 := len(s1)
	c1 := cap(s1)
	return l1, c1
}
