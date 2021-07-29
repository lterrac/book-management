package apis

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

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

// Collection represents a set of books
type Collection struct {
}
