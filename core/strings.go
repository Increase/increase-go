package core

func CoalesceStrings[P ~string | ~*string](strings ...P) string {
	for _, str := range strings {
		if str, ok := any(str).(string); ok && len(str) != 0 {
			return str
		}
		if str, ok := any(str).(*string); ok && len(*str) != 0 {
			return *str
		}
	}
	return ""
}
