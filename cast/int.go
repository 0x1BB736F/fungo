package cast

func CastInt(i interface{}) int {
	switch i.(type) { //nolint
	case int8:
		return int(i.(int8))
	case int16:
		return int(i.(int16))
	case int32:
		return int(i.(int32))
	case int64:
		return int(i.(int64))
	case int:
		return i.(int)
	default:
		return 0
	}
}

func CastableInt(i interface{}) bool {
	switch i.(type) { //nolint
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	case int:
		return true
	default:
		return false
	}
}
