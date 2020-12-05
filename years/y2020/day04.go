package y2020

import (
	"bytes"
	"strconv"
)

func d4ValidateDate(a []byte, from, to uint64, size int) uint8 {
	if size > 0 {
		if len(a) != size {
			return 0
		}
	}
	i, err := strconv.ParseUint(string(a), 10, 64)
	if err == nil && (from <= i) && (i <= to) {
		return 1
	} else {
		return 0
	}
}

func d4ValidateHeight(a []byte) uint8 {
	i, err := strconv.ParseUint(string(a[:len(a)-2]), 10, 64)
	if err != nil {
		return 0
	}

	switch string(a[len(a)-2:]) {
	case "cm":
		if 150 <= i && i <= 193 {
			return 1
		}
	case "in":
		if 59 <= i && i <= 76 {
			return 1
		}
	}

	return 0
}

func d4ValidatEyeColor(a []byte) uint8 {
	switch string(a) {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return 1
	default:
		return 0
	}
}

func d4ValidateHairColor(a []byte) uint8 {
	if len(a) != 7 || a[0] != '#' || len(bytes.Trim(a[1:], "abcdef0123456789")) != 0 {
		return 0
	} else {
		return 1
	}
}

// Return the hash of the key
func d4hash(p []byte, validate bool) uint8 {
	sw := uint8(0)
	if validate {
		sw = 1
	}
	switch string(p[:4]) {
	case "byr:":
		return 1 << ((1-sw)*1 + sw*d4ValidateDate(p[4:], 1920, 2002, 4)*1)
	case "iyr:":
		return 1 << ((1-sw)*2 + sw*d4ValidateDate(p[4:], 2010, 2020, 4)*2)
	case "eyr:":
		return 1 << ((1-sw)*3 + sw*d4ValidateDate(p[4:], 2020, 2030, 4)*3)
	case "hgt:":
		return 1 << ((1-sw)*4 + sw*d4ValidateHeight(p[4:])*4)
	case "hcl:":
		return 1 << ((1-sw)*5 + sw*d4ValidateHairColor(p[4:])*5)
	case "ecl:":
		return 1 << ((1-sw)*6 + sw*d4ValidatEyeColor(p[4:])*6)
	case "pid:":
		return 1 << ((1-sw)*7 + sw*d4ValidateDate(p[4:], 0, 999999999, 9)*7)
	case "cid:":
		return 0
	default:
		return 1
	}
}

func Day04Part01(input []byte) (string, error) {
	ans := uint64(0)
	for _, a := range bytes.Split(bytes.TrimSpace(input), []byte{'\n', '\n'}) {
		a = bytes.ReplaceAll(a, []byte{'\n'}, []byte{' '})

		// Use bit flags
		f := uint8(0)
		for _, p := range bytes.Split(a, []byte{' '}) {
			f |= d4hash(p, false)
		}
		if f == 0b11111110 {
			ans++
		}
	}
	return strconv.FormatUint(ans, 10), nil
}

func Day04Part02(input []byte) (string, error) {
	ans := uint64(0)
	for _, a := range bytes.Split(bytes.TrimSpace(input), []byte{'\n', '\n'}) {
		a = bytes.ReplaceAll(a, []byte{'\n'}, []byte{' '})

		// Use bit flags
		f := uint8(0)
		for _, p := range bytes.Split(a, []byte{' '}) {
			f |= d4hash(p, true)
		}
		if f == 0b11111110 {
			ans++
		}
	}
	return strconv.FormatUint(ans, 10), nil
}
