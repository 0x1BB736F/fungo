package compare

import "github.com/zyltrex/fungo/cast"

func LessUint(i, j interface{}) bool {
	v1 := cast.CastUint(i)
	v2 := cast.CastUint(j)
	return v1 < v2
}

func LessEqualUint(i, j interface{}) bool {
	v1 := cast.CastUint(i)
	v2 := cast.CastUint(j)
	return v1 <= v2
}

func MoreUint(i, j interface{}) bool {
	v1 := cast.CastUint(i)
	v2 := cast.CastUint(j)
	return v1 > v2
}

func MoreEqualUint(i, j interface{}) bool {
	v1 := cast.CastUint(i)
	v2 := cast.CastUint(j)
	return v1 >= v2
}

func EqualUint(i, j interface{}) bool {
	v1 := cast.CastUint(i)
	v2 := cast.CastUint(j)
	return v1 == v2
}

func NotEqualUint(i, j interface{}) bool {
	return Not(EqualUint(i, j))
}

func isNullUint(i interface{}) bool {
	v1 := cast.CastUint(i)
	return v1 == 0
}
