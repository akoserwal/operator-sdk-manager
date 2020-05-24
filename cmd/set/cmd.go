package set

import (
	"fmt"
	genutil "github.com/akoserwal/operator-sdk-manager/cmd/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	DEFAULT_OPERATOR_SDK_PATH = "/usr/local/bin/operator-sdk"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "set as the operator-sdk version",
		Long:  "",
		RunE:  SetOperatorSdk,
	}

	cmd.AddCommand(setGoVersionCMD())

	return cmd
}

func SetOperatorSdk(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		version := strings.ToLower(args[0])
		defaultPath := filepath.Join(DEFAULT_OPERATOR_SDK_PATH)
		opSdkVersion := genutil.GetOperatorSdkFilePath(version)

		if genutil.IsOperatorAvailable(version) == true {
			if _, err := os.Lstat(defaultPath); err == nil {
				os.Remove(defaultPath)
			}

			err := os.Symlink(opSdkVersion, defaultPath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Version is not available")
		}


	} else {
		fmt.Print("Specify version to set. For eg: operator-sdk-manager set v0.17.0")
	}

	return nil
}

