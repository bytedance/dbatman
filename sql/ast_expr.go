package sql

import (
	"strconv"
	"strings"
)

// IExpr represents an expression.
type IExpr interface {
	IExpr()
}

type IExprs []IExpr

// expr and boolean_primary
func (*OrExpr) IExpr()      {}
func (*XorExpr) IExpr()     {}
func (*AndExpr) IExpr()     {}
func (*NotExpr) IExpr()     {}
func (*IsCheck) IExpr()     {}
func (*NullCheck) IExpr()   {}
func (*CompareExpr) IExpr() {}
func (*Predicate) IExpr()   {}

// predicate
func (*InCond) IExpr()    {}
func (*RangeCond) IExpr() {}
func (*LikeCond) IExpr()  {}

// simple_expr
func (StrVal) IExpr()        {}
func (NumVal) IExpr()        {}
func (BoolVal) IExpr()       {}
func (HexVal) IExpr()        {}
func (BinVal) IExpr()        {}
func (ValArg) IExpr()        {}
func (*NullVal) IExpr()      {}
func (*SchemaObject) IExpr() {}
func (IValExprs) IExpr()     {}
func (*SubQuery) IExpr()     {}
func (*BinaryExpr) IExpr()   {}
func (*UnaryExpr) IExpr()    {}
func (*IntervalExpr) IExpr() {}
func (*FuncExpr) IExpr()     {}
func (*CaseExpr) IExpr()     {}
func (*CollateExpr) IExpr()  {}

// BoolExpr represents a boolean expression.
type IBoolExpr interface {
	IBoolExpr()
	IExpr
}

// expr
func (*OrExpr) IBoolExpr()    {}
func (*AndExpr) IBoolExpr()   {}
func (*XorExpr) IBoolExpr()   {}
func (*NotExpr) IBoolExpr()   {}
func (*IsCheck) IBoolExpr()   {}
func (*NullCheck) IBoolExpr() {}

// boolean_primary
func (*CompareExpr) IBoolExpr() {}
func (*Predicate) IBoolExpr()   {}

// AndExpr represents an AND expression.
type AndExpr struct {
	Left, Right IExpr
}

// OrExpr represents an OR expression.
type OrExpr struct {
	Left, Right IExpr
}

// XorExpr represents an OR expression.
type XorExpr struct {
	Left, Right IExpr
}

// NotExpr represents a NOT expression.
type NotExpr struct {
	Expr IExpr
}

// IsCheck represents an IS TRUE | FALSE | UNKNOWN expression.
type IsCheck struct {
	Operator string
	Expr     IBoolExpr
}

// IsCheck.Operator
const (
	OP_IS_TRUE        = "is true"
	OP_IS_NOT_TRUE    = "is not true"
	OP_IS_FALSE       = "is false"
	OP_IS_NOT_FALSE   = "is not false"
	OP_IS_UNKNOWN     = "is unknown"
	OP_IS_NOT_UNKNOWN = "is not unknown"
)

// NullCheck represents an IS NULL or an IS NOT NULL expression.
type NullCheck struct {
	Operator string
	Expr     IBoolExpr
}

// NullCheck.Operator
const (
	OP_IS_NULL     = "is null"
	OP_IS_NOT_NULL = "is not null"
)

// CompareExpr represents a two-value comparison expression.
type CompareExpr struct {
	Operator string
	Left     IBoolExpr
	Right    IValExpr
}

// CompareExpr.Operator
const (
	OP_EQ  = "="
	OP_LT  = "<"
	OP_GT  = ">"
	OP_LE  = "<="
	OP_GE  = ">="
	OP_NE  = "!="
	OP_NSE = "<=>"
)

type Predicate struct {
	Expr IValExpr
}

// IValExpr represents a value expression.
type IValExpr interface {
	IValExpr()
	IExpr
}

// subquery
func (*SubQuery) IValExpr() {}

// predicate
func (*InCond) IValExpr()    {}
func (*RangeCond) IValExpr() {}
func (*LikeCond) IValExpr()  {}

// bit_expr
func (*BinaryExpr) IValExpr() {}

// simple_expr
func (StrVal) IValExpr()        {} // literal
func (NumVal) IValExpr()        {}
func (BoolVal) IValExpr()       {}
func (HexVal) IValExpr()        {}
func (BinVal) IValExpr()        {}
func (*NullVal) IValExpr()      {}
func (*SchemaObject) IValExpr() {} // identifier
func (*FuncExpr) IValExpr()     {} // function
func (*CollateExpr) IValExpr()  {} // function
func (*Variable) IValExpr()     {} // variable
func (*Variable) IExpr()        {}
func (*OrOrExpr) IValExpr()     {} // || expr
func (*OrOrExpr) IExpr()        {}
func (*UnaryExpr) IValExpr()    {} // [+|-|~|!|BINARY] simple_expr
func (IExprs) IValExpr()        {} // (expr [, expr] ...)
func (IExprs) IExpr()           {}
func (IValExprs) IValExpr()     {}
func (*ExistsExpr) IValExpr()   {} // Exists (subquery)
func (*ExistsExpr) IExpr()      {}
func (ValArg) IValExpr()        {}
func (*IdentExpr) IValExpr()    {}
func (*IdentExpr) IExpr()       {}
func (*MatchExpr) IValExpr()    {}
func (*MatchExpr) IExpr()       {}
func (*CaseExpr) IValExpr()     {}
func (*IntervalExpr) IValExpr() {}

