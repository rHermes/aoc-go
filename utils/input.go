package utils

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// LinesToInt64 convert lines to int64
func LinesToInt64(rdr io.Reader) ([]int64, error) {
	sc := bufio.NewScanner(rdr)
	nums := make([]int64, 0, 100)
	for sc.Scan() {
		num, err := strconv.ParseInt(sc.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, sc.Err()
}

// LinesToUint64 convert lines to uint64
func LinesToUint64(rdr io.Reader) ([]uint64, error) {
	sc := bufio.NewScanner(rdr)
	nums := make([]uint64, 0, 100)
	for sc.Scan() {
		num, err := strconv.ParseUint(sc.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, sc.Err()
}

// LineToUint64s converts a line to uint64s, seperated by some seperator
func LineToUint64s(line, sep string) ([]uint64, error) {
	raw := strings.Split(line, sep)
	nums := make([]uint64, 0, len(raw))
	for _, rs := range raw {
		a, err := strconv.ParseUint(rs, 10, 64)
		if err != nil {
			return nil, err
		}
		nums = append(nums, a)
	}
	return nums, nil
}
