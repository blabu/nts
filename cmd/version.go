package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = ""

func CreateRootCmd() *cobra.Command {
	root := &cobra.Command{}
	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  "This command print a version of application and exit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version " + version)
		},
	})

	return root
}
