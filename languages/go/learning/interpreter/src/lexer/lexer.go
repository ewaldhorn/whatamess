package lexer

// ----------------------------------------------------------------------------
type Lexer struct {
	input        string
	position     int  // current pos in input, points to current char
	readPosition int  // current reading position in input, after current char
	ch           byte // current char being inspected
}

// ----------------------------------------------------------------------------
// New creates a new lexer that will process the input string.
// Returns a pointer to the newly created lexer.
func New(input string) *Lexer {
	newLexer := &Lexer{input: input}
	newLexer.readChar()

	return newLexer
}

// ----------------------------------------------------------------------------
// readChar reads the next character in the input and advances l.position and
// l.readPosition. When it reaches the end of input, it sets l.ch to 0 (NUL).
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
