package apis

import (
	"container/list"

	"github.com/iancoleman/strcase"
)

// Operator is the type of the filter operation
type Operator string

const (
	And      Operator = "and"
	Equals   Operator = "eq"
	NotEqual Operator = "ne"

	ErrInvalidFilter string = "invalid filter: %v"
)

func (o Operator) String() string {
	return string(o)
}

// Symbol returns the SQL symbol for the operator
func (o Operator) Symbol() (symbol string) {
	switch o {
	case And:
		symbol = "&"
	case Equals:
		symbol = "="
	case NotEqual:
		symbol = "<>"
	}
	return
}

// Filter is a filter for a field and an operation
type Filter struct {
	Field     string
	Operation Operator
	Value     string
}

func (f *Filter) FieldName() string {
	return f.Field
}

func (f *Filter) SQL() (string, []interface{}) {
	return strcase.ToSnake(f.Field) + " " + f.Operation.Symbol() + " ?", []interface{}{f.Value}
}

// FilterChain is a chain of filters
type FilterChain struct {
	chain *list.List
}

func newFilterChain() *FilterChain {
	return &FilterChain{
		chain: list.New(),
	}
}

func (f *FilterChain) add(filter Filter) *FilterChain {
	f.chain.PushBack(filter)
	return f
}
