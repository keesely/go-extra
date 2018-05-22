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
}
