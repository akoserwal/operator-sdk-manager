package internal

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func GetGoPath() string {
	home := GetHomeDir()
	return filepath.Join(home, ".osm/go-ver/")
}


func GetOpSdkManagerVersionPath(version string) string {
	home := GetHomeDir()
	opSdkMgmrPath := filepath.Join(home, ".osm/versions/", version)
	return opSdkMgmrPath
}

func GetGoVersionPath(version string) string {
	home := GetHomeDir()
	goDownloadPath := filepath.Join(home, ".osm/go-ver/", version)
	return goDownloadPath
}

func IsOperatorAvailable(version string) bool {
	opSdkVersion := GetOperatorSdkFilePath(version)
	if _, err := os.Stat(opSdkVersion); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func IsGoVerAvailable(version string) bool {
	goVersion := GetGoVersionPath(version)
	if _, err := os.Stat(goVersion); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func GetOperatorSdkFilePath(version string) string {
	opSdkVersionPath := GetOpSdkManagerVersionPath(version)
	opSdkVersion := filepath.Join(opSdkVersionPath, "operator-sdk")
	return opSdkVersion
}

func CreateDir(name string) string {
	home := GetHomeDir()
	opSdkMgmrOperator := filepath.Join(home, ".osm/"+name)
	if _, err := os.Stat(opSdkMgmrOperator); os.IsNotExist(err) {
		os.MkdirAll(opSdkMgmrOperator, os.ModePerm)
	}
	return opSdkMgmrOperator
}


// untarrer extract contant of file tarName into location xpath
func Untar(tarName, xpath string) (err error) {
	tarFile, err := os.Open(tarName)
	if err != nil {
		return err
	}
	defer func() {
		err = tarFile.Close()
	}()

	absPath, err := filepath.Abs(xpath)
	if err != nil {
		return err
	}

	tarReader, err := getTarReader(tarFile, tarName)
	if err != nil {
		return err
	}

	for {
		tarHeader, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fileName := tarHeader.Name
		fileNamePath, err := checkFilePath(fileName, err)
		if err != nil {
			return err
		}
		extFilePath := filepath.Join(absPath, fileNamePath)

		fileInfo := tarHeader.FileInfo()
		if fileInfo.Mode().IsDir() {
			if err := os.MkdirAll(extFilePath, 0755); err != nil {
				return err
			}
			continue
		}

		err, done := createExtractedFile(err, extFilePath, fileInfo, tarReader)
		if done {
			return err
		}
	}
	return nil
}

func createExtractedFile(err error, goFileName string, fileInfo os.FileInfo, tarReader *tar.Reader) (error, bool) {
	// create new file with original file mode
	file, err := os.OpenFile(goFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, fileInfo.Mode().Perm())
	if err != nil {
		return err, true
	}
	written, copyErr := io.Copy(file, tarReader)
	if closeErr := file.Close(); closeErr != nil { // close file immediately
		return err, true
	}
	if copyErr != nil {
		return copyErr, true
	}
	if written != fileInfo.Size() {
		return fmt.Errorf("unexpected bytes written: wrote %d, want %d", written, fileInfo.Size()), true
	}
	return nil, false
}

func checkFilePath(fileName string, err error) (string, error) {
	if filepath.IsAbs(fileName) {
		fmt.Printf("removing / prefix from %s\n", fileName)
		fileName, err = filepath.Rel("/", fileName)
		if err != nil {
			return "", err
		}
	}
	return fileName, err
}

func getTarReader(tarFile *os.File, tarName string) (*tar.Reader, error) {
	tr := tar.NewReader(tarFile)
	if strings.HasSuffix(tarName, ".gz") || strings.HasSuffix(tarName, ".gzip") {
		gz, err := gzip.NewReader(tarFile)
		if err != nil {
			return nil, err
		}
		defer gz.Close()
		tr = tar.NewReader(gz)
	}
	return tr, nil
}