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

const DefaultGoPath = "/usr/local/bin/go"

func setGoVersionCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go",
		Short: "set go version",
		Long:  "operator-sdk set go x.x.x",
		RunE:  SetGoVerion,
	}

	return cmd
}

func SetGoVerion(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		version := strings.ToLower(args[0])
		defaultPath := filepath.Join(DefaultGoPath)
		goVersionPath := genutil.GetGoVersionPath(version)
		goPath := filepath.Join(goVersionPath, "go/bin", "go")
		fmt.Println(goPath)

		if genutil.IsGoVerAvailable(version) == true {
			if _, err := os.Lstat(defaultPath); err == nil {
				os.Remove(defaultPath)
			}

			err := os.Symlink(goPath, defaultPath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Version is not available")
		}


	} else {
		fmt.Print("Specify version to set. For eg: operator-sdk-manager set go 0.17.0")
	}

	return nil
}
