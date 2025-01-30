package maximum_path_sum_1

import (
	"strconv"
	"strings"
)

func Run(input string) int {
	tree := parseInput(input)
	dp := make([]int, len(tree[len(tree)-1]))
	copy(dp, tree[len(tree)-1])

	for i := len(tree) - 2; i >= 0; i-- {
		for j := 0; j < len(tree[i]); j++ {
			dp[j] = tree[i][j] + max(dp[j], dp[j+1])
		}
	}

	return dp[0]
}

func parseInput(input string) [][]int {
	s := strings.Split(input, "\n")
	out := make([][]int, len(s), len(s))

	for i := 0; i < len(s); i++ {
		split := strings.Split(s[i], " ")
		for _, char := range split {
			n, _ := strconv.Atoi(char)
			out[i] = append(out[i], n)
		}
	}

	return out
}
