package gen

func asString(node interface{}, defaultValue string) string {
	if node == nil {
		return defaultValue
	}
	switch node.(type) {
	case string:
		return node.(string)
	case []uint8:
		return string(node.([]uint8))
	}
	return defaultValue
}
