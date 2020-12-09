package y2020

import (
	"bytes"
	"errors"
	"math"
	"strconv"

	"github.com/rhermes/aoc-go/utils"
)

func d9Solve1(nums []uint64, N uint64) (uint64, error) {
	if uint64(len(nums)) < N {
		return 0, errors.New("Invalid input")
	}

outer:
	for i := N; i < uint64(len(nums)); i++ {
		// Create the "set". We want to use the struct{} type as it takes less space
		items := make(map[uint64]struct{}, N)
		for _, x := range nums[i-N : i] {
			items[x] = struct{}{}
		}

		for k := range items {
			if nums[i] > k && nums[i]-k != k {
				if _, ok := items[nums[i]-k]; ok {
					continue outer
				}
			}
		}

		return nums[i], nil
	}

	return 0, errors.New("No answer found!")
}

func Day09Part01(input []byte) (string, error) {
	nums, err := utils.LinesToUint64(bytes.NewReader(bytes.TrimSpace(input)))
	if err != nil {
		return "", err
	}
	sol, err := d9Solve1(nums, 25)
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(sol, 10), nil
}

func Day09Part02(input []byte) (string, error) {
	nums, err := utils.LinesToUint64(bytes.NewReader(bytes.TrimSpace(input)))
	if err != nil {
		return "", err
	}

	// first solve 1.
	target, err := d9Solve1(nums, 25)
	if err != nil {
		return "", err
	}

	var sum uint64
	var back uint64
	for front, x := range nums {
		sum += x

		for sum > target {
			sum -= nums[back]
			back++
		}

		if sum == target {
			var min, max uint64
			min = math.MaxUint64
			for _, y := range nums[back:front] {
				if y < min {
					min = y
				}
				if y > max {
					max = y
				}
			}

			return strconv.FormatUint(min+max, 10), nil
		}
	}

	return "", errors.New("No solution found")
}
