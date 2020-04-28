package internal

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home
}

func GetOpSdkManagerPath() string {
	home := GetHomeDir()
	opSdkMgmrPath := filepath.Join(home, ".osm/versions/")
	return opSdkMgmrPath
}

func GetOpSdkManagerVersionPath(version string) string {
	home := GetHomeDir()
	opSdkMgmrPath := filepath.Join(home, ".osm/versions/", version)
	return opSdkMgmrPath
}

func IsOperatorAvailable(opSdkVersion string) bool {
	if _, err := os.Stat(opSdkVersion); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
