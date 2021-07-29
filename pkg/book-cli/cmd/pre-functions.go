package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// PreModifierFunction checks whether the resource definition is passed as arg or a valid path to a file is supplied with -f flag
func PreModifierFunction(cmd *cobra.Command, args []string) error {
	if len(args) == 1 && cmd.Flag("file").Value.String() == "" {
		return fmt.Errorf("provide a resource definition on the command line or using -f flag")
	}
	return nil
}
