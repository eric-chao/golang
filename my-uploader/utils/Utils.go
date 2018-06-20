package utils

func ToStringSlice(params ...interface{}) []string {
	var stringSlice []string
	for _, param := range params {
		stringSlice = append(stringSlice, param.(string))
	}
	// aa := strings.Join(paramSlice, "_")
	// Join 方法第2个参数是 string 而不是 rune
	return stringSlice
}
