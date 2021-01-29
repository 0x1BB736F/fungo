package cast

func CastUint(i interface{}) uint {
	switch i.(type) { //nolint
	case uint8:
		return uint(i.(uint8))
	case uint16:
		return uint(i.(uint16))
	case uint32:
		return uint(i.(uint32))
	case uint64:
		return uint(i.(uint64))
	case uint:
		return i.(uint)
	default:
		return 0
	}
}

func CastableUint(i interface{}) bool {
	switch i.(type) { //nolint
	case uint8:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	case uint:
		return true
	default:
		return false
	}
}
