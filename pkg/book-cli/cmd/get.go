package cmd

import (
	"book-management/pkg/book-cli/pkg/options"
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get a resource",
	Short:   "retrieve a resource",
	Long:    `used to retrieve resources. Example: book-cli get <TYPE> [FLAGS|RESOURCE_IDENTIFIER]`,
	PreRunE: PreRetrieverFunction,
	Args:    cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, err := options.NewRetrieverOptions(cmd, options.Get, host, args)

		if err != nil {
			return fmt.Errorf("Error: invalid options: %v", err)
		}

		return RunCommand(opts)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().String("author", "", "book author")
	getCmd.Flags().String("title", "", "book title")
	getCmd.Flags().String("genre", "", "book genre")
	getCmd.Flags().String("dates", "", "range of published dates")
}
