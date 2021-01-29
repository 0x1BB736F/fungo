package cast

func CastableBool(v interface{}) bool {
	_, ok := v.(bool)
	return ok
}

func CastBool(v interface{}) bool {
	b, ok := v.(bool)
	if !ok {
		return false
	}
	return b
}
