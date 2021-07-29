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

// PreRetrieverFunction checks whether the resource identifier is passed as arg or at least one of the filtering args is supplied as flag
func PreRetrieverFunction(cmd *cobra.Command, args []string) error {
	if len(args) == 1 &&
		cmd.Flag("author").Value.String() == "" &&
		cmd.Flag("title").Value.String() == "" &&
		cmd.Flag("dates").Value.String() == "" &&
		cmd.Flag("genre").Value.String() == "" {
		return fmt.Errorf("provide resource identifier or at least one valid filter")
	}
	return nil
}
