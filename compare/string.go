package compare

import "github.com/zyltrex/fungo/cast"

func EqualString(i, v interface{}) bool {
	s1 := cast.CastString(i)
	s2 := cast.CastString(v)
	return s1 == s2
}

func NotEqualString(i, v interface{}) bool {
	return Not(EqualString(i, v))
}

func isNullString(v interface{}) bool {
	s := cast.CastString(v)
	return s == ""
}
