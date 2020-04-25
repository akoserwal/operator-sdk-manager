package search

import (
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "search",
		Short: "search avaliable version",
		Long:  "",
	}

	return cmd
}

