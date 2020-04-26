package list

import (
	"fmt"
	genutil "github.com/akoserwal/operator-sdk-manager/cmd/internal"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all installed versions",
		Long:  "",
		RunE: listInstalledOperators,
	}

	return cmd
}

func listInstalledOperators(cmd *cobra.Command, args []string) error {
	opSdkMgmrPath := genutil.GetOpSdkManagerPath()
	files, err := ioutil.ReadDir(opSdkMgmrPath)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		fmt.Println("No version is installed")
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
	return nil
}

