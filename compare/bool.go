package compare

import "github.com/zyltrex/fungo/cast"

func Not(b bool) bool {
	return !b
}

func EqualBool(i, v bool) bool {
	return i == v
}

func NotEqualBool(i, v bool) bool {
	return Not(EqualBool(i, v))
}

func isNullBool(v interface{}) bool {
	b := cast.CastBool(v)
	return Not(b)
}
