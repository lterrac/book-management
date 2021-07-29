package apis

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseFilters(t *testing.T) {
	testcases := []struct {
		description string
		input       string
		desired     *FilterChain
		assert      func(desired, actual *FilterChain, err error)
	}{
		{
			description: "test equal",
			input:       "title_eq_William-Shakespeare",
			desired: newFilterChain().add(&Filter{
				Field:     "Title",
				Operation: Equals,
				Value:     "William Shakespeare",
			}),
			assert: func(desired, actual *FilterChain, err error) {
				require.Nil(t, err)
				require.Equal(t, *desired, *actual)
			},
		},
		{
			description: "test not equal",
			input:       "title_ne_William-Shakespeare",
			desired: newFilterChain().add(&Filter{
				Field:     "Title",
				Operation: NotEqual,
				Value:     "William Shakespeare",
			}),
			assert: func(desired, actual *FilterChain, err error) {
				require.Nil(t, err)
				require.Equal(t, *desired, *actual)
			},
		},
		{
			description: "test validation value uint8",
			input:       "edition_ne_2",
			desired: newFilterChain().add(&Filter{
				Field:     "Edition",
				Operation: NotEqual,
				Value:     "2",
			}),
			assert: func(desired, actual *FilterChain, err error) {
				require.Nil(t, err)
				require.Equal(t, *desired, *actual)
			},
		},
		{
			description: "test validation value Date",
			input:       "published-date_ne_2020-01-01",
			desired: newFilterChain().add(&Filter{
				Field:     "PublishedDate",
				Operation: NotEqual,
				Value:     "2020-01-01",
			}),
			assert: func(desired, actual *FilterChain, err error) {
				require.Nil(t, err)
				require.Equal(t, *desired, *actual)
			},
		},
		{
			description: "test unknown operator",
			input:       "title_invalidop_William-Shakespeare",
			desired:     newFilterChain(),
			assert: func(desired, actual *FilterChain, err error) {
				require.EqualError(t, err, fmt.Sprintf(ErrInvalidFilter, "invalidop does not exists"))
			},
		},
		{
			description: "test missing field",
			input:       "invalidfield_eq_William-Shakespeare",
			desired:     newFilterChain(),
			assert: func(desired, actual *FilterChain, err error) {
				require.EqualError(t, err, fmt.Sprintf(ErrInvalidFilter, "Invalidfield does not exists"))
			},
		},
		{
			description: "test invalid value",
			input:       "edition_eq_William-Shakespeare",
			desired:     newFilterChain(),
			assert: func(desired, actual *FilterChain, err error) {
				require.EqualError(t, err, fmt.Sprintf(ErrInvalidFilter, "Edition has a mismatching type"))
			},
		},
		{
			description: "test concat filters",
			input:       "published-date_ne_2020-01-01_and_title_eq_pippo",
			desired: newFilterChain().add(
				&Filter{
					Field:     "PublishedDate",
					Operation: NotEqual,
					Value:     "2020-01-01",
				},
			).add(
				&Filter{
					Field:     "Title",
					Operation: Equals,
					Value:     "pippo",
				},
			),
			assert: func(desired, actual *FilterChain, err error) {
				require.Nil(t, err)
				require.Equal(t, *desired, *actual)
			},
		},
		{
			description: "keep only isbn",
			input:       "published-date_ne_2020-01-01_and_title_eq_pippo_and_isbn_eq_1234",
			desired: newFilterChain().add(
				&Filter{
					Field:     "Isbn",
					Operation: Equals,
					Value:     "1234",
				},
			),
			assert: func(desired, actual *FilterChain, err error) {
				require.Nil(t, err)
				require.Equal(t, *desired, *actual)
			},
		},
		{
			description: "isbn supports only eq operator",
			input:       "published-date_ne_2020-01-01_and_title_eq_pippo_and_isbn_ne_1234",
			desired:     newFilterChain(),
			assert: func(desired, actual *FilterChain, err error) {
				require.EqualError(t, err, fmt.Sprintf(ErrInvalidFilter+" not admitted for isbn filter", NotEqual.String()))
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			actual, err := ParseFilters(tt.input, ValidateBookField, ValidateBookValue)

			tt.assert(tt.desired, actual, err)
		})
	}
}

func TestSQL(t *testing.T) {
	testcases := []struct {
		description    string
		input          *FilterChain
		desiredPrepare string
		desiredValues  []string
	}{
		{
			description: "test simple SQL",
			input: newFilterChain().add(&Filter{
				Field:     "Title",
				Operation: Equals,
				Value:     "William Shakespeare",
			}),
			desiredPrepare: "title = ?",
			desiredValues:  []string{"'William Shakespeare'"},
		},
		{
			description: "test string and int",
			input: newFilterChain().add(&Filter{
				Field:     "Title",
				Operation: Equals,
				Value:     "William Shakespeare",
			}).add(
				&Filter{
					Field:     "Edition",
					Operation: NotEqual,
					Value:     "2",
				},
			),
			desiredPrepare: "title = ? AND edition <> ?",
			desiredValues:  []string{"'William Shakespeare'", "'2'"},
		},
		{
			description: "test string, int and date",
			input: newFilterChain().add(&Filter{
				Field:     "Title",
				Operation: Equals,
				Value:     "William Shakespeare",
			}).add(
				&Filter{
					Field:     "Edition",
					Operation: NotEqual,
					Value:     "2",
				},
			).add(
				&Filter{
					Field:     "PublishedDate",
					Operation: NotEqual,
					Value:     "2020-01-01",
				},
			),
			desiredPrepare: "title = ? AND edition <> ? AND published_date <> ?",
			desiredValues:  []string{"'William Shakespeare'", "'2'", "'2020-01-01'"},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.description, func(t *testing.T) {
			actualPrepare, actualQueryValues := tt.input.SQLStatement()
			require.Equal(t, tt.desiredPrepare, actualPrepare)
			require.Equal(t, tt.desiredValues, actualQueryValues)
		})
	}
}
