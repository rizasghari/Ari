package lexer

import "github.com/rizasghari/ari/token"

// We only read in integers. What about floats? Or numbers in hex notation? Octal notation?
// We ignore them and just say that Arı doesn’t support this.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// The isLetter helper function just checks whether the given argument is a letter.
// That sounds easy enough, but what’s noteworthy about isLetter is that changing
// this function has a larger impact on the language our interpreter will be able to
// parse than one would expect from such a small function. As you can see, in our case
// it contains the check ch == '_', which means that we’ll treat _ as a letter and allow
// it in identifiers and keywords. That means we can use variable names like foo_bar. Other
// programming languages even allow ! and ? in identifiers. If you want to allow that too,
// this is the place to sneak it in.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
