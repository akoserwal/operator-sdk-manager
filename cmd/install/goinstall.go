package install

import (
	"fmt"
	genutil "github.com/akoserwal/operator-sdk-manager/cmd/internal"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const (
	GO_DOWNLOAD_URL      = "https://dl.google.com/go/"
	DARWIN_GO_URI_PREFIX = "darwin-amd64"
	LINUX_64_URI_PREFIX  = "linux-amd64"
	LINUX_386_URI_PREFIX = "linux-386"
	GO_STR               = "go"
)

//tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz

func InstallGOCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go",
		Short: "go x.x.x",
		Long:  "operator-sdk-manager install -go v1.13.1",
		RunE:  installGoVersion,
	}

	return cmd
}

func installGoVersion(cmd *cobra.Command, args []string) error {
	osType := runtime.GOOS
	osArch := runtime.GOARCH
	if len(args) > 0 {
		version := strings.ToLower(args[0])
		if version != "" {
			url := GO_DOWNLOAD_URL + GO_STR + version + "." + osType + "-" + osArch + ".tar.gz"
			fmt.Println(url)
			goPathDir := genutil.CreateDir("go-ver")
			os.Chdir(goPathDir)
			downloadGoVersion(url, version)
		}
	}
	return nil
}

func downloadGoVersion(url string, ver string) {
	version := ver
	goFileName := "go" + ver + ".tar.gz"

	goVerPath := genutil.GetGoVersionPath(version)
	os.MkdirAll(goVerPath, os.ModePerm)
	os.Chdir(goVerPath)

	if _, err := os.Stat(goFileName); os.IsNotExist(err) {
		fmt.Println("Downloading", url, "to", goFileName)
		output, err := os.Create(goFileName)
		if err != nil {
			fmt.Println("Error while creating", goFileName, "-", err)
			return
		}
		defer output.Close()

		response, err := http.Get(url)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}
		defer response.Body.Close()

		_, err = io.Copy(output, response.Body)
		if err != nil {
			fmt.Println("Error while downloading", url, "-", err)
			return
		}

		os.Chmod(goFileName, os.ModePerm)
		fmt.Println("downloaded", goFileName)
		genutil.Untar(goFileName, goVerPath)
		os.RemoveAll(goVerPath + "/" + goFileName)
	} else {
		fmt.Println("Version " + version + " is already downloaded")
	}
}
