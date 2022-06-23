package rootpath

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// 取得資料夾目錄
func GetFilePath(file string) (path string, err error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return currentPath, errors.New(fmt.Sprintf("[os.Getwd]: %v", err))
	}
	index := strings.Index(currentPath, file)
	if index == -1 {
		return currentPath, errors.New(fmt.Sprintf("[File not found in path] file: %v, path: %v\n", file, currentPath))
	}
	path = currentPath[:index] + file
	return path, nil
}
