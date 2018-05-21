/*************************************************************************
   > File Name: file.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.21
************************************************************************/
package files

import (
	"io"
	"io/ioutil"
	"os"
)

const (
	APPEND = (os.O_APPEND | os.O_WRONLY)
	WRONLY = os.O_WRONLY
)

func GET(file string) (string, error) {
	handle, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer handle.Close()

	content, err := ioutil.ReadAll(handle)
	return string(content), err
}

func PUT(file string, content string, _append int) (bool, error) {
	if true != Exists(file) {
		_, err := os.Create(file)

		if err != nil {
			return false, err
		}
	}

	op := os.O_WRONLY | os.O_TRUNC
	if _append != 0 {
		op = _append
	}

	handle, err := os.OpenFile(file, op, 0755)

	nr, err := io.WriteString(handle, content)
	defer handle.Close()

	if err != nil {
		return false, err
	}
	return bool(nr > len(content)), err
}

func Exists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
