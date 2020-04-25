package install

import (
	"fmt"
	genutil "github.com/akoserwal/operator-sdk-manager/cmd/internal"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//# Linux
//$ curl -LO https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu
//# macOS
//$ curl -LO https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk-${RELEASE_VERSION}-x86_64-apple-darwin

const (
	OPERATOR_SKD_URL  = "https://github.com/operator-framework/operator-sdk/releases/download/"
	DARWIN_URI_PREFIX = "-x86_64-apple-darwin"
	LINUX_URI_PREFIX  = "-x86_64-linux-gnu"
	OPERATOR_SDK      = "operator-sdk-"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "install version of operator-sdk",
		Long:  "",
		RunE:  installOperatorSdk,
	}

	return cmd
}

func installOperatorSdk(cmd *cobra.Command, args []string) error {
	osType := runtime.GOOS
	if len(args) > 0 {
		version := args[0]
		if version != "" {
			if osType == "darwin" {
				url := OPERATOR_SKD_URL + version + "/" + OPERATOR_SDK + version + DARWIN_URI_PREFIX
				downloadOperatorSdk(url)

			} else if osType == "linux" {
				url := OPERATOR_SKD_URL + version + "/" + OPERATOR_SDK + version + LINUX_URI_PREFIX
				downloadOperatorSdk(url)
			}
		}
	} else {
		fmt.Println("Version number should be provider like: operator-sdk-manager install V0.17.0")
	}

	return nil
}

func downloadOperatorSdk(url string) {
	tokens := strings.Split(url, "/")
	version := tokens[len(tokens)-2]
	opSdkFileName := tokens[len(tokens)-1]

	opSdkMgmrPath := getOpSdkManagerVersionPath(version)
	os.MkdirAll(opSdkMgmrPath, os.ModePerm)
	os.Chdir(opSdkMgmrPath)

	if _, err := os.Stat(opSdkFileName); os.IsNotExist(err) {
		fmt.Println("Downloading", url, "to", opSdkFileName)
		output, err := os.Create(opSdkFileName)
		if err != nil {
			fmt.Println("Error while creating", opSdkFileName, "-", err)
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

		os.Chmod(opSdkFileName, os.ModePerm)
		fmt.Println("downloaded", opSdkFileName)
	} else {
		fmt.Println("Version " + version + " is already downloaded")
	}
}

func getOpSdkManagerVersionPath(version string) string {
	home := genutil.GetHomeDir()
	opSdkMgmrPath := filepath.Join(home, ".operator-sdk-manager/versions/"+version)
	return opSdkMgmrPath
}


