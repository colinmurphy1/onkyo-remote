package lib

// Ternary operator function
func TOp(condition bool, ifTrue interface{}, ifFalse interface{}) interface{} {
	if condition {
		return ifTrue
	}
	return ifFalse
}
