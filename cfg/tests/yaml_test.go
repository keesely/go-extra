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

	cfgs, _ := cfg.Yaml(yaml)

	fmt.Println("===========================================")
	fmt.Println("================ YAML Parse ===============")
	fmt.Println("===========================================")
	fmt.Println("===========================================")

	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- GET version ----------------")
	fmt.Println(cfgs.Get("version"))
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- GET A.Aa -------------------")
	fmt.Println(cfgs.Get("A.Aa"))
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- GET A.Ab.b2 ----------------")
	fmt.Println(cfgs.Get("A.Ab.b2"))
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- GET Int --------------------")
	fmt.Println(cfgs.Get("A.Ab.int"))
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- GET Default ----------------")
	fmt.Println(cfgs.Get("A.Ab.b1.nif"))
	fmt.Println(cfgs.Get("A.Ab.b1.nif", "DEFAULT"))
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- SET Value ------------------")
	fmt.Println(cfgs.Set("version", "1024.1"))

	fmt.Println(cfgs.Set("services.data", "ADD TO DATA"))
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- SET Value2 -----------------")
	fmt.Println(cfgs.Set("A.Ab.b1.ba", "add sub node"))
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- SET Value3 -----------------")
	fmt.Println(cfgs.Set("A.Ab.b1.ba", []string{"add value 1", "Val 2"}))
	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- SET Value4 -----------------")
	fmt.Println(cfgs.Set("A.Ab.b1.int", []int{1, 2, 3, 4, 5}))

	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- Save to New ----------------")
	fmt.Println(cfgs.Save("example2.yml"))

	fmt.Println("-------------------------------------------")
	fmt.Println("-------------- Show to All ----------------")
	ncfg, _ := cfg.Yaml("example2.yml")
	fmt.Println(ncfg.ToString())

}
