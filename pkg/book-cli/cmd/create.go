package cmd

import (
	"book-management/pkg/book-cli/cmd/pkg/options"
	"book-management/pkg/book-cli/cmd/pkg/validation"
	"fmt"

	"github.com/spf13/cobra"
)

// createCmd creates a new resource
var createCmd = &cobra.Command{
	Use:     "create resource",
	Short:   "create a new resource",
	Long:    `used to create new resources. Example: book-cli create <TYPE> [OPTIONS] [ -f FILE-PATH | OBJECT]`,
	PreRunE: PreModifierFunction,
	Args:    cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, err := options.NewModifierOptions(cmd, options.Create, host, args)

		if err != nil {
			return fmt.Errorf("Error: invalid options: %v", err)
		}

		err = validation.ValidateResource(opts.Resource, opts.Object)

		if err != nil {
			return fmt.Errorf("Error: invalid resource definition: %v", err)
		}

		return RunCommand(opts)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("file", "f", "", "path to JSON resource file")
}
