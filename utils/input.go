package utils

import (
	"bufio"
	"io"
	"strconv"
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
