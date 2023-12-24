package helper

import (
	"fmt"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type FileResponse struct {
	FilePath string `json:"file_path"`
}

func GetDownloadLink(b *gotgbot.Bot, fileId string) string {
	filePath, err := b.GetFile(fileId, nil)
	if err != nil {
		return ""
	}

	downloadLink := filePath.URL(b, nil)

	return downloadLink
}

func DownloadFile(url string) (string, error) {
	log.Printf("DownloadFile: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error while downloading file: %v", err)
	}
	defer resp.Body.Close()

	// Create a temporary file
	tempPath := filepath.Join(os.TempDir(), filepath.Base(url))
	file, err := os.Create(tempPath)
	if err != nil {
		return "", fmt.Errorf("error while creating file: %v", err)
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return "", fmt.Errorf("error while copying file: %v", err)
	}

	return tempPath, nil
}

func StringToInt64(s string) (int64, bool) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, false
	}
	return i, true
}
