package lexer

import (
	"monkey_lang/token"
)

type Lexer struct {
	input        string
	position     int  // 所输入字符中的当前位置（指向当前字符）
	readPosition int  // 所输入字符中的当前读取位置（指向当前字符之后的一个字符）
	ch           byte // 当前正在查看的字符
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()

	return l
}

// readChar 把当前字符的指针+1
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// pickChar 读取下一个字符，并且不会移动readPosition指针的位置
func (l *Lexer) pickChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// NextToken 获取一下个词法单元（词元）
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.pickChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = token.NewToken(token.PLUS, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '-':
		tok = token.NewToken(token.MINUS, l.ch)
	case '!':
		if l.pickChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '*':
		tok = token.NewToken(token.ASTERISK, l.ch)
	case '/':
		tok = token.NewToken(token.SLASH, l.ch)
	case '<':
		tok = token.NewToken(token.LT, l.ch)
	case '>':
		tok = token.NewToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// 如果是字符，则按标识符规则读取
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// 如果是数字，则按整形数字读取，暂不支持浮点数
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			// 其它无法识别的情况标识为ILLEGAL
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readIdentifier 读取标识符字符串
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter 判断是否为Identifier支持的字符，当前支持大小写字母与下划线
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || (ch == '_')
}

// skipWhitespace 跳过所有空白字符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
