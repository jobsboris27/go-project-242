package main_tests

import (
	code "code"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	KB        = 1024
	MB        = KB * 1024
	GB        = MB * 1024
	TB        = GB * 1024
	sizeFile1 = 100
	sizeFile2 = 200
	sizeFile3 = 150
	sizeFile4 = 50
)

func TestGetSize_File(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "testFile.txt"
	filePath := tempDir + "/" + fileName
	var targetSize int64 = 100

	createFileWithSize(t, filepath.Join(tempDir, fileName), targetSize)

	size, err := code.GetSize(filePath, false, false)

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

	size, err := code.GetSize(tempDir, false, false)

	require.NoError(t, err)
	require.Equal(t, size, targetSize*int64(len(files)))
}

func TestGetSize_NonExistentPath(t *testing.T) {
	filePath := "./non_existent_file.txt"
	_, err := code.GetSize(filePath, false, false)

	require.NotNil(t, err)
}

func TestGetSize_EmptyDir(t *testing.T) {
	tempDir := t.TempDir()

	size, err := code.GetSize(tempDir, false, false)

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
		message := code.FormatSize(size, true)
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
		message := code.FormatSize(size, false)
		require.Equal(t, size, sizeNames[message], "Size should match the expected value")
	}
}

func TestCalculateDirSize_WithHidden(t *testing.T) {
	t.Run("size with hidden files", func(t *testing.T) {
		tempDir := t.TempDir()

		createFileWithSize(t, filepath.Join(tempDir, "visible.txt"), 100)
		createFileWithSize(t, filepath.Join(tempDir, ".hidden"), 50)

		size, _ := code.GetSize(tempDir, true, false)
		require.Equal(t, int64(150), size)
	})

	t.Run("size without hidden files", func(t *testing.T) {
		tempDir := t.TempDir()

		createFileWithSize(t, filepath.Join(tempDir, "visible.txt"), 100)
		createFileWithSize(t, filepath.Join(tempDir, ".hidden"), 50)

		size, _ := code.GetSize(tempDir, false, false)
		require.Equal(t, int64(100), size)
	})

	t.Run("only hidden files", func(t *testing.T) {
		tempDir := t.TempDir()

		createFileWithSize(t, filepath.Join(tempDir, ".hidden1"), 30)
		createFileWithSize(t, filepath.Join(tempDir, ".hidden2"), 40)

		size, _ := code.GetSize(tempDir, true, false)
		require.Equal(t, int64(70), size)
	})
}

func TestCalculateDirSize_Recursive(t *testing.T) {
	setupTestEnv := func(t *testing.T) string {
		tempDir := t.TempDir()

		createFileWithSize(t, filepath.Join(tempDir, "file1.txt"), sizeFile1)
		createFileWithSize(t, filepath.Join(tempDir, "file2.txt"), sizeFile2)

		subDir := filepath.Join(tempDir, "subdir")
		require.NoError(t, os.Mkdir(subDir, 0755))
		createFileWithSize(t, filepath.Join(subDir, "file3.txt"), sizeFile3)

		nestedDir := filepath.Join(subDir, "nested")
		require.NoError(t, os.Mkdir(nestedDir, 0755))
		createFileWithSize(t, filepath.Join(nestedDir, "file4.txt"), sizeFile4)

		return tempDir
	}

	t.Run("non-recursive counts only root files", func(t *testing.T) {
		dir := setupTestEnv(t)
		size, err := code.GetSize(dir, true, false)
		expected := int64(sizeFile1 + sizeFile2)

		require.NoError(t, err)
		require.Equal(t, expected, size)
	})

	t.Run("recursive counts all files", func(t *testing.T) {
		dir := setupTestEnv(t)
		size, err := code.GetSize(dir, true, true)
		expected := int64(sizeFile1 + sizeFile2 + sizeFile3 + sizeFile4)

		require.NoError(t, err)
		require.Equal(t, expected, size)
	})
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
