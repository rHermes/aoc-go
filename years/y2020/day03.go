package y2020

import (
	"bytes"
	"strconv"
)

func Day03Part01(input []byte) (string, error) {
	ans := uint64(0)
	for i, line := range bytes.Split(input, []byte{'\n'}) {
		if len(line) > 0 && line[(i*3)%len(line)] == '#' {
			ans++
		}
	}
	return strconv.FormatUint(ans, 10), nil
}

func Day03Part02(input []byte) (string, error) {
	type Slope struct {
		dx int
		dy int
	}
	slopes := map[Slope]uint64{
		Slope{1, 1}: 0,
		Slope{3, 1}: 0,
		Slope{5, 1}: 0,
		Slope{7, 1}: 0,
		Slope{1, 2}: 0,
	}

	for i, line := range bytes.Split(input, []byte{'\n'}) {
		if len(line) == 0 {
			continue
		}

		for k, _ := range slopes {
			y := i / k.dy
			if i%k.dy == 0 && line[(y*k.dx)%len(line)] == '#' {
				slopes[k]++
			}
		}
	}

	ans := uint64(1)
	for _, v := range slopes {
		ans *= v
	}
	return strconv.FormatUint(ans, 10), nil
}
