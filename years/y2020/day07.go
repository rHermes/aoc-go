package y2020

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type d7bags map[string]map[string]uint64

func d7ParseBags(input []byte) (d7bags, error) {
	ans := map[string]map[string]uint64{}
	for _, line := range bytes.Split(bytes.TrimSpace(input), []byte{'\n'}) {
		prts := bytes.SplitN(line[:len(line)-1], []byte(" bags contain "), 2)
		if len(prts) != 2 {
			return nil, errors.New("invalid format")
		}
		name, rest := string(prts[0]), string(prts[1])

		xs := map[string]uint64{}
		ans[name] = xs
		if rest == "no other bags" {
			continue
		}

		for _, ys := range strings.Split(rest, ", ") {
			parts := strings.Split(ys, " ")
			if len(parts) < 3 {
				return nil, errors.New("invalid format")
			}
			num, err := strconv.ParseUint(parts[0], 10, 64)
			if err != nil {
				return nil, err
			}

			name := strings.Join(parts[1:len(parts)-1], " ")

			xs[name] = num
		}
	}
	return ans, nil
}

// we pass a cache along side here, to avoid recomputing values we have looked
// at before. This is not needed for such a small input, but it has much better
// runtime for larger inputs.
func d7Contains(cache map[string]bool, bags d7bags, bag string) bool {
	ans, ok := cache[bag]
	if ok {
		return ans
	}

	for k := range bags[bag] {
		if d7Contains(cache, bags, k) {
			ans = true
			break
		}
	}

	cache[bag] = ans
	return ans
}

// we pass a cache along side here, to avoid recomputing values we have looked
// at before. This is not needed for such a small input, but it has much better
// runtime for larger inputs.
func d7Required(cache map[string]uint64, bags d7bags, bag string) uint64 {
	ans, ok := cache[bag]
	if ok {
		return ans
	}

	for sub, num := range bags[bag] {
		ans += num * (1 + d7Required(cache, bags, sub))
	}

	cache[bag] = ans
	return ans
}

func Day07Part01(input []byte) (string, error) {
	bags, err := d7ParseBags(input)
	if err != nil {
		return "", err
	}
	ans := uint64(0)
	cache := map[string]bool{
		"shiny gold": true,
	}
	for k := range bags {
		if d7Contains(cache, bags, k) {
			ans++
		}
	}
	return strconv.FormatUint(ans-1, 10), nil
}

func Day07Part02(input []byte) (string, error) {
	bags, err := d7ParseBags(input)
	if err != nil {
		return "", err
	}
	cache := map[string]uint64{}
	return strconv.FormatUint(d7Required(cache, bags, "shiny gold"), 10), nil
}
