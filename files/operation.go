/*************************************************************************
   > File Name: operation.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.22
************************************************************************/
package files

import (
	"io"
	"os"
)

// 拷贝文件
func Copy(src string, dst string) error {
	handle, err := os.Open(src)
	if err != nil {
		return err
	}
	defer handle.Close()

	dHandle, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dHandle.Close()

	if _, err := io.Copy(dHandle, handle); err != nil {
		return err
	}
	return nil
}

// 迁移文件
func Move(src string, dst string) error {
	return os.Rename(src, dst)
}

// 移除文件
func Remove(src string) error {
	return os.Remove(src)
}
