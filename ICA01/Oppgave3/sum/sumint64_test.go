package sum

import "testing"


var sum_tests_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	{1, "dksld", 2.0},
	{10, -5, -5},
	{1, 57238472947289, 6},
}

func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}