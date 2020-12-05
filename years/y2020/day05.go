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
	var min, max, ans uint64
	min = math.MaxUint64
	for _, l := range bytes.Split(bytes.TrimSpace(input), []byte{'\n'}) {
		if x, err := d5PassToId(l); err != nil {
			return "", err
		} else {
			ans ^= x
			if x < min {
				min = x
			}
			if x > max {
				max = x
			}
		}
	}
	// Since A ^ A = 0, when we reapply all the XORs in the range, the one missing
	// will be the one we are missing who will be left. This is the same thing done in
	// storage RAIDs for parity!
	for i := min; i <= max; i++ {
		ans ^= i
	}

	return strconv.FormatUint(ans, 10), nil
}
