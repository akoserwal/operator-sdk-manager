package uninstall

import (
	"fmt"
	genutil "github.com/akoserwal/operator-sdk-manager/cmd/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall version of operator-sdk",
		Long:  "",
		RunE: removeOperatorVersion,
	}

	return cmd
}

func removeOperatorVersion(cmd *cobra.Command, args []string)  error {
	if len(args) > 0 {
		version := args[0]
		opFilePath :=genutil.GetOpSdkManagerVersionPath(version)
		if _, err := os.Stat(opFilePath); err != nil {
			if os.IsNotExist(err) {
				log.Fatal(err)
				os.Exit(1)
			}
		} else {
			os.RemoveAll(opFilePath)
		}

	} else {
		fmt.Println("Specific a version to uninstall")
	}

	return nil
}