/*************************************************************************
   > File Name: file_test.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.22
************************************************************************/
package files

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// 文件是否存在 - 不存在确认
	exists := Exists("./test.txt")
	if exists == false {
		fmt.Println("test.txt 文件不存在")
	}

	// 不追加写入
	_, err := PUT("./test.txt", "测试文本BEGIN\n", 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("文件写入")
	}

	if false == Exists("./test.txt") {
		fmt.Println("test.txt 文件还是不存在")
	} else {
		fmt.Println("test.txt 文件已存在")
	}

	// 追加写入
	_, err = PUT("./test.txt", "追加文本\n", APPEND)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("追加写入成功")
	}

	_, err = PUT("./test.txt", "ENDING.", APPEND)
	if err != nil {
		fmt.Println("End")
	}

	if false != Exists("./test.txt") {
		str, err := GET("./test.txt")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Print("读取文本内容\n", "```\n", str, "\n```\n")
		}
	}

}
