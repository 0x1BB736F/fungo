package cast

func CastFloat(i interface{}) float64 {
	switch i.(type) { //nolint
	case float32:
		return float64(i.(float32))
	case float64:
		return i.(float64)
	default:
		return 0.0
	}
}

func CastableFloat(i interface{}) bool {
	switch i.(type) { //nolfloat
	case float32:
		return true
	case float64:
		return true
	default:
		return false
	}
}
