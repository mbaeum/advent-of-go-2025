package util

import (
	"errors"
	"fmt"
	"log/slog"
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
	r string
	l *slog.Logger
}

// Creates a new FileManager which can manipulate files
// in the `path` folder.
func NewFileManger(path string) (*FileManager, error) {
	l := NewLogger()
	return &FileManager{path, l}, nil
}

func (f FileManager) GetRoot() string {
	return f.r
}

// Creates a new challenge package based on the id given
func (f FileManager) NewChallenge(id int) error {
	d := fmt.Sprintf("%s/challenge%02d", f.r, id)
	err := os.MkdirAll(d, os.ModePerm)
	if err != nil {
		return err
	}
	f.l.Debug("Created directory", "dir", d)

	fileNames := []string{
		"data_test.txt",
		"data.txt",
		"challenge.go",
		"challenge_test.go",
	}

	for _, fileName := range fileNames {
		filePath := filepath.Join(d, fileName)
		if checkFileExists(filePath) {
			f.l.Warn("File already exists, skipping", "file", filePath)
			continue
		}
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		f.l.Debug("Created File", "file", filePath)
		filePathParts := strings.Split(fileName, ".")
		if filePathParts[1] == "go" {

			var content string

			if !strings.Contains(filePathParts[0], "_test") {
				content += "package challenges\n\n"
				content += fmt.Sprintf("type Challenge%02d struct { }\n\n", id)
			} else {
				content += "package challenges_test\n"
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
