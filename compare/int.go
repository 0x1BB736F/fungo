package compare

import "github.com/zyltrex/fungo/cast"

func LessInt(i, j interface{}) bool {
	v1 := cast.CastInt(i)
	v2 := cast.CastInt(j)
	return v1 < v2
}

func LessEqualInt(i, j interface{}) bool {
	v1 := cast.CastInt(i)
	v2 := cast.CastInt(j)
	return v1 <= v2
}

func MoreInt(i, j interface{}) bool {
	v1 := cast.CastInt(i)
	v2 := cast.CastInt(j)
	return v1 > v2
}

func MoreEqualInt(i, j interface{}) bool {
	v1 := cast.CastInt(i)
	v2 := cast.CastInt(j)
	return v1 >= v2
}

func EqualInt(i, j interface{}) bool {
	v1 := cast.CastInt(i)
	v2 := cast.CastInt(j)
	return v1 == v2
}

func NotEqualInt(i, j interface{}) bool {
	return Not(EqualInt(i, j))
}

func isNullInt(i interface{}) bool {
	v1 := cast.CastInt(i)
	return v1 == 0
}
