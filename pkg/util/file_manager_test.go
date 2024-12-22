package util_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mbaeum/advent-of-go-2025/pkg/util"
)

func TestNewFileManager(t *testing.T) {
	// Create a temporary directory to act as the root
	tempDir, err := os.MkdirTemp("", "test_root")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Cleanup

	// Test successful FileManager creation
	fm, err := util.NewFileManger(tempDir, nil)
	if err != nil {
		t.Fatalf("NewFileManger failed: %v", err)
	}
	if fm.GetRoot() != tempDir {
		t.Errorf("Expected root: %s, got: %s", tempDir, fm.GetRoot())
	}
}

func TestNewChallenge(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "test_root")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Cleanup

	// Initialize FileManager
	fm, err := util.NewFileManger(tempDir, nil)
	if err != nil {
		t.Fatalf("NewFileManger failed: %v", err)
	}

	// Test creating a new challenge
	challengeID := 1
	err = fm.NewChallenge(challengeID)
	if err != nil {
		t.Fatalf("NewChallenge failed: %v", err)
	}

	// Verify files in the challenge directory
	expectedFiles := []string{
		"challenge01.go",
		"challenge01_test.go",
	}
	for _, fileName := range expectedFiles {
		filePath := filepath.Join(tempDir, fileName)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist, but it does not", filePath)
		}
	}

	// Verify challenge.go content
	challengeFile := filepath.Join(tempDir, "challenge01.go")
	content, err := os.ReadFile(challengeFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", challengeFile, err)
	}

	if string(content) == "" {
		t.Errorf("Content mismatch in %s:\nExpected:non empty", challengeFile)
	}
}

func TestNewChallenge_SkipExistingFiles(t *testing.T) {
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "test_root")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Cleanup

	// Initialize FileManager
	fm, err := util.NewFileManger(tempDir, nil)
	if err != nil {
		t.Fatalf("NewFileManger failed: %v", err)
	}

	// Create a new challenge
	challengeID := 2
	err = fm.NewChallenge(challengeID)
	if err != nil {
		t.Fatalf("NewChallenge failed: %v", err)
	}

	// Modify an existing file
	existingFile := filepath.Join(tempDir, "challenge02.go")
	err = os.WriteFile(existingFile, []byte("Modified content"), os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to modify file %s: %v", existingFile, err)
	}

	// Run NewChallenge again
	err = fm.NewChallenge(challengeID)
	if err != nil {
		t.Fatalf("NewChallenge failed on second run: %v", err)
	}

	// Ensure the modified file was not overwritten
	content, err := os.ReadFile(existingFile)
	if err != nil {
		t.Fatalf("Failed to read file %s: %v", existingFile, err)
	}
	if string(content) != "Modified content" {
		t.Errorf("Existing file %s was overwritten. Expected content: 'Modified content', got: %s", existingFile, content)
	}
}
