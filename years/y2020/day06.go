package y2020

import (
	"bytes"
	"strconv"
)

func Day06Part01(input []byte) (string, error) {
	var ans int
	for _, group := range bytes.Split(input, []byte{'\n', '\n'}) {
		s := map[byte]struct{}{}
		for _, c := range group {
			s[c] = struct{}{}
		}
		delete(s, '\n')
		ans += len(s)
	}

	return strconv.Itoa(ans), nil
}

func Day06Part02(input []byte) (string, error) {
	var ans int
	for _, group := range bytes.Split(input, []byte{'\n', '\n'}) {
		s := map[byte]uint64{}
		for _, c := range group {
			s[c]++
		}
		lines := s['\n'] + 1
		for _, v := range s {
			if v == lines {
				ans++
			}
		}
	}

	return strconv.Itoa(ans), nil
}
