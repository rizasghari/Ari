Lexing
Lexer or Tokenizer or Scanner

"let x = 5 + 5;"
[
    LET,
    IDENTIFIER("x"),
    EQUAL_SIGN,
    INTEGER(5),
    PLUS_SIGN,
    INTEGER(5),
    SEMICOLON
]


Features and limitations:
-  Ari only supports ASCII characters instead of the full Unicode range