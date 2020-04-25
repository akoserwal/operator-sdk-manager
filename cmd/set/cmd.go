package set

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "set as the operator-sdk version",
		Long:  "",
	}

	return cmd
}
