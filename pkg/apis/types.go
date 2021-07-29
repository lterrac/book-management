package apis

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	// BookType represents books
	BookType ResourceType = "book"
	// CollectionType represents collections
	CollectionType ResourceType = "collection"
	// NotSupported represents a type not currently supported
	NotSupported ResourceType = "type not supported"
)

// ResourceType is used to represent the resources managed by the application
type ResourceType string

// String returns a string representation of the resource type
func (r ResourceType) String() string {
	return string(r)
}

// Plural returns a resource plural
func (r ResourceType) Plural() string {
	return fmt.Sprintf("%vs", r)
}

// GetResource returns the corresponding resource type
func GetResource(s string) ResourceType {
	switch s {
	case BookType.String():
		return BookType
	case CollectionType.String():
		return CollectionType
	default:
		return NotSupported
	}
}

// Book represents the book object
type Book struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Isbn          string `json:"isbn"`
	PublishedDate Date   `json:"published_date"`
	Edition       uint8  `json:"edition"`
	Description   string `json:"description"`
	Genre         string `json:"genre"`
}

const dateLayout = "2006-01-02"

// Date is a custom date format to represnt date in the format of  `dateLayout` format
type Date time.Time

// UnmarshalJSON implement Unmarshaler interface
func (j *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(dateLayout, s)
	fmt.Println(t.Format(dateLayout))
	if err != nil {
		return err
	}
	*j = Date(t)
	return nil
}

// MarshalJSON implement Marshaler interface
func (j Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j).Format(dateLayout))
}

// Maybe a Format function for printing your date
func (j Date) String() string {
	t := time.Time(j)
	return t.Format(dateLayout)
}

// ValidateBookField check if a field exists in book struct
func ValidateBookField(f string) (exists bool) {
	b := Book{}
	_, exists = reflect.TypeOf(b).FieldByName(f)
	return
}

// ValidateBookValue check if a value can be assigned to the field in book struct
func ValidateBookValue(f string, v string) (ok bool) {
	b := Book{}
	field, _ := reflect.TypeOf(b).FieldByName(f)

	var value reflect.Value
	switch field.Type {
	case reflect.TypeOf(uint8(0)):
		num, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return false
		}
		value = reflect.ValueOf(uint8(num))
	case reflect.TypeOf(Date{}):
		t, err := time.Parse(dateLayout, v)
		if err != nil {
			return false
		}
		value = reflect.ValueOf(Date(t))
	default:
		value = reflect.ValueOf(v)
	}

	return value.Type().AssignableTo(field.Type)
}

// Collection represents a set of books
type Collection struct {
}
