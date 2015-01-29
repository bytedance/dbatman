// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import (
	"bytes"
	"fmt"
	"strings"

	"micode.be.xiaomi.com/wangjing1/go-mysql-proxy/sqltypes"
)

const EOFCHAR = 0x100

// Tokenizer is the struct used to generate SQL
// tokens for the parser.
type Tokenizer struct {
	InStream      *strings.Reader
	AllowComments bool
	ForceEOF      bool
	lastChar      uint16
	Position      int
	errorToken    []byte
	LastError     string
	posVarIndex   int
	ParseTree     Statement
}

// NewStringTokenizer creates a new Tokenizer for the
// sql string.
func NewStringTokenizer(sql string) *Tokenizer {
	return &Tokenizer{InStream: strings.NewReader(sql)}
}

var keywords = map[string]int{
	"all":           ALL,
	"alter":         ALTER,
	"analyze":       ANALYZE,
	"and":           AND,
	"as":            AS,
	"asc":           ASC,
	AST_AUTO_INCR:   AUTO_INCREMENT,
	"between":       BETWEEN,
	AST_BIGINT:      BIGINT,
	"blob":          BLOB,
	"bit":           BIT,
	"bool":          BOOL,
	"by":            BY,
	"case":          CASE,
	AST_CHAR:        CHAR,
	AST_CHARACTER:   CHARACTER,
	"columns":       COLUMNS,
	"create":        CREATE,
	"cross":         CROSS,
	"database":      DATABASE,
	"databases":     DATABASES,
	"date":          DATE,
	"datetime":      DATETIME,
	"decimal":       DECIMAL,
	"default":       DEFAULT,
	"delete":        DELETE,
	"desc":          DESC,
	"describe":      DESCRIBE,
	"distinct":      DISTINCT,
	"drop":          DROP,
	"double":        DOUBLE,
	"duplicate":     DUPLICATE,
	"else":          ELSE,
	"end":           END,
	"enum":          ENUM,
	AST_ERRORS:      ERRORS,
	"except":        EXCEPT,
	"exists":        EXISTS,
	"explain":       EXPLAIN,
	"float":         FLOAT,
	"for":           FOR,
	"force":         FORCE,
	"from":          FROM,
	"fulltext":      FULLTEXT,
	"function":      FUNCTION,
	"global":        GLOBAL,
	"group":         GROUP,
	"having":        HAVING,
	"if":            IF,
	"ignore":        IGNORE,
	"in":            IN,
	AST_INT:         INT,
	"index":         INDEX,
	"indexes":       INDEXES,
	"inner":         INNER,
	"insert":        INSERT,
	"integer":       INTEGER,
	"intersect":     INTERSECT,
	"into":          INTO,
	"is":            IS,
	"join":          JOIN,
	"key":           KEY,
	"keys":          KEYS,
	"left":          LEFT,
	"like":          LIKE,
	"limit":         LIMIT,
	"lock":          LOCK,
	"long":          LONG,
	"longblob":      LONGBLOB,
	"longtext":      LONGTEXT,
	"mediumblob":    MEDIUMBLOB,
	"mediumint":     MEDIUMINT,
	"mediumtext":    MEDIUMTEXT,
	"minus":         MINUS,
	"names":         NAMES,
	"natural":       NATURAL,
	"not":           NOT,
	"null":          NULL,
	"numeric":       NUMERIC,
	"on":            ON,
	"or":            OR,
	"order":         ORDER,
	"outer":         OUTER,
	"precision":     PRECISION,
	AST_PRIMARY:     PRIMARY,
	"procedure":     PROCEDURE,
	"real":          REAL,
	"rename":        RENAME,
	"replace":       REPLACE,
	"right":         RIGHT,
	AST_STATUS:      STATUS,
	"select":        SELECT,
	"session":       SESSION,
	"set":           SET,
	"show":          SHOW,
	"signed":        SIGNED,
	AST_SMALLINT:    SMALLINT,
	"straight_join": STRAIGHT_JOIN,
	"table":         TABLE,
	"tables":        TABLES,
	AST_TEMPORARY:   TEMPORARY,
	"text":          TEXT,
	AST_TIME:        TIME,
	AST_TIMESTAMP:   TIMESTAMP,
	"tinyblob":      TINYBLOB,
	AST_TINYINT:     TINYINT,
	"tinytext":      TINYTEXT,
	"then":          THEN,
	"to":            TO,
	"union":         UNION,
	"unique":        UNIQUE,
	"unsigned":      UNSIGNED,
	"update":        UPDATE,
	"use":           USE,
	"using":         USING,
	"varchar":       VARCHAR,
	"values":        VALUES,
	"value":         VALUE,
	"varbinary":     VARBINARY,
	"variables":     VARIABLES,
	"view":          VIEW,
	AST_WARNINGS:    WARNINGS,
	"when":          WHEN,
	"where":         WHERE,
	"year":          YEAR,
	"begin":         BEGIN,
	"rollback":      ROLLBACK,
	"commit":        COMMIT,

	//for proxy admin
	"admin": ADMIN,
	"proxy": PROXY,

	//special
}

