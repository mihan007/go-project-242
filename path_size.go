package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := GetSize(path, all, recursive)
	if err != nil {
		return "", err
	}

	if human {
		return HumanReadableSize(size), nil
	}
	return fmt.Sprintf("%dB", size), nil
}

func GetSize(path string, all, recursive bool) (int64, error) {
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
				return 0, err
			}
			if fileInfo.IsDir() {
				if recursive {
					s, err := GetSize(filepath.Join(path, file.Name()), all, recursive)
					if err != nil {
						return 0, err
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

func HumanReadableSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%dB", bytes)
	}

	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	size := float64(bytes)
	i := 0
	for size >= unit && i < len(sizes)-1 {
		size /= unit
		i++
	}

	if size == float64(int64(size)) {
		return fmt.Sprintf("%d%s", int64(size), sizes[i])
	}
	return fmt.Sprintf("%.1f%s", size, sizes[i])
}
