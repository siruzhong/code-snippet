package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
)

// CalculateSha256 calculates the sha256 hash value of the file
func CalculateSha256(path string) (string, error) {
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

// GetSha256CheckSum get file sha256 checksum from http response
// Not recommended, cause the "X-SHA256" is not a standard header, the ones that start with X- indicate custom headers
// The general recommendation is not to use X- headers, but make sure you don't overuse standard headers
func GetSha256CheckSum(resp *http.Response) string {
	// Avoid network redirection that causes only the subrequest response body's header to have an X-SHA256 field
	for resp.Header.Get("X-SHA256") == "" && resp.Request != nil {
		resp = resp.Request.Response
	}
	return resp.Header.Get("X-SHA256")
}
