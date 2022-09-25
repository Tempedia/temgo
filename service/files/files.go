package files

import (
	"errors"
	"fmt"
	"os"
	"path"
)

var _filesfolder string

func Init(folder string) error {
	_filesfolder = folder
	if info, err := os.Stat(_filesfolder); err == nil {
		if !info.IsDir() {
			return fmt.Errorf("文件目录不是目录")
		}
	} else if errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("文件目录不存在")
	} else {
		return fmt.Errorf("文件目录不合法")
	}
	return nil
}

func Folder() string {
	return _filesfolder
}

func FilePath(file string) string {
	return path.Join(_filesfolder, file)
}
