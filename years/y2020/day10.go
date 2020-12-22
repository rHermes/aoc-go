package y2020

import (
	"bytes"
	"sort"
	"strconv"

	"github.com/rhermes/aoc-go/utils"
)

func Day10Part01(input []byte) (string, error) {
	ints, err := utils.LinesToUint64(bytes.NewReader(input))
	if err != nil {
		return "", err
	}
	sort.Slice(ints, func(i, j int) bool { return ints[i] < ints[j] })

	prevs := uint64(0)
	diffs := []uint64{0, 0, 1}
	for _, k := range ints {
		diffs[k-prevs-1]++
		prevs = k
	}
	return strconv.FormatUint(diffs[0]*diffs[2], 10), nil
}

func Day10Part02(input []byte) (string, error) {
	ints, err := utils.LinesToUint64(bytes.NewReader(input))
	if err != nil {
		return "", err
	}
	sort.Slice(ints, func(i, j int) bool { return ints[i] < ints[j] })

	var a, b, c, df, last uint64
	switch ints[0] {
	case 3:
		a = 1
	case 2:
		b = 1
	case 1:
		c = 1
	}
	for _, x := range ints {
		df, last = x-last, x
		switch df {
		case 1:
			a, b, c = b, c, a+b+c
		case 2:
			a, b, c = c, 0, a+b+c
		default:
			a, b, c = 0, 0, c
		}
	}
	return strconv.FormatUint(c, 10), nil
}
