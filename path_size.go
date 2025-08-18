package pathsize

import (
	"fmt"
	"log"
	"os"
)

func GetSize(path string) (string, error) {
	info, err := os.Lstat(path)

	if err != nil {
		log.Fatal(err)
	}

	if info.IsDir() {
		return fmt.Sprintf("%d\t%s\n", calculateDirSize(path), path), nil
	} else {
		return fmt.Sprintf("%d\t%s\n", info.Size(), path), nil
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
