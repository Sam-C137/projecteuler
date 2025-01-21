package collatz_sequence

func Run(upperbound int) int {
	maxSeq := 0
	maxOcc := 0

	for n := range upperbound {
		if newSeq := seq(n); newSeq > maxSeq {
			maxSeq = newSeq
			maxOcc = n
		}
	}

	return maxOcc
}

func seq(start int) int {
	count := 1

	for start > 1 {
		if (start & 1) == 0 {
			start /= 2
		} else {
			start = 3*start + 1
		}
		count++
	}

	return count
}