// Lex returns the next token form the Tokenizer.
// This function is used by go yacc.
func (tkn *Tokenizer) Lex(lval *yySymType) int {
	typ, val := tkn.Scan()
	for typ == COMMENT {
		if tkn.AllowComments {
			break
		}
		typ, val = tkn.Scan()
	}
	switch typ {
	case ID, STRING, NUMBER, VALUE_ARG, LIST_ARG, COMMENT:
		lval.bytes = val
	}
	tkn.errorToken = val
	return typ
}

// Error is called by go yacc if there's a parsing error.
func (tkn *Tokenizer) Error(err string) {
	buf := bytes.NewBuffer(make([]byte, 0, 32))
	if tkn.errorToken != nil {
		fmt.Fprintf(buf, "%s at position %v near %s", err, tkn.Position, tkn.errorToken)
	} else {
		fmt.Fprintf(buf, "%s at position %v", err, tkn.Position)
	}
	tkn.LastError = buf.String()
}

// Scan scans the tokenizer for the next token and returns
// the token type and an optional value.
func (tkn *Tokenizer) Scan() (int, []byte) {
	if tkn.ForceEOF {
		return 0, nil
	}

	if tkn.lastChar == 0 {
		tkn.next()
	}
	tkn.skipBlank()
	switch ch := tkn.lastChar; {
	case isLetter(ch):
		return tkn.scanIdentifier()
	case isDigit(ch):
		return tkn.scanNumber(false)
	case ch == ':':
		return tkn.scanBindVar()
	default:
		tkn.next()
		switch ch {
		case EOFCHAR:
			return 0, nil
		case '=', ',', ';', '(', ')', '+', '*', '%', '&', '|', '^', '~':
			return int(ch), nil
		case '?':
			tkn.posVarIndex++
			buf := new(bytes.Buffer)
			fmt.Fprintf(buf, ":v%d", tkn.posVarIndex)
			return VALUE_ARG, buf.Bytes()
		case '.':
			if isDigit(tkn.lastChar) {
				return tkn.scanNumber(true)
			} else {
				return int(ch), nil
			}
		case '/':
			switch tkn.lastChar {
			case '/':
				tkn.next()
				return tkn.scanCommentType1("//")
			case '*':
				tkn.next()
				return tkn.scanCommentType2()
			default:
				return int(ch), nil
			}
		case '-':
			if tkn.lastChar == '-' {
				tkn.next()
				return tkn.scanCommentType1("--")
			} else {
				return int(ch), nil
			}
		case '<':
			switch tkn.lastChar {
			case '>':
				tkn.next()
				return NE, nil
			case '=':
				tkn.next()
				switch tkn.lastChar {
				case '>':
					tkn.next()
					return NULL_SAFE_EQUAL, nil
				default:
					return LE, nil
				}
			default:
				return int(ch), nil
			}
		case '>':
			if tkn.lastChar == '=' {
				tkn.next()
				return GE, nil
			} else {
				return int(ch), nil
			}
		case '!':
			if tkn.lastChar == '=' {
				tkn.next()
				return NE, nil
			} else {
				return LEX_ERROR, []byte("!")
			}
		case '\'', '"':
			return tkn.scanString(ch, STRING)
		case '`':
			return tkn.scanLiteralIdentifier()
		default:
			return LEX_ERROR, []byte{byte(ch)}
		}
	}
}

