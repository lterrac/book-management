package cmd

import (
	"book-management/pkg/apis"
	"fmt"

	"github.com/spf13/cobra"
)

var host string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "book-cli",
	Short: "book-cli is used to interact with the book management software",
	Long:  `book-cli is used to interact with the book management software`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		kind := apis.GetResource(args[0])

		if kind == apis.NotSupported {
			return fmt.Errorf("%v is not a valid type", args[0])
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVar(&host, "host", "127.0.0.1:8080", "address of book-server")
}
