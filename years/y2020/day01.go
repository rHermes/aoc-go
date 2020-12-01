package y2020

import (
	"errors"
	"io"
	"strconv"

	"github.com/rhermes/aoc-go/utils"
)

func d1solve(nums []int64, target, left int64) (int64, bool) {
	for i, x := range nums {
		nx := target - x
		if nx == 0 && left == 1 {
			return x, true
		}
		if nx > 0 && left > 1 {
			dx, ok := d1solve(nums[i:], nx, left-1)
			if ok {
				return x * dx, true
			}
		}
	}
	return -1, false
}

func Day01Part01(rdr io.Reader) (string, error) {
	nums, err := utils.LinesToInt64(rdr)
	if err != nil {
		return "", err
	}

	ans, ok := d1solve(nums, 2020, 2)
	if !ok {
		return "", errors.New("No valid solution")
	}
	return strconv.FormatInt(ans, 10), nil
}

func Day01Part02(rdr io.Reader) (string, error) {
	nums, err := utils.LinesToInt64(rdr)
	if err != nil {
		return "", err
	}

	ans, ok := d1solve(nums, 2020, 3)
	if !ok {
		return "", errors.New("No valid solution")
	}
	return strconv.FormatInt(ans, 10), nil
}
