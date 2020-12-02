package y2020

import (
	"regexp"
	"strconv"
)

var d2re = regexp.MustCompile(`(?m)^(\d+)-(\d+) ([a-z]): ([a-z]+)$`)

type d2func func(lo, hi uint64, c byte, pass []byte) bool

func d2solve(input []byte, f d2func) (uint64, error) {
	var ans uint64
	for _, match := range d2re.FindAllSubmatch(input, -1) {
		lo, err := strconv.ParseUint(string(match[1]), 10, 64)
		if err != nil {
			return 0, err
		}
		hi, err := strconv.ParseUint(string(match[2]), 10, 64)
		if err != nil {
			return 0, err
		}

		if f(lo, hi, match[3][0], match[4]) {
			ans++
		}
	}
	return ans, nil
}

func Day02Part01(input []byte) (string, error) {
	ans, err := d2solve(input, func(lo, hi uint64, c byte, pass []byte) bool {
		var cnt uint64
		for _, x := range pass {
			if x == c {
				cnt++
			}
		}
		return lo <= cnt && cnt <= hi
	})
	return strconv.FormatUint(ans, 10), err
}

func Day02Part02(input []byte) (string, error) {
	ans, err := d2solve(input, func(lo, hi uint64, c byte, pass []byte) bool {
		a := pass[lo-1] == c
		b := pass[hi-1] == c

		return (a && !b) || (!a && b)
	})
	return strconv.FormatUint(ans, 10), err
}
