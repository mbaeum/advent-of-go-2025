package util

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetTargetPath(name string) (string, error) {
	p, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}
	p = filepath.Join(p, name)
	return p, nil
}

type FileManager struct {
	r   string
	l   *slog.Logger
	cfg *Config
}

// Creates a new FileManager which can manipulate files
// in the `path` folder.
func NewFileManger(path string, cfg *Config) (*FileManager, error) {
	l := NewLogger()
	return &FileManager{path, l, cfg}, nil
}

func NewChallengeFileManager(id int, cfg *Config) (*FileManager, error) {
	path, err := GetTargetPath(fmt.Sprintf("challenges/challenge%02d", id))
	if err != nil {
		return nil, err
	}
	return NewFileManger(path, cfg)

}

func (f *FileManager) GetRoot() string {
	return f.r
}

// Creates a new challenge package based on the id given
func (fm *FileManager) NewChallenge(id int) error {

	fileNames := []string{
		fmt.Sprintf("challenge%02d.go", id),
		fmt.Sprintf("challenge%02d_test.go", id),
	}

	for _, fileName := range fileNames {
		filePath := filepath.Join(fm.r, fileName)
		if checkFileExists(filePath) {
			fm.l.Warn("File already exists, skipping", "file", filePath)
			continue
		}
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		fm.l.Debug("Created File", "file", filePath)
		filePathParts := strings.Split(fileName, ".")
		if filePathParts[1] == "go" {

			var content string
			i := templaterInput{Id: id}

			if !strings.Contains(filePathParts[0], "_test") {
				content, err = createChallengeFileString(i, challengeFileTemplate)
			} else {
				content, err = createChallengeFileString(i, challengeFileTestTemplate)
			}
			if err != nil {
				return err
			}

			_, err := file.WriteString(content)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func GetChallengeData(id int, sc string) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sc))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
