package version

import (
	"fmt"
	ver "github.com/akoserwal/operator-sdk-manager/version"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Prints the version of operator-sdk-manager",
		Run: func(cmd *cobra.Command, args []string) {
			version := ver.Version
			if version == "unknown" {
				version = ver.Version
			}
			fmt.Printf("operator-sdk-manager version: %q, commit: %q, go version: %q\n",
				version, ver.GoVersion)
		},
	}

	return versionCmd
}
