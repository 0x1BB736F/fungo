package cast

import "testing"

func Test_CastableUint(t *testing.T) {
	type test struct {
		val      interface{}
		expected bool
	}

	var tt = []test{
		{uint8(1), true},
		{uint16(1), true},
		{uint32(1), true},
		{uint64(1), true},
		{uint(1), true},
		{string("qwerty"), false},
		{int(1), false},
	}

	for _, elem := range tt {
		res := CastableUint(elem.val)
		if res != elem.expected {
			t.Errorf("%v %v %v error", elem.val, res, elem.expected)
		}
	}
}

func Test_CastUint(t *testing.T) {
	type test struct {
		val      interface{}
		expected interface{}
	}

	var tt = []test{
		{uint8(1), uint(1)},
		{uint16(1), uint(1)},
		{uint32(1), uint(1)},
		{uint64(1), uint(1)},
		{uint(1), uint(1)},
		{string("qwerty"), uint(0)},
		{int(1), uint(0)},
	}

	for _, elem := range tt {
		res := CastUint(elem.val)
		if res != elem.expected {
			t.Errorf("%v %v %v error", elem.val, res, elem.expected)
		}
	}
}