func (tkn *Tokenizer) skipBlank() {
	ch := tkn.lastChar
	for ch == ' ' || ch == '\n' || ch == '\r' || ch == '\t' {
		tkn.next()
		ch = tkn.lastChar
	}
}

func (tkn *Tokenizer) scanIdentifier() (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteByte(byte(tkn.lastChar))
	for tkn.next(); isLetter(tkn.lastChar) || isDigit(tkn.lastChar); tkn.next() {
		buffer.WriteByte(byte(tkn.lastChar))
	}
	lowered := bytes.ToLower(buffer.Bytes())
	if keywordId, found := keywords[string(lowered)]; found {
		return keywordId, lowered
	}
	return ID, buffer.Bytes()
}

func (tkn *Tokenizer) scanLiteralIdentifier() (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteByte(byte(tkn.lastChar))
	if !isLetter(tkn.lastChar) {
		return LEX_ERROR, buffer.Bytes()
	}
	for tkn.next(); isLetter(tkn.lastChar) || isDigit(tkn.lastChar); tkn.next() {
		buffer.WriteByte(byte(tkn.lastChar))
	}
	if tkn.lastChar != '`' {
		return LEX_ERROR, buffer.Bytes()
	}
	tkn.next()
	return ID, buffer.Bytes()
}

func (tkn *Tokenizer) scanBindVar() (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteByte(byte(tkn.lastChar))
	token := VALUE_ARG
	tkn.next()
	if tkn.lastChar == ':' {
		token = LIST_ARG
		buffer.WriteByte(byte(tkn.lastChar))
		tkn.next()
	}
	if !isLetter(tkn.lastChar) {
		return LEX_ERROR, buffer.Bytes()
	}
	for isLetter(tkn.lastChar) || isDigit(tkn.lastChar) || tkn.lastChar == '.' {
		buffer.WriteByte(byte(tkn.lastChar))
		tkn.next()
	}
	return token, buffer.Bytes()
}

func (tkn *Tokenizer) scanMantissa(base int, buffer *bytes.Buffer) {
	for digitVal(tkn.lastChar) < base {
		tkn.ConsumeNext(buffer)
	}
}

func (tkn *Tokenizer) scanNumber(seenDecimalPoint bool) (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	if seenDecimalPoint {
		buffer.WriteByte('.')
		tkn.scanMantissa(10, buffer)
		goto exponent
	}

	if tkn.lastChar == '0' {
		// int or float
		tkn.ConsumeNext(buffer)
		if tkn.lastChar == 'x' || tkn.lastChar == 'X' {
			// hexadecimal int
			tkn.ConsumeNext(buffer)
			tkn.scanMantissa(16, buffer)
		} else {
			// octal int or float
			seenDecimalDigit := false
			tkn.scanMantissa(8, buffer)
			if tkn.lastChar == '8' || tkn.lastChar == '9' {
				// illegal octal int or float
				seenDecimalDigit = true
				tkn.scanMantissa(10, buffer)
			}
			if tkn.lastChar == '.' || tkn.lastChar == 'e' || tkn.lastChar == 'E' {
				goto fraction
			}
			// octal int
			if seenDecimalDigit {
				return LEX_ERROR, buffer.Bytes()
			}
		}
		goto exit
	}

	// decimal int or float
	tkn.scanMantissa(10, buffer)

fraction:
	if tkn.lastChar == '.' {
		tkn.ConsumeNext(buffer)
		tkn.scanMantissa(10, buffer)
	}

exponent:
	if tkn.lastChar == 'e' || tkn.lastChar == 'E' {
		tkn.ConsumeNext(buffer)
		if tkn.lastChar == '+' || tkn.lastChar == '-' {
			tkn.ConsumeNext(buffer)
		}
		tkn.scanMantissa(10, buffer)
	}

exit:
	return NUMBER, buffer.Bytes()
}

