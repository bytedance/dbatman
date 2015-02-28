package sql

// Expr represents an expression.
type Expr interface {
	IExpr()
}

func (*AndExpr) IExpr()        {}
func (*OrExpr) IExpr()         {}
func (*XorExpr) IExpr()        {}
func (*NotExpr) IExpr()        {}
func (*ParenBoolExpr) IExpr()  {}
func (*ComparisonExpr) IExpr() {}
func (*RangeCond) IExpr()      {}
func (*NullCheck) IExpr()      {}
func (*IsCheck) IExpr()        {}
func (*ExistsExpr) IExpr()     {}
func (StrVal) IExpr()          {}
func (NumVal) IExpr()          {}
func (ValArg) IExpr()          {}
func (*NullVal) IExpr()        {}
func (*ColName) IExpr()        {}
func (ValExprs) IExpr()        {}
func (*SubQuery) IExpr()       {}
func (*BinaryExpr) IExpr()     {}
func (*UnaryExpr) IExpr()      {}
func (*FuncExpr) IExpr()       {}
func (*CaseExpr) IExpr()       {}

// BoolExpr represents a boolean expression.
type BoolExpr interface {
	IBoolExpr()
	Expr
}

func (*AndExpr) IBoolExpr()        {}
func (*OrExpr) IBoolExpr()         {}
func (*XorExpr) IBoolExpr()        {}
func (*NotExpr) IBoolExpr()        {}
func (*ParenBoolExpr) IBoolExpr()  {}
func (*ComparisonExpr) IBoolExpr() {}
func (*RangeCond) IBoolExpr()      {}
func (*NullCheck) IBoolExpr()      {}
func (*IsCheck) IBoolExpr()        {}
func (*ExistsExpr) IBoolExpr()     {}

// AndExpr represents an AND expression.
type AndExpr struct {
	Left, Right BoolExpr
}

// OrExpr represents an OR expression.
type OrExpr struct {
	Left, Right BoolExpr
}

// XorExpr represents an OR expression.
type XorExpr struct {
	Left, Right BoolExpr
}

// NotExpr represents a NOT expression.
type NotExpr struct {
	Expr BoolExpr
}

// ParenBoolExpr represents a parenthesized boolean expression.
type ParenBoolExpr struct {
	Expr BoolExpr
}

// ComparisonExpr represents a two-value comparison expression.
type ComparisonExpr struct {
	Operator    string
	Left, Right ValExpr
}

// ComparisonExpr.Operator
const (
	OP_EQ          = "="
	OP_LT          = "<"
	OP_GT          = ">"
	OP_LE          = "<="
	OP_GE          = ">="
	OP_NE          = "!="
	OP_NSE         = "<=>"
	OP_IN          = "in"
	OP_NOT_IN      = "not in"
	OP_LIKE        = "like"
	OP_NOT_LIKE    = "not like"
	OP_SOUNDS_LIKE = "sounds like"
	OP_REGEXP      = "regexp"
	OP_NOT_REGEXP  = "not regexp"
)

// RangeCond represents a BETWEEN or a NOT BETWEEN expression.
type RangeCond struct {
	Operator string
	Left     ValExpr
	From, To ValExpr
}

// RangeCond.Operator
const (
	OP_BETWEEN     = "between"
	OP_NOT_BETWEEN = "not between"
)

// IsCheck represents an IS TRUE | FALSE | UNKNOWN expression.
type IsCheck struct {
	Operator string
	Expr     ValExpr
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
	Expr     ValExpr
}

// NullCheck.Operator
const (
	OP_IS_NULL     = "is null"
	OP_IS_NOT_NULL = "is not null"
)

// ExistsExpr represents an EXISTS expression.
type ExistsExpr struct {
	SubQuery *SubQuery
}

// ValExpr represents a value expression.
type ValExpr interface {
	IValExpr()
	Expr
}

func (StrVal) IValExpr()      {}
func (NumVal) IValExpr()      {}
func (ValArg) IValExpr()      {}
func (*NullVal) IValExpr()    {}
func (*ColName) IValExpr()    {}
func (ValExprs) IValExpr()    {}
func (*SubQuery) IValExpr()   {}
func (*BinaryExpr) IValExpr() {}
func (*UnaryExpr) IValExpr()  {}
func (*FuncExpr) IValExpr()   {}
func (*CaseExpr) IValExpr()   {}

// StrVal represents a string value.
type StrVal []byte

// NumVal represents a number.
type NumVal []byte

// ValArg represents a named bind var argument.
type ValArg []byte

// NullVal represents a NULL value.
type NullVal struct{}

// ColName represents a column name.
type ColName struct {
	Name, Qualifier []byte
}

// ValExprs represents a list of value expressions.
// It's not a valid expression because it's not parenthesized.
type ValExprs []ValExpr

// BinaryExpr represents a binary value expression.
type BinaryExpr struct {
	Operator    byte
	Left, Right Expr
}

// BinaryExpr.Operator
const (
	OP_BITAND = '&'
	OP_BITOR  = '|'
	OP_BITXOR = '^'
	OP_PLUS   = '+'
	OP_MINUS  = '-'
	OP_MULT   = '*'
	OP_DIV    = '/'
	OP_MOD    = '%'
)

// UnaryExpr represents a unary value expression.
type UnaryExpr struct {
	Operator byte
	Expr     Expr
}

// UnaryExpr.Operator
const (
	OP_UPLUS  = '+'
	OP_UMINUS = '-'
	OP_TILDA  = '~'
)

// FuncExpr represents a function call.
type FuncExpr struct {
	Name     []byte
	Distinct bool
	Exprs    []Expr
}

// CaseExpr represents a CASE expression.
type CaseExpr struct {
	Expr  ValExpr
	Whens []*When
	Else  ValExpr
}

// When represents a WHEN sub-expression.
type When struct {
	Cond BoolExpr
	Val  ValExpr
}

// Values represents a VALUES clause.
type Values []Tuple

// GroupBy represents a GROUP BY clause.
type GroupBy []ValExpr

// OrderBy represents an ORDER By clause.
type OrderBy []*Order

// Order represents an ordering expression.
type Order struct {
	Expr      ValExpr
	Direction string
}

// Order.Direction
const (
	OP_ASC  = "asc"
	OP_DESC = "desc"
)

// Limit represents a LIMIT clause.
type Limit struct {
	Offset, Rowcount ValExpr
}
