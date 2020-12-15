package y2020

import (
	"bytes"
	"strconv"
	"strings"
)

func d15Solve(init []uint64, N uint64) uint64 {
	seen := map[uint64]uint64{}
	for i, x := range init {
		seen[x] = uint64(i)
	}
	nPrev := init[len(init)-1]
	for i := uint64(len(init)); i < N; i++ {
		nPrevPrev, ok := seen[nPrev]
		if !ok {
			nPrevPrev = i - 1
		}
		seen[nPrev] = i - 1
		nPrev = i - 1 - nPrevPrev
	}
	return nPrev
}

func d15FullSolve(input []byte, N uint64) (string, error) {
	raw := strings.Split(string(bytes.TrimSpace(input)), ",")
	nums := make([]uint64, 0, len(raw))
	for _, rs := range raw {
		a, err := strconv.ParseUint(rs, 10, 64)
		if err != nil {
			return "", err
		}
		nums = append(nums, a)
	}

	return strconv.FormatUint(d15Solve(nums, N), 10), nil
}

func Day15Part01(input []byte) (string, error) {
	return d15FullSolve(input, 2020)
}

func Day15Part02(input []byte) (string, error) {
	return d15FullSolve(input, 30000000)
}
