package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "owndatum",
	Short: "owndatum framework CLI",
}

func Execute() error {
	return RootCmd.Execute()
}
