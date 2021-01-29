package compare

import "github.com/zyltrex/fungo/cast"

func LessFloat(i, j interface{}) bool {
	v1 := cast.CastFloat(i)
	v2 := cast.CastFloat(j)
	return v1 < v2
}

func LessEqualFloat(i, j interface{}) bool {
	v1 := cast.CastFloat(i)
	v2 := cast.CastFloat(j)
	return v1 <= v2
}

func MoreFloat(i, j interface{}) bool {
	v1 := cast.CastFloat(i)
	v2 := cast.CastFloat(j)
	return v1 > v2
}

func MoreEqualFloat(i, j interface{}) bool {
	v1 := cast.CastFloat(i)
	v2 := cast.CastFloat(j)
	return v1 >= v2
}

func EqualFloat(i, j interface{}) bool {
	v1 := cast.CastFloat(i)
	v2 := cast.CastFloat(j)
	return v1 == v2
}

func NotEqualFloat(i, j interface{}) bool {
	return Not(EqualFloat(i, j))
}

func isNullFloat(i interface{}) bool {
	v1 := cast.CastFloat(i)
	return v1 == 0.0
}
