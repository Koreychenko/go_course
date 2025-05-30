package leetcode

import (
	"strconv"
	"strings"
)

func equalPairs(grid [][]int) int {
	hashedI := make(map[string]int, 200)
	hashedJ := make(map[string]int, 200)
	conv1 := strings.Builder{}
	conv2 := strings.Builder{}

	/* Create a hash from values for every row */
	for j := 0; j < len(grid); j++ {
		conv1.Reset()
		conv2.Reset()
		for i := 0; i < len(grid); i++ {
			conv1.WriteString(strconv.Itoa(grid[j][i]))
			conv1.WriteString("_")

			conv2.WriteString(strconv.Itoa(grid[i][j]))
			conv2.WriteString("_")
		}

		hashedI[conv1.String()]++
		hashedJ[conv2.String()]++
	}

	// As we have 2 hashes all we need is to iterate over one of them and check for pairs in the other
	pairs := 0

	for k, v1 := range hashedI {
		if v2, ok := hashedJ[k]; ok {
			pairs += v1 * v2
		}
	}

	return pairs
}

func Leetcode2352(grid [][]int) int {
	return equalPairs(grid)
}
