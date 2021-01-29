package cast

import "testing"

func Test_CastableInt(t *testing.T) {
	type test struct {
		val      interface{}
		expected bool
	}

	var tt = []test{
		{int8(1), true},
		{int16(1), true},
		{int32(1), true},
		{int64(1), true},
		{int(1), true},
		{string("qwerty"), false},
		{uint(1), false},
	}

	for _, elem := range tt {
		res := CastableInt(elem.val)
		if res != elem.expected {
			t.Errorf("%v %v %v error", elem.val, res, elem.expected)
		}
	}
}

func Test_CastInt(t *testing.T) {
	type test struct {
		val      interface{}
		expected interface{}
	}

	var tt = []test{
		{int8(1), int(1)},
		{int16(1), int(1)},
		{int32(1), int(1)},
		{int64(1), int(1)},
		{int(1), int(1)},
		{string("qwerty"), int(0)},
		{uint(1), int(0)},
	}

	for _, elem := range tt {
		res := CastInt(elem.val)
		if res != elem.expected {
			t.Errorf("%v %v %v error", elem.val, res, elem.expected)
		}
	}
}
