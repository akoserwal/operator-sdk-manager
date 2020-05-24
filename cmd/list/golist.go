package list

import (
	"fmt"
	genutil "github.com/akoserwal/operator-sdk-manager/cmd/internal"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

func GetGoVersionListCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go",
		Short: "operator-sdk-manager list go",
		Long:  "operator-sdk-manager list go",
		RunE:  getGoList,
	}

	return cmd
}

func getGoList(cmd *cobra.Command, args []string) error {
	goInstallationPath := genutil.GetGoPath()
	files, err := ioutil.ReadDir(goInstallationPath)
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

