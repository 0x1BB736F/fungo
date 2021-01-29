package cast

func CastString(v interface{}) string {
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return s
}

func CastableString(v interface{}) bool {
	_, ok := v.(string)
	return ok
}