// InCond
type InCond struct {
	Operator string
	Left     IValExpr
	Right    IExprs
}

const (
	OP_IN     = "in"
	OP_NOT_IN = "not in"
)

// RangeCond represents a BETWEEN or a NOT BETWEEN expression.
type RangeCond struct {
	Operator string
	Left     IValExpr
	From, To IValExpr
}

// RangeCond.Operator
const (
	OP_BETWEEN     = "between"
	OP_NOT_BETWEEN = "not between"
)

type LikeCond struct {
	Operator string
	Left     IValExpr
	Right    IValExpr
}

const (
	OP_LIKE        = "like"
	OP_NOT_LIKE    = "not like"
	OP_SOUNDS_LIKE = "sounds like"
	OP_REGEXP      = "regexp"
	OP_NOT_REGEXP  = "not regexp"
)

// StrVal represents a string value.
type StrVal []byte

func (s StrVal) Trim() string {
	if len(s) < 1 {
		return ""
	}

	return strings.Trim(string([]byte(s)), `"'`)
}

// NumVal represents a number.
type NumVal []byte

func (n NumVal) ParseInt() (int, error) {
	if i, err := strconv.Atoi(string([]byte(n))); err != nil {
		return 0, err
	} else {
		return i, nil
	}
}

type BoolVal bool
type HexVal []byte
type BinVal []byte

// ValArg represents a named bind var argument.
type ValArg []byte

// NullVal represents a NULL value.
type NullVal struct{}

type TemporalVal struct {
	Prefix []byte
	Text   []byte
}

// IValExprs represents a list of value expressions.
// It's not a valid expression because it's not parenthesized.
type IValExprs []IValExpr

// BinaryExpr represents a binary value expression.
type BinaryExpr struct {
	Operator    string
	Left, Right IExpr
}

// BinaryExpr.Operator
const (
	OP_BITAND     = "&"
	OP_BITOR      = "|"
	OP_BITXOR     = "^"
	OP_PLUS       = "+"
	OP_MINUS      = "-"
	OP_MULT       = "*"
	OP_DIV        = "/"
	OP_MOD        = "%"
	OP_SHIFTLEFT  = "<<"
	OP_SHIFTRIGHT = ">>"
)

// UnaryExpr represents a unary value expression.
type UnaryExpr struct {
	Operator string
	Expr     IExpr
}

// UnaryExpr.Operator
const (
	OP_UPLUS   = "+"
	OP_UMINUS  = "-"
	OP_TILDA   = "~"
	OP_NOT2    = "!"
	OP_UBINARY = "binary"
)

// IntervalExpr represents a date and time function param expression
// -- http://dev.mysql.com/doc/refman/5.6/en/date-and-time-functions.html
type IntervalExpr struct {
	Expr     IExpr
	Interval []byte
}

type OrOrExpr struct {
	Left, Right IValExpr
}

// FuncExpr represents a function call.
type FuncExpr struct {
	Name     []byte
	Distinct bool
	Exprs    []IExpr
}

// CaseExpr represents a CASE expression.
type CaseExpr struct {
	Expr  IValExpr
	Whens []*When
	Else  IValExpr
}

// When represents a WHEN sub-expression.
type When struct {
	Cond IBoolExpr
	Val  IValExpr
}

// GroupBy represents a GROUP BY clause.
type GroupBy []IValExpr

// OrderBy represents an ORDER By clause.
type OrderBy []*Order

// Order represents an ordering expression.
type Order struct {
	Expr      IValExpr
	Direction string
}

// Order.Direction
const (
	OP_ASC  = "asc"
	OP_DESC = "desc"
)

// Limit represents a LIMIT clause.
type Limit struct {
	Offset, Rowcount IValExpr
}

// SchemaObject
type SchemaObject struct {
	Schema []byte
	Table  []byte
	Column []byte
}

// ExistsExpr
type ExistsExpr struct {
	SubQuery *SubQuery
}

// IdentExpr
type IdentExpr struct {
	Ident []byte
	Expr  IExpr
}

// MatchExpr
type MatchExpr struct {
}

// CollateExpr
type CollateExpr struct {
	Expr    IValExpr
	Collate []byte
}
