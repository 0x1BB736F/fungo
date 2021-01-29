package compare

import "github.com/zyltrex/fungo/cast"

func IsNull(v interface{}) bool {
	switch {
	case cast.CastableUint(v):
		return isNullUint(v)
	case cast.CastableInt(v):
		return isNullInt(v)
	case cast.CastableFloat(v):
		return isNullFloat(v)
	case cast.CastableString(v):
		return isNullString(v)
	case cast.CastableBool(v):
		return isNullBool(v)
	default:
		return false
	}
}

func NotNull(v interface{}) bool {
	return Not(IsNull(v))
}
