package cond

func Interface(condition bool, v interface{}, x interface{}) interface{} {
	if condition {
		return v
	}

	return x
}

func String(condition bool, v string, x string) string {
	if condition {
		return v
	}

	return x
}

func Bool(condition bool, v bool, x bool) bool {
	if condition {
		return v
	}

	return x
}

func Int(condition bool, v int, x int) int {
	if condition {
		return v
	}

	return x
}

func Int8(condition bool, v int8, x int8) int8 {
	if condition {
		return v
	}

	return x
}

func Int16(condition bool, v int16, x int16) int16 {
	if condition {
		return v
	}

	return x
}

func Int32(condition bool, v int32, x int32) int32 {
	if condition {
		return v
	}

	return x
}

func Int64(condition bool, v int64, x int64) int64 {
	if condition {
		return v
	}

	return x
}

func Uint(condition bool, v uint, x uint) uint {
	if condition {
		return v
	}

	return x
}

func Uint8(condition bool, v uint8, x uint8) uint8 {
	if condition {
		return v
	}

	return x
}

func Uint16(condition bool, v uint16, x uint16) uint16 {
	if condition {
		return v
	}

	return x
}

func Uint32(condition bool, v uint32, x uint32) uint32 {
	if condition {
		return v
	}

	return x
}

func Uint64(condition bool, v uint64, x uint64) uint64 {
	if condition {
		return v
	}

	return x
}
