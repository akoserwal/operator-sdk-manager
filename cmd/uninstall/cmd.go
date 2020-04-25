package uninstall

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall version of operator-sdk",
		Long:  "",
	}

	return cmd
}
