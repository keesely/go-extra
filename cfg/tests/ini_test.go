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

	fmt.Println(cfg.String("mongo:User", "default"))
	fmt.Println(cfg.String("mongo:AAA", 0))
	fmt.Println(cfg.String("Admin", "Locel"))
	fmt.Println(cfg.String("mongo:Admin", "NOFOUND"))
	fmt.Println(cfg.String("Admin:DNS", "NOFOUND"))
	fmt.Println(cfg.String("Admin:QUI", "NOFOUND"))

	cfg.SetString("production:ADDKEY", "110110100101")

	fmt.Println("PRINT TO `production`")
	data := cfg.All("production")
	for k, v := range data {
		fmt.Println(k, "=>", v)
	}

	cfg.SetString("ADDKEY", "New&KEY:Value")
	fmt.Println("PRINT TO `nil`")
	data2 := cfg.All()
	for k, v := range data2 {
		fmt.Println(k, "=>", v)
	}

	fmt.Println(cfg.Save("new.ini"))
	fmt.Println(cfg.Save())
}
