package apis

import (
	"container/list"
	"fmt"
	"regexp"
	"strings"

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

type fieldValidatorFunc func(field string) bool
type valueValidatorFunc func(field string, value string) bool

// SQLConverter converts a filter into a SQL statement.
type SQLConverter interface {
	// SQL returns the SQL query and the parameters to bind to it.
	SQL() (prepare string, query []interface{})
	// FieldName returns the resource field name
	FieldName() string
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

// DateRangeFilter is a filter for a date range
type DateRangeFilter struct {
	Field     string
	StartDate string
	EndDate   string
}

func (d *DateRangeFilter) SQL() (string, []interface{}) {
	return strcase.ToSnake(d.Field) + " >= ? AND " + d.Field + " <= ?", []interface{}{d.StartDate, d.EndDate}
}

func (d *DateRangeFilter) FieldName() string {
	return d.Field
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

func (f *FilterChain) Get(field string) (SQLConverter, bool) {
	for e := f.chain.Front(); e != nil; e = e.Next() {
		if e.Value.(SQLConverter).FieldName() == field {
			return e.Value.(SQLConverter), true
		}
	}
	return nil, false
}

func (f *FilterChain) add(filter SQLConverter) *FilterChain {
	f.chain.PushBack(filter)
	return f
}

// SQLStatement converts all filters in their corresponding SQL statement.
func (f *FilterChain) SQLStatement() (prepare string, query []interface{}) {
	var prepares []string

	for e := f.chain.Front(); e != nil; e = e.Next() {
		prepare, values := e.Value.(SQLConverter).SQL()
		prepares = append(prepares, prepare)
		query = append(query, values...)
	}

	return strings.Join(prepares, " AND "), query
}

// ParseFilters build a filterchain from a string. If anything goes wrong, it returns an ErrInvalidFilter. It requires a fieldvalidator and a valuevalidator to perform type checking on the struct field.
func ParseFilters(filters string, validateField fieldValidatorFunc, validateValue valueValidatorFunc) (chain *FilterChain, err error) {
	filterChain, err := parseFilters(filters, validateField, validateValue)

	if err != nil {
		return nil, err
	}

	if filterChain.chain.Len() == 0 {
		return nil, fmt.Errorf("no filters")
	}

	var isbnFilter *Filter
	converter, exists := filterChain.Get("Isbn")

	// Use only isbn to retrieve the book if specified
	if exists {
		isbnFilter = converter.(*Filter)
		singleChain := newFilterChain()

		if isbnFilter.Operation != Equals {
			return nil, fmt.Errorf(ErrInvalidFilter+" not admitted for isbn filter", isbnFilter.Operation)
		}
		singleChain.add(isbnFilter)
		return singleChain, nil
	}

	return filterChain, nil
}

// ParseFilters build a filterchain from a string. If anything goes wrong, it returns an ErrInvalidFilter. It requires a fieldvalidator and a valuevalidator to perform type checking on the struct field.
func parseFilters(filters string, validateField fieldValidatorFunc, validateValue valueValidatorFunc) (chain *FilterChain, err error) {
	chain = newFilterChain()

	// Split all the single filters
	filtersRaw := strings.Split(filters, "_"+And.String()+"_")

	for _, filter := range filtersRaw {
		// Split field, operator and value
		parts := strings.Split(filter, "_")

		if len(parts)%3 != 0 {
			return nil, fmt.Errorf(ErrInvalidFilter, "wrong number of filter parts")
		}

		// Transform field string to UpperCamelCase and remove spaces
		field := strcase.ToCamel(parts[0])
		operator := parts[1]
		value := parts[2]

		var converter SQLConverter

		// date filter should be handle as a special case
		if field == "Dates" {
			converter, err = parseDateRange(value)
			if err != nil {
				return nil, fmt.Errorf(ErrInvalidFilter, err)
			}
			chain.add(converter)
			continue
		}

		if !validateField(field) {
			return nil, fmt.Errorf(ErrInvalidFilter+" does not exists", field)
		}

		f := &Filter{}
		f.Field = field

		switch operator {
		case Equals.String():
			f.Operation = Equals
			break
		case NotEqual.String():
			f.Operation = NotEqual
			break
		default:
			return nil, fmt.Errorf(ErrInvalidFilter+" does not exists", operator)
		}

		// replace all '-' if the value is not a date
		if isDate := regexp.MustCompile(`^[0-9|-]+$`).MatchString; !isDate(value) {
			value = strings.ReplaceAll(value, "-", " ")
		}

		if !validateValue(field, value) {
			return nil, fmt.Errorf(ErrInvalidFilter+" has a mismatching type", field)
		}

		f.Value = value
		chain.add(f)
	}

	return chain, nil
}

func parseDateRange(dateRange string) (*DateRangeFilter, error) {
	parts := strings.Split(dateRange, "-to-")

	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid date range")
	}

	startDate := parts[0]
	endDate := parts[1]

	return &DateRangeFilter{
		Field:     "published_date",
		StartDate: startDate,
		EndDate:   endDate,
	}, nil
}
