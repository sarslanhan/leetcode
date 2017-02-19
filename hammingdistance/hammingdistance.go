package hammingdistance

func hammingDistance(x int, y int) int {
	distance := 0

	xor := x ^ y

	for xor != 0 {
		if 1&xor == 1 {
			distance++
		}

		xor = xor >> 1
	}

	return distance
}
