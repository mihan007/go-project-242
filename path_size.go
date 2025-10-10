package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, recursive, all)
	if err != nil {
		return "", fmt.Errorf("error processing %s: %v", path, err)
	}

	if human {
		return humanReadableSize(size), nil
	}
	return fmt.Sprintf("%dB", size), nil
}

func getSize(path string, recursive, all bool) (int64, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	isDir := fileInfo.IsDir()
	if isDir {
		files, err := os.ReadDir(path)
		if err != nil {
			return 0, err
		}
		size := int64(0)
		for _, file := range files {
			if !all && strings.HasPrefix(file.Name(), ".") {
				continue
			}
			fileInfo, err := file.Info()
			if err != nil {
				fmt.Printf("Error getting file info for %s: %s\n", file.Name(), err)
				continue
			}
			if fileInfo.IsDir() {
				if recursive {
					s, err := getSize(filepath.Join(path, file.Name()), all, recursive)
					if err != nil {
						continue
					}
					size += s
				}
			} else {
				size += fileInfo.Size()
			}
		}
		return size, nil
	} else {
		return fileInfo.Size(), nil
	}
}

func humanReadableSize(bytes int64) string {
	switch {
	case bytes >= 1<<30: // 1 GB
		return fmt.Sprintf("%.1fGB", float64(bytes)/(1<<30))
	case bytes >= 1<<20: // 1 MB
		return fmt.Sprintf("%.1fMB", float64(bytes)/(1<<20))
	case bytes >= 1<<10: // 1 KB
		return fmt.Sprintf("%.1fKB", float64(bytes)/(1<<10))
	default:
		return fmt.Sprintf("%dB", bytes)
	}
}