func (tkn *Tokenizer) scanString(delim uint16, typ int) (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	for {
		ch := tkn.lastChar
		tkn.next()
		if ch == delim {
			if tkn.lastChar == delim {
				tkn.next()
			} else {
				break
			}
		} else if ch == '\\' {
			if tkn.lastChar == EOFCHAR {
				return LEX_ERROR, buffer.Bytes()
			}
			if decodedChar := sqltypes.SqlDecodeMap[byte(tkn.lastChar)]; decodedChar == sqltypes.DONTESCAPE {
				ch = tkn.lastChar
			} else {
				ch = uint16(decodedChar)
			}
			tkn.next()
		}
		if ch == EOFCHAR {
			return LEX_ERROR, buffer.Bytes()
		}
		buffer.WriteByte(byte(ch))
	}
	return typ, buffer.Bytes()
}

func (tkn *Tokenizer) scanCommentType1(prefix string) (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteString(prefix)
	for tkn.lastChar != EOFCHAR {
		if tkn.lastChar == '\n' {
			tkn.ConsumeNext(buffer)
			break
		}
		tkn.ConsumeNext(buffer)
	}
	return COMMENT, buffer.Bytes()
}

func (tkn *Tokenizer) scanCommentType2() (int, []byte) {
	buffer := bytes.NewBuffer(make([]byte, 0, 8))
	buffer.WriteString("/*")
	for {
		if tkn.lastChar == '*' {
			tkn.ConsumeNext(buffer)
			if tkn.lastChar == '/' {
				tkn.ConsumeNext(buffer)
				break
			}
			continue
		}
		if tkn.lastChar == EOFCHAR {
			return LEX_ERROR, buffer.Bytes()
		}
		tkn.ConsumeNext(buffer)
	}
	return COMMENT, buffer.Bytes()
}

func (tkn *Tokenizer) ConsumeNext(buffer *bytes.Buffer) {
	if tkn.lastChar == EOFCHAR {
		// This should never happen.
		panic("unexpected EOF")
	}
	buffer.WriteByte(byte(tkn.lastChar))
	tkn.next()
}

func (tkn *Tokenizer) next() {
	if ch, err := tkn.InStream.ReadByte(); err != nil {
		// Only EOF is possible.
		tkn.lastChar = EOFCHAR
	} else {
		tkn.lastChar = uint16(ch)
	}
	tkn.Position++
}

func isLetter(ch uint16) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}

func digitVal(ch uint16) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch) - '0'
	case 'a' <= ch && ch <= 'f':
		return int(ch) - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return int(ch) - 'A' + 10
	}
	return 16 // larger than any legal digit val
}

func isDigit(ch uint16) bool {
	return '0' <= ch && ch <= '9'
}

const (
	AST_AUTO_INCR = "auto_increment"
	AST_BIGINT    = "bigint"
	AST_CHAR      = "char"
	AST_CHARACTER = "character"
	AST_DATE      = "date"
	AST_ERRORS    = "errors"
	AST_EXPLAIN   = "explain"
	AST_INDEX     = "index"
	AST_INDEXES   = "indexes"
	AST_INT       = "int"
	AST_KEYS      = "keys"
	AST_FROM      = "from"
	AST_SMALLINT  = "smallint"
	AST_STATUS    = "status"
	AST_TEMPORARY = "temporary"
	AST_TIME      = "time"
	AST_TIMESTAMP = "timestamp"
	AST_TINYINT   = "tinyint"
	AST_VCHAR     = "varchar"
	AST_WARNINGS  = "warnings"
	AST_PRIMARY   = "primary"
)
