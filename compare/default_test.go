package compare

import "testing"

func Test_IsNull(t *testing.T) {
	type test struct {
		val      interface{}
		expected bool
	}

	tt := []test{
		{int(1), false},
		{int(0), true},
		{uint(1), false},
		{uint(0), true},
		{float32(1), false},
		{float32(0), true},
		{byte(1), false},
		{byte(0), true},
	}

	for _, elem := range tt {
		res := IsNull(elem.val)
		if res != elem.expected {
			t.Errorf("%v %v %v error", elem.val, elem.expected, res)
		}
	}
}

func Benchmark_IsNull(b *testing.B) {
	v1 := int8(1)
	v2 := uint8(1)
	for i := 0; i < b.N; i++ {
		IsNull(v1)
		IsNull(v2)
	}
}
