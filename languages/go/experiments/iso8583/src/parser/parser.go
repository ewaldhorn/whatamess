package parser

type MessageTypeIndicator struct {
	MTI string
}

// readMTIFromString returns the MTI (first 4 characters) from the source message string.
//
// assumes the string is at least 4 characters long.
func readMTIFromString(source string) MessageTypeIndicator {
	return MessageTypeIndicator{MTI: source[0:4]}
}
