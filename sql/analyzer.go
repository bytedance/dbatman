// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

// analyzer.go contains utility analysis functions.

import (
	"fmt"

	"micode.be.xiaomi.com/wangjing1/go-mysql-proxy/sqltypes"
)

// GetDBName parses the specified DML and returns the
// db name if it was used to qualify the table name.
// It returns an error if parsing fails or if the statement
// is not a DML.
func GetDBName(sql string) (string, error) {
	statement, err := Parse(sql)
	if err != nil {
		return "", err
	}

	return GetStmtDB(statement)
}

func GetStmtDB(statement Statement) (string, error) {
	switch stmt := statement.(type) {
	case *Insert:
		return string(stmt.Table.Qualifier), nil
	case *Update:
		return string(stmt.Table.Qualifier), nil
	case *Delete:
		return string(stmt.Table.Qualifier), nil
	case *Replace:
		return string(stmt.Table.Qualifier), nil
	case *Select:
		if len(stmt.From) != 0 {
			return GetTableExprDB(stmt.From[0])
		} else {
			return "", nil
		}
	case *Create, *Set, *Drop:
		return "", nil
	case *CreateTable:
		return string(stmt.Table.Qualifier), nil
	case *Show:
		if stmt.Table != nil {
			return string(stmt.Table.Qualifier), nil
		} else {
			return "", nil
		}
	case *Explain:
		if stmt.Table != nil {
			return string(stmt.Table.Qualifier), nil
		} else {
			// explain select
			return "", nil
		}
	}

	return "", fmt.Errorf("statement '%v' is unsupported for parse dbname", statement)
}

func GetTableExprDB(node TableExpr) (string, error) {
	switch s := node.(type) {
	case SimpleTableExpr:
		return GetSimpleTableExprDB(s), nil
	case *AliasedTableExpr:
		return GetSimpleTableExprDB(s.Expr), nil
	default:
		//return "", fmt.Errorf("unsupport get db from %v", s)
		return "", nil
	}
}

func GetSimpleTableExprDB(node SimpleTableExpr) string {
	if n, ok := node.(*TableName); ok && n.Qualifier != nil {
		return string(n.Qualifier)
	}

	return ""
}

// GetTableName returns the table name from the SimpleTableExpr
// only if it's a simple expression. Otherwise, it returns "".
func GetTableName(node SimpleTableExpr) string {
	if n, ok := node.(*TableName); ok && n.Qualifier == nil {
		return string(n.Name)
	}
	// sub-select or '.' expression
	return ""
}

// GetColName returns the column name, only if
// it's a simple expression. Otherwise, it returns "".
func GetColName(node Expr) string {
	if n, ok := node.(*ColName); ok {
		return string(n.Name)
	}
	return ""
}

// IsColName returns true if the ValExpr is a *ColName.
func IsColName(node ValExpr) bool {
	_, ok := node.(*ColName)
	return ok
}

// IsValue returns true if the ValExpr is a string, number or value arg.
// NULL is not considered to be a value.
func IsValue(node ValExpr) bool {
	switch node.(type) {
	case StrVal, NumVal, ValArg:
		return true
	}
	return false
}

// HasINCaluse returns true if any of the conditions has an IN clause.
func HasINClause(conditions []BoolExpr) bool {
	for _, node := range conditions {
		if c, ok := node.(*ComparisonExpr); ok && c.Operator == AST_IN {
			return true
		}
	}
	return false
}

// IsSimpleTuple returns true if the ValExpr is a ValTuple that
// contains simple values or if it's a list arg.
func IsSimpleTuple(node ValExpr) bool {
	switch vals := node.(type) {
	case ValTuple:
		for _, n := range vals {
			if !IsValue(n) {
				return false
			}
		}
		return true
	case ListArg:
		return true
	}
	// It's a subquery
	return false
}

// AsInterface converts the ValExpr to an interface. It converts
// ValTuple to []interface{}, ValArg to string, StrVal to sqltypes.String,
// NumVal to sqltypes.Numeric, NullVal to nil.
// Otherwise, it returns an error.
func AsInterface(node ValExpr) (interface{}, error) {
	switch node := node.(type) {
	case ValTuple:
		vals := make([]interface{}, 0, len(node))
		for _, val := range node {
			v, err := AsInterface(val)
			if err != nil {
				return nil, err
			}
			vals = append(vals, v)
		}
		return vals, nil
	case ValArg:
		return string(node), nil
	case ListArg:
		return string(node), nil
	case StrVal:
		return sqltypes.MakeString(node), nil
	case NumVal:
		n, err := sqltypes.BuildNumeric(string(node))
		if err != nil {
			return nil, fmt.Errorf("type mismatch: %s", err)
		}
		return n, nil
	case *NullVal:
		return nil, nil
	}
	return nil, fmt.Errorf("unexpected node %v", node)
}

// StringIn is a convenience function that returns
// true if str matches any of the values.
func StringIn(str string, values ...string) bool {
	for _, val := range values {
		if str == val {
			return true
		}
	}
	return false
}
