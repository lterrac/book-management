package validation

import (
	"book-management/pkg/apis"
	"encoding/json"
	"fmt"
)

// ValidateResource ensures that input string can be unmarshaled in the correct data structure
func ValidateResource(kind apis.ResourceType, obj string) error {
	switch kind {
	case apis.BookType:
		var book apis.Book
		return json.Unmarshal([]byte(obj), &book)
	case apis.CollectionType:
		var collection apis.Collection
		return json.Unmarshal([]byte(obj), &collection)
	default:
		return fmt.Errorf("unsupported type")
	}
}
