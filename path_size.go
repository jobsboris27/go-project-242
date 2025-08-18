package pathsize

import (
	"os"
)

func GetSize(path string) (int64, error) {
	info, err := os.Lstat(path)

	if err != nil {
		return 0, err
	}

	if info.IsDir() {
		return calculateDirSize(path), nil
	} else {
		return info.Size(), nil
	}
}

func calculateDirSize(path string) int64 {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0
	}
	var size int64

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		if !entry.IsDir() {
			size += info.Size()
		}
	}

	return size
}
