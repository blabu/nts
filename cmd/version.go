package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.0.1"

func CreateRootCmd() *cobra.Command {
	root := &cobra.Command{}
	root.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  `This is version`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version " + version)
		},
	})

	return root
}
