package options

import (
	"book-management/pkg/apis"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// ResourceOperation is an abstraction of the HTTP verb that will be used to interact with the database
type ResourceOperation string

const (
	Create ResourceOperation = http.MethodPost
	Get    ResourceOperation = http.MethodGet
	Update ResourceOperation = http.MethodPut
	Delete ResourceOperation = http.MethodDelete
)

// String returns a string representation of the resource type
func (r ResourceOperation) String() string {
	return string(r)
}

// CommandOptions are the settings that define a resource operation
type CommandOptions struct {
	Server    string
	Resource  apis.ResourceType
	Operation ResourceOperation
	Object    string
	Filters   []string
}

// NewModifierOptions forms the options for a modifier command
func NewModifierOptions(cmd *cobra.Command, op ResourceOperation, host string, args []string) (*CommandOptions, error) {
	kind := apis.GetResource(args[0])

	var obj string
	fileFlag := cmd.Flag("file").Value.String()

	if fileFlag != "" {
		bytes, err := os.ReadFile(fileFlag)

		if err != nil {
			return nil, fmt.Errorf("reading file path: %v", err)
		}
		obj = string(bytes)
	} else {
		obj = args[1]
	}

	return newCommandOptions(kind, op, obj, host, []string{}), nil
}

// NewModifierOptions forms the options for a retriever command
func NewRetrieverOptions(cmd *cobra.Command, op ResourceOperation, host string, args []string) (*CommandOptions, error) {
	kind := apis.GetResource(args[0])

	var id string

	// use only resource identifier if provided
	if len(args) > 1 {
		id = strings.ReplaceAll(strings.ToLower(args[1]), " ", "-")
		return newCommandOptions(kind, op, "", host, []string{"isbn_eq_" + id}), nil
	}

	filters := []string{}
	authorFlag := cmd.Flag("author").Value.String()
	titleFlag := cmd.Flag("title").Value.String()
	genreFlag := cmd.Flag("genre").Value.String()
	datesFlag := cmd.Flag("dates").Value.String()

	if authorFlag != "" {
		filters = append(filters, "author_eq_"+strings.ReplaceAll(strings.ToLower(authorFlag), " ", "-"))
	}

	if titleFlag != "" {
		filters = append(filters, "title_eq_"+strings.ReplaceAll(strings.ToLower(titleFlag), " ", "-"))
	}

	if genreFlag != "" {
		filters = append(filters, "genre_eq_"+strings.ReplaceAll(strings.ToLower(genreFlag), " ", "-"))
	}

	if datesFlag != "" {
		filters = append(filters, "dates_eq_"+strings.ReplaceAll(strings.ToLower(datesFlag), " ", "-"))
	}

	return newCommandOptions(kind, op, "", host, filters), nil
}

func newCommandOptions(kind apis.ResourceType, op ResourceOperation, obj string, server string, filters []string) *CommandOptions {
	return &CommandOptions{
		Server:    server,
		Resource:  kind,
		Operation: op,
		Object:    obj,
		Filters:   filters,
	}
}

// URL forms the correct URL for a command
func (opts *CommandOptions) URL() string {
	baseURL := fmt.Sprintf("http://%s/api/v1/%s", opts.Server, opts.Resource.Plural())
	if len(opts.Filters) == 0 {
		return baseURL
	}

	baseFilterURL := fmt.Sprintf("%s?filter=", baseURL)

	return baseFilterURL + strings.Join(opts.Filters, "_and_")
}
