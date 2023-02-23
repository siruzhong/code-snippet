package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
)

// CalculateMd5 Calculates the md5 hash value of the file
func CalculateMd5(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	hex := fmt.Sprintf("%x", h.Sum(nil))
	return hex, nil
}

// GetMd5CheckSum Get file md5 checksum from http response
// Recommend, cause the "Content-Md5" is a standard header
// ref: https://www.oreilly.com/library/view/http-the-definitive/1565925092/re17.html
func GetMd5CheckSum(resp *http.Response) string {
	return resp.Header.Get("Content-Md5")
}
