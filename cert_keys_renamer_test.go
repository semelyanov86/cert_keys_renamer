package main

import (
	"os"
	"path/filepath"
	"testing"
)

func createTestFiles(dir string) {
	// Create test files in the temporary directory
	fileNames := []string{"chain.pem", "chain1.pem", "chain2.pem", "cert.pem", "cert1.pem", "cert2.pem"}
	for _, fileName := range fileNames {
		filePath := filepath.Join(dir, fileName)
		err := os.WriteFile(filePath, []byte{}, 0644)
		if err != nil {
			panic(err)
		}
	}

	// Change current working directory to the temporary directory
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

}

func TestCopyLatestMatchingFile(t *testing.T) {
	// Prepare the command line arguments
	find := "chain.pem"
	newName := "check.crt"
	dir, err := os.MkdirTemp("", "test-certs")
	defer os.RemoveAll(dir)
	createTestFiles(dir)
	if err != nil {
		t.Errorf("Error reading new file: %v", err)
	}
	dir2, err := os.MkdirTemp("", "result-certs")
	if err != nil {
		t.Errorf("Error reading new file: %v", err)
	}

	// Run the script
	err = FindAndCopyFile(find, newName, &dir, &dir2)
	if err != nil {
		t.Errorf("Error reading new file: %v", err)
	}
	// Check if the new file was created and has the correct content
	content, err := os.ReadFile(dir2 + "/check.crt")
	if err != nil {
		t.Errorf("Error reading new file: %v", err)
	}
	if len(content) > 0 {
		t.Errorf("New file should be empty but contains %d bytes", len(content))
	}
}

func TestMissingRequiredFlags(t *testing.T) {
	// Prepare the command line arguments
	find := ""
	newName := ""
	dir, err := os.MkdirTemp("", "test-certs")
	if err != nil {
		panic(err)
	}
	dir2, err := os.MkdirTemp("", "result-certs")
	if err != nil {
		t.Errorf("Error reading new file: %v", err)
	}
	defer os.RemoveAll(dir)
	createTestFiles(dir)

	err = FindAndCopyFile(find, newName, &dir, &dir2)
	if err == nil {
		t.Error("Expected to be an error that files not found")
	}
}
