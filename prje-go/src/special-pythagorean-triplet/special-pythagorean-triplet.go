package special_pythagorean_triplet

func Run(target int) int {
	for c := range target {
		for b := range target {
			for a := range target {
				if a+b+c == target && (a*a)+(b*b) == (c*c) {
					return a * b * c
				}
			}
		}
	}

	return 0
}
