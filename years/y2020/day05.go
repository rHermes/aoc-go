package y2020

import (
	"bytes"
	"errors"
	"math"
	"strconv"
)

func d5PassToId(input []byte) (uint64, error) {
	if len(input) != 10 {
		return 0, errors.New("invalid input")
	}

	ans := uint64(0)
	for i, c := range input {
		switch c {
		case 'B', 'R':
			ans |= 1 << (9 - i)
		case 'F', 'L':
		default:
			return 0, errors.New("invalid input")
		}
	}
	return ans, nil
}

func Day05Part01(input []byte) (string, error) {
	var ans uint64
	for _, l := range bytes.Split(bytes.TrimSpace(input), []byte{'\n'}) {
		if x, err := d5PassToId(l); err != nil {
			return "", err
		} else {
			if x > ans {
				ans = x
			}
		}
	}
	return strconv.FormatUint(ans, 10), nil
}

func Day05Part02(input []byte) (string, error) {
	var min, max, inperf, perf uint64
	min = math.MaxUint64
	for _, l := range bytes.Split(bytes.TrimSpace(input), []byte{'\n'}) {
		if x, err := d5PassToId(l); err != nil {
			return "", err
		} else {
			inperf ^= x
			if x < min {
				min = x
			}
			if x > max {
				max = x
			}
		}
	}

	// We find the missing number by assuming it's the only missing
	// and taking the xor of the perfect series and the xor of the missing
	// series and seeing what comes out. This is like doing a parity check
	// in a raid.
	for i := min; i <= max; i++ {
		perf ^= i
	}

	return strconv.FormatUint(inperf^perf, 10), nil
}
