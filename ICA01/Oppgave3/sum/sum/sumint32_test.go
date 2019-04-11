package sum

import "testing"


var sum_tests_int32 = []struct {
	n1       int32
	n2       int32
	expected int32
}{
	{1, -1, 80},
	{10, 3844239205, 1.5},
	{1, -1, 0},
}

func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}