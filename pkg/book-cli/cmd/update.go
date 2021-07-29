package cmd

import (
	"book-management/pkg/book-cli/pkg/options"
	"book-management/pkg/book-cli/pkg/validation"
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update resource",
	Short:   "update a new resource",
	Long:    `used to update new resources. Example: book-cli update <TYPE> [OPTIONS] [ -f FILE-PATH | OBJECT]`,
	PreRunE: PreModifierFunction,
	Args:    cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		opts, err := options.NewModifierOptions(cmd, options.Update, host, args)

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
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("file", "f", "", "path to JSON resource file")
}
