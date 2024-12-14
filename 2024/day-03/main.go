package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		result = result + parse(line)
	}
}

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	MUL    = "MUL"
	INT    = "INT"
	COMMA  = ","
	LPAREN = "("
	RPAREN = ")"
	DO     = "DO"
	DONT   = "DONT"

	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)

var keywords = map[string]TokenType{
	"mul": MUL,
}

type Lexer struct {
	input           string
	currentPosition int
	nextPosition    int
	ch              byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() Token {
	var tok Token
	switch l.ch {
	case ',':
		tok = newToken(COMMA, l.ch)
	case '(':
		tok = newToken(LPAREN, l.ch)
	case ')':
		tok = newToken(RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	case 'd':
		if l.peekChar() == 'o' {
			l.readChar()
			if l.peekChar() == '(' {
				l.readChar()
				if l.peekChar() == ')' {
					l.readChar()
					tok = Token{Type: DO, Literal: "do()"}
				}
			} else if l.peekChar() == 'n' {
				l.readChar()
				if l.peekChar() == '\'' {
					l.readChar()
					if l.peekChar() == 't' {
						l.readChar()
						if l.peekChar() == '(' {
							l.readChar()
							if l.peekChar() == ')' {
								l.readChar()
								tok = Token{Type: DONT, Literal: "don't()"}
							}
						}
					}
				}
			}
		}
	case 'm':
		if l.peekChar() == 'u' {
			l.readChar()
			if l.peekChar() == 'l' {
				l.readChar()
				tok = Token{Type: MUL, Literal: "mul"}
			}
		}
	default:
		if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) readNumber() string {
	position := l.currentPosition
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.currentPosition
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.currentPosition]
}

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9')
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func newToken(t TokenType, b byte) Token {
	return Token{Type: t, Literal: string(b)}
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.currentPosition = l.nextPosition
	l.nextPosition++
}

var disable = false

func parse(program string) int64 {
	l := New(program)
	tokens := make([]Token, 0)
	i := 0
	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == EOF {
			break
		}
		i++
	}

	curr := 0
	result := int64(0)
	for true {
		currToken := tokens[curr]
		if currToken.Type == MUL {
			if curr+5 >= len(tokens) {
				break
			}

			isValid := tokens[curr+1].Type == LPAREN
			isValid = tokens[curr+2].Type == INT
			isValid = tokens[curr+3].Type == COMMA
			isValid = tokens[curr+4].Type == INT
			isValid = tokens[curr+5].Type == RPAREN
			if !isValid || disable {
				curr++
				continue
			}

			leftInt, _ := strconv.Atoi(tokens[curr+2].Literal)
			rightInt, _ := strconv.Atoi(tokens[curr+4].Literal)
			result += int64(leftInt) * int64(rightInt)
		} else if currToken.Type == DONT {
			disable = true
		} else if currToken.Type == DO {
			disable = false
		} else if currToken.Type == EOF {
			break
		}
		curr++
	}

	return result
}
