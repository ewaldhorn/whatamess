package parser

type Iso8583VersionType string
type MessageClassType string

// Iso8583 version types.
// Types 3..7 are reserved by ISO.
const (
	Iso8583_1987       Iso8583VersionType = "0"
	Iso8583_1993       Iso8583VersionType = "1"
	Iso8583_2002       Iso8583VersionType = "2"
	Iso8583_Reserved_3 Iso8583VersionType = "3"
	Iso8583_Reserved_4 Iso8583VersionType = "4"
	Iso8583_Reserved_5 Iso8583VersionType = "5"
	Iso8583_Reserved_6 Iso8583VersionType = "6"
	Iso8583_Reserved_7 Iso8583VersionType = "7"
	Iso8583_National   Iso8583VersionType = "8"
	Iso8583_Private    Iso8583VersionType = "9"
)

const (
	MessageClass_Reserved_0     MessageClassType = "0"
	MessageClass_Authorization  MessageClassType = "1"
	MessageClass_Financial      MessageClassType = "2"
	MessageClass_FileActions    MessageClassType = "3"
	MessageClass_Reversal       MessageClassType = "4"
	MessageClass_Reconciliation MessageClassType = "5"
	MessageClass_Administrative MessageClassType = "6"
	MessageClass_FeeCollection  MessageClassType = "7"
	MessageClass_Network        MessageClassType = "8"
	MessageClass_Reserved_9     MessageClassType = "9"
)

type MessageTypeIndicator struct {
	MTI string
}

// readMTIFromString returns the MTI (first 4 characters) from the source message string.
//
// assumes the string is at least 4 characters long.
func readMTIFromString(source string) MessageTypeIndicator {
	return MessageTypeIndicator{MTI: source[0:4]}
}
