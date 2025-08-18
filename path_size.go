package pathsize

import (
	"fmt"
	"os"
	"strings"
)

func GetSize(path string, withHidden bool) (int64, error) {
	info, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}

	if info.IsDir() {
		return calculateDirSize(path, withHidden), nil
	} else {
		return info.Size(), nil
	}
}

func calculateDirSize(path string, withHidden bool) int64 {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0
	}
	var size int64

	for _, entry := range entries {
		info, err := entry.Info()

		if withHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if err != nil {
			continue
		}

		if !entry.IsDir() {
			size += info.Size()
		}
	}

	return size
}

func FormatSize(size int64, human bool) string {
	if human {
		return getHumanSize(size)
	} else {
		return fmt.Sprintf("%dB", size)
	}
}

func getHumanSize(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	unitIndex := 0
	floatSize := float64(size)

	for floatSize >= 1024 && unitIndex < len(units)-1 {
		floatSize /= 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", floatSize, units[unitIndex])
}
