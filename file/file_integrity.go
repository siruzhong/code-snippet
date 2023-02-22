package file

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/imroc/req/v3"
)

// client a simple go http client
var client = req.C()

// DownloadFile Download the file and verify its integrity
func DownloadFile(originUrl string) error {
	filePath := os.ExpandEnv("$PWD.tar.gz")
	// download file from origin url
	resp, err := client.R().SetOutputFile(filePath).Get(originUrl)
	if err != nil {
		return err
	}
	// verify the download file integrity
	checksum := GetCheckSum(resp.Response)
	if checksum != "" {
		fileChecksum, err := Sha256(filePath)
		if err != nil {
			return err
		}
		if !strings.EqualFold(checksum, fileChecksum) {
			return errors.New("the download file is incomplete")
		}
	}
	return nil
}

// GetCheckSum get file checksum from http response(prevent http request nesting)
func GetCheckSum(resp *http.Response) string {
	for resp.Header.Get("X-SHA256") == "" && resp.Request != nil {
		resp = resp.Request.Response
	}
	return resp.Header.Get("X-SHA256")
}

// Sha256 Obtain the SHA256 verification value of the file
func Sha256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	hex := fmt.Sprintf("%x", h.Sum(nil))
	return hex, nil
}
