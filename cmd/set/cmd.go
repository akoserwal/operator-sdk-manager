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

	return cmd
}

func SetOperatorSdk(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		version := strings.ToLower(args[0])
		defaultPath := filepath.Join(DEFAULT_OPERATOR_SDK_PATH)
		opSdkVersionPath := genutil.GetOpSdkManagerVersionPath(version)
		opSdkVersion := filepath.Join(opSdkVersionPath, "operator-sdk")

		isOperatorAvailable(opSdkVersion)

		if _, err := os.Lstat(defaultPath); err == nil {
			os.Remove(defaultPath)
		}

		err := os.Symlink(opSdkVersion, defaultPath)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print("Specific version to set like: operator-sdk-manager set v0.17.0")
	}

	return nil
}

func isOperatorAvailable(opSdkVersion string) bool {
	if _, err := os.Stat(opSdkVersion); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
