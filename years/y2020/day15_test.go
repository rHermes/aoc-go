package y2020

import "testing"

func TestDay15Part01(t *testing.T) {
	type TC struct {
		Input         string
		Expected      string
		ExpectedError error
	}

	testCases := []TC{
		{Input: "0,3,6", Expected: "436", ExpectedError: nil},
		{Input: "1,3,2", Expected: "1", ExpectedError: nil},
		{Input: "2,1,3", Expected: "10", ExpectedError: nil},
		{Input: "1,2,3", Expected: "27", ExpectedError: nil},
		{Input: "2,3,1", Expected: "78", ExpectedError: nil},
		{Input: "3,2,1", Expected: "438", ExpectedError: nil},
	}

	for _, tc := range testCases {
		ans, err := Day15Part01([]byte(tc.Input))
		if ans != tc.Expected || err != tc.ExpectedError {
			t.Errorf("Expected \"%s\" for \"%s\" but got \"%s\"\n", tc.Expected, tc.Input, ans)
		}
	}
}
