/*************************************************************************
   > File Name: file_test.go
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

func Test(t *testing.T) {
	file := "./test.txt"
	// 文件是否存在 - 不存在确认
	exists := files.Exists(file)
	if exists == false {
		fmt.Println(file, "文件不存在")
	} else {
		if err := files.Remove(file); err == nil {
			fmt.Println("移除存在的文件")
		}
	}

	// 不追加写入
	_, err := files.Put(file, "测试文本BEGIN\n", 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("文件写入")
	}

	if false == files.Exists(file) {
		fmt.Println(file, "文件还是不存在")
	} else {
		fmt.Println(file, "文件已存在")
	}

	// 追加写入
	_, err = files.Put(file, "追加文本\n", files.APPEND)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("追加写入成功")
	}

	_, err = files.Put(file, "ENDING.", files.APPEND)
	if err != nil {
		fmt.Println("End")
	}

	if false != files.Exists(file) {
		str, err := files.Get(file)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Print("读取文本内容\n", "```\n", str, "\n```\n")

			if err = files.Copy(file, "new-test.txt"); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("文件复制成功")

				if err = files.Move("new-test.txt", "move.txt"); err == nil {
					fmt.Println("文件转移成功")
				}
			}
		}
	}

}
