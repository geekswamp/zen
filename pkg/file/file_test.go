package file_test

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/geekswamp/zen/pkg/errors"
	"github.com/geekswamp/zen/pkg/file"
	"github.com/stretchr/testify/assert"
)

func TestNewReadLineFromEnd(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "readlinetest")
	assert.NoError(t, err, "Failed to create temp directory")
	defer os.RemoveAll(tmpDir)

	tests := []struct {
		name     string
		content  string
		wantErr  bool
		errCheck func(error) bool
	}{
		{
			name:    "Valid file with content",
			content: "Line 1\nLine 2\nLine 3",
			wantErr: false,
		},
		{
			name:    "Empty file",
			content: "",
			wantErr: false,
		},
		{
			name:    "Non-existent file",
			content: "",
			wantErr: true,
			errCheck: func(err error) bool {
				return os.IsNotExist(err)
			},
		},
		{
			name:    "Directory instead of file",
			content: "",
			wantErr: true,
			errCheck: func(err error) bool {
				return err == errors.ErrNotAFile
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var path string

			switch tt.name {
			case "Valid file with content", "Empty file":
				file, err := os.CreateTemp(tmpDir, "test_file_*.txt")
				assert.NoError(t, err, "Failed to create temp file")
				defer file.Close()

				_, err = file.WriteString(tt.content)
				assert.NoError(t, err, "Failed to write to temp file")
				path = file.Name()

			case "Non-existent file":
				path = filepath.Join(tmpDir, "non_existent_file.txt")

			case "Directory instead of file":
				path = tmpDir
			}

			rd, err := file.NewReadLineFromEnd(path)

			if tt.wantErr {
				assert.Error(t, err, "Expected an error for case: %s", tt.name)
				if tt.errCheck != nil {
					assert.True(t, tt.errCheck(err), "Expected specific error condition for case: %s", tt.name)
				}
				return
			}

			assert.NoError(t, err, "Unexpected error for case: %s", tt.name)
			assert.NotNil(t, rd, "Expected non-nil reader for case: %s", tt.name)
		})
	}
}

func TestIsExist(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "existtest")
	assert.NoError(t, err, "Failed to create temp directory")
	defer os.RemoveAll(tmpDir)

	testFile := filepath.Join(tmpDir, "test_file.txt")
	err = os.WriteFile(testFile, []byte("test content"), 0644)
	assert.NoError(t, err, "Failed to create test file")

	testDir := filepath.Join(tmpDir, "test_dir")
	err = os.Mkdir(testDir, 0755)
	assert.NoError(t, err, "Failed to create test directory")

	tests := []struct {
		name      string
		path      string
		wantExist bool
		isDir     bool
	}{
		{
			name:      "Existing file",
			path:      testFile,
			wantExist: true,
			isDir:     false,
		},
		{
			name:      "Existing directory",
			path:      testDir,
			wantExist: true,
			isDir:     true,
		},
		{
			name:      "Non-existent path",
			path:      filepath.Join(tmpDir, "non_existent_file.txt"),
			wantExist: false,
			isDir:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fileInfo, exists := file.IsExist(tt.path)

			assert.Equal(t, tt.wantExist, exists, "Expected correct existence status for case: %s", tt.name)

			if !exists {
				assert.Nil(t, fileInfo, "Expected nil fileInfo for non-existent path in case: %s", tt.name)
				return
			}

			assert.Equal(t, tt.isDir, fileInfo.IsDir(), "Expected correct IsDir status for case: %s", tt.name)
		})
	}
}

func BenchmarkNewReadLineFromEnd(b *testing.B) {
	tmpFile, err := os.CreateTemp("", "benchmark_file_*.txt")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	content := bytes.Repeat([]byte("This is a line of text for benchmarking purposes.\n"), 1000)
	if _, err := tmpFile.Write(content); err != nil {
		b.Fatalf("Failed to write to temp file: %v", err)
	}

	if err := tmpFile.Sync(); err != nil {
		b.Fatalf("Failed to sync file: %v", err)
	}

	b.ResetTimer()
}

func BenchmarkIsExist(b *testing.B) {
	tmpFile, err := os.CreateTemp("", "benchmark_exist_*.txt")
	if err != nil {
		b.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	path := tmpFile.Name()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, exists := file.IsExist(path)
		if !exists {
			b.Fatalf("File should exist: %s", path)
		}
	}
}

func BenchmarkIsExistNonExistent(b *testing.B) {
	path := "/tmp/non_existent_file_for_benchmark.txt"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, exists := file.IsExist(path)
		if exists {
			b.Fatalf("File should not exist: %s", path)
		}
	}
}
