package math

import "testing"

// func TestAbs(t *testing.T) {
// 	val1, exp1 := 1, 1
// 	val2, exp2 := 4, 4
// 	val3, exp3 := -3, 3

// 	if Abs(val1) != exp1 {
// 		t.Errorf("Expected %d but got %d, for %d", exp1, Abs(val1), val1)
// 	}
// 	if Abs(val2) != exp2 {
// 		t.Errorf("Expected %d but got %d, for %d", exp2, Abs(val2), val2)
// 	}
// 	if Abs(val3) != exp3 {
// 		t.Errorf("Expected %d but got %d,for %d", exp3, Abs(val3), val3)
// 	}

// }

func TestContains(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{"ExistInLists", []int{1, 2, 3, 4, 5}, 3, true},
		{"DoesNotExistLists", []int{1, 2, 3, 4, 5}, 6, false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := Contains(tc.nums, tc.target)
			if result != tc.expected {
				t.Errorf("For %s,expected %v,go %v", tc.name, tc.expected, result)
			}
		})
	}
}
