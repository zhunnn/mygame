package rootpath

import (
	"fmt"
	"os"
	"strings"
)

// 取得資料夾目錄
func GetFilePath(file string) (path string, err error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return currentPath, err
	}
	index := strings.Index(currentPath, file)
	if index == -1 {

		return currentPath, fmt.Errorf("file: %v, not found in path: %v\n", file, currentPath)
	}
	path = currentPath[:index] + file
	return path, nil
}
