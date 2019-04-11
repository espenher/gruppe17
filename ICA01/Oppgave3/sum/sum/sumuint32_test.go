package sum

import "testing"


var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{1, 1, 2},
	{10, -5, 15},
	{1, -1, 2},
}

func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}