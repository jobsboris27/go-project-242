package pathsize_test

import (
	pathsize "code"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	KB = 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
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

func TestGetSize_EmptyDir(t *testing.T) {
	tempDir := t.TempDir()

	size, err := pathsize.GetSize(tempDir)

	require.NoError(t, err)
	require.Equal(t, int64(0), size)
}

func TestFormatSize_Human(t *testing.T) {
	sizes := []int64{100, 1024, 1048576, 1073741824, 1099511627776}
	sizeNames := map[string]int64{
		"100B":  100,
		"1.0KB": KB,
		"1.0MB": MB,
		"1.0GB": GB,
		"1.0TB": TB,
	}

	for _, size := range sizes {
		message := pathsize.FormatSize(size, true)
		require.Equal(t, size, sizeNames[message], "Size should match the expected value")
	}
}

func TestFormatSize_NotHuman(t *testing.T) {
	sizes := []int64{100, 1024, 1048576, 1073741824, 1099511627776}
	sizeNames := map[string]int64{
		"100B":           100,
		"1024B":          KB,
		"1048576B":       MB,
		"1073741824B":    GB,
		"1099511627776B": TB,
	}

	for _, size := range sizes {
		message := pathsize.FormatSize(size, false)
		require.Equal(t, size, sizeNames[message], "Size should match the expected value")
	}
}

func createFileWithSize(t *testing.T, path string, size int64) {
	t.Helper()

	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			t.Errorf("Failed to close file: %v", err)
		}
	}()

	if err := file.Truncate(size); err != nil {
		t.Fatalf("Failed to set file size: %v", err)
	}
}
