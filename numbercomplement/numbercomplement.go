package numbercomplement

func findComplement(num int) int {
	mask := ^0

	for num&mask != 0 {
		mask = mask << 1
	}

	mask = ^mask

	return num ^ mask
}
