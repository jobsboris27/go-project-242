package pathsize_test

import (
	pathsize "code"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSize_File(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "testFile.txt"
	filePath := tempDir + "/" + fileName
	var targetSize int64 = 100

	createFileWithSize(t, filepath.Join(tempDir, fileName), targetSize)

	size, err := pathsize.GetSize(filePath)

	require.Nil(t, err)
	require.Equal(t, size, targetSize)
}

func TestGetSize_Directory(t *testing.T) {
	tempDir := t.TempDir()
	files := []string{"dir.txt", "dir1.txt", "dir2.txt"}

	var targetSize int64 = 20

	for _, file := range files {
		createFileWithSize(t, filepath.Join(tempDir, file), targetSize)
	}

	size, err := pathsize.GetSize(tempDir)

	require.NoError(t, err)
	require.Equal(t, size, targetSize*int64(len(files)))
}

func TestGetSize_NonExistentPath(t *testing.T) {
	filePath := "./non_existent_file.txt"
	_, err := pathsize.GetSize(filePath)

	require.NotNil(t, err)
}

func TestCalculateDirSize_Empty(t *testing.T) {
	tempDir := t.TempDir()

	size, err := pathsize.GetSize(tempDir)

	require.NoError(t, err)
	require.Equal(t, int64(0), size)
}

func createFileWithSize(t *testing.T, path string, size int64) {
	t.Helper()

	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	if err := file.Truncate(size); err != nil {
		t.Fatalf("Failed to set file size: %v", err)
	}
}
