package files

import (
	"errors"
	"fmt"
	"os"
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
