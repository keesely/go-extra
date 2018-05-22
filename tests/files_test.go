/*************************************************************************
   > File Name: files_test.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.22
************************************************************************/
package tests

import (
	"fmt"
	"github.com/keesely/go-extra/files"
	"testing"
)

func Test2(t *testing.T) {
	file := "test.txt"

	if exists := files.Exists(file); exists == true {
		contents, _ := files.Files(file)
		//i := 0
		for i, str := range contents {
			//i++
			fmt.Println("LN", i, str)
		}
	}
}
