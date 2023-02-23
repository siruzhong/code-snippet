package file

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"

	"code-snippet/utils"
	"github.com/imroc/req/v3"
)

// During file transmission, you can use MD5/SHA256 to calculate the hash value of the file,
// and calculate the hash value again after receiving the file
// for comparison to ensure that the file is not damaged during transmission.

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
	err = CheckIntegrityByMd5(resp, filePath) // or err = CheckIntegrityBySha256(resp, filePath)
	if err != nil {
		return err
	}
	return nil
}

// CheckIntegrityByMd5 Verify the integrity of downloaded files through md5
func CheckIntegrityByMd5(file *req.Response, filePath string) error {
	checksum := utils.GetMd5CheckSum(file.Response)
	decodeChecksum, err := base64.StdEncoding.DecodeString(checksum)
	if err != nil {
		return err
	}
	if checksum != "" {
		fileChecksum, err := utils.CalculateMd5(filePath)
		if err != nil {
			return err
		}

		if !strings.EqualFold(fmt.Sprintf("%x", decodeChecksum), fileChecksum) {
			return errors.New("the download file is incomplete")
		}
	}
	return nil
}

// CheckIntegrityBySha256 Verify the integrity of downloaded files through sha256
func CheckIntegrityBySha256(file *req.Response, filePath string) error {
	checksum := utils.GetSha256CheckSum(file.Response)
	if checksum != "" {
		fileChecksum, err := utils.CalculateSha256(filePath)
		if err != nil {
			return err
		}

		if !strings.EqualFold(checksum, fileChecksum) {
			return errors.New("the download file is incomplete")
		}
	}
	return nil
}
