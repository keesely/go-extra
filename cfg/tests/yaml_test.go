/*************************************************************************
   > File Name: yaml_test.go
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

func Test2(t *testing.T) {
	yaml := "example.yml"

	cfg, _ := cfg.Yaml(yaml)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("===========================================")

	fmt.Println(cfg.Get("version"))
	fmt.Println(cfg.Get("services:node1"))
	fmt.Println(cfg.Get("services:node1.data"))

	fmt.Println(cfg.Set("version", "1122222"))
	fmt.Println(cfg.Set("services:data", "1122222"))
	fmt.Println(cfg.Set("services:DBS:DW:dws", "1122222"))
	//fmt.Println(cfg.Get("services:DBS:data"))

	fmt.Println(string(cfg.Save()))
}
