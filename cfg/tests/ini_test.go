/*************************************************************************
   > File Name: ini_test.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.23
************************************************************************/
package tests

import (
	"fmt"
	"github.com/keesely/go-extra/cfg"
	"testing"
)

func Test(t *testing.T) {
	ini := "app.ini"

	cfg := cfg.Ini(ini)

	fmt.Println(cfg)

	fmt.Println(cfg.Get("mongo:User", "default"))
	fmt.Println(cfg.Get("mongo:AAA", 0))
	fmt.Println(cfg.Get("Admin", "Locel"))
	fmt.Println(cfg.Get("mongo:Admin", "NOFOUND"))
	fmt.Println(cfg.Get("Admin:DNS", "NOFOUND"))
	fmt.Println(cfg.Get("Admin:QUI", "NOFOUND"))

	cfg.Set("production:ADDKEY", 110110100101)
	cfg.Set("production:ADDKEY2", []string{"hi admin", "hello", "world"})
	cfg.Set("production:ADDINT", []int{1, 2, 3, 4, 5, 6})

	fmt.Println("PRINT TO `production`")
	data := cfg.All("production")
	for k, v := range data {
		fmt.Println(k, "=>", v)
	}

	cfg.Set("ADDKEY", "New&KEY:Value")
	fmt.Println("PRINT TO `nil`")

	data2 := cfg.All()
	for k, v := range data2 {
		fmt.Println(k, "=>", v)
	}

	fmt.Println(cfg.Save("new.ini"))
	fmt.Println(cfg.Save())
}
