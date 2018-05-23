/*************************************************************************
   > File Name: ini.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.22
************************************************************************/
package cfg

import (
	"fmt"
	"github.com/keesely/go-extra/files"
	"strings"
)

type IniCfg struct {
	cfg  string
	data map[interface{}]map[string]string `json:"data"`
}

func Ini(cfg string) *IniCfg {
	cfgs, _ := files.Files(cfg)

	this := &IniCfg{
		cfg:  cfg,
		data: make(map[interface{}]map[string]string),
	}

	partter := ""
	this.data[0] = make(map[string]string)
	for _, text := range cfgs {
		text = strings.TrimSpace(text)
		n := len(text)
		if n == 0 {
			continue
		}

		// 处理区域
		if "[" == text[0:1] && "]" == text[n-1:] {
			partter = text[1:]
			partter = strings.Replace(partter, "]", "", -1)

			if _, exists := this.data[partter]; exists == false {
				this.data[partter] = make(map[string]string)
			}
		}

		// 不解析注释
		if "#" == text[0:1] || ";" == text[0:1] {
			continue
		}

		split := strings.SplitN(text, "=", 2)
		if 2 > len(split) {
			continue
		}
		key, value := psVal(split)

		if "" == key {
			continue
		}

		if len(partter) > 0 {
			this.data[partter][key] = value
		} else {
			//this.data[key] = value
			this.data[0][key] = value
		}
	}

	return this
}

func (this *IniCfg) Get(key string, def interface{}) interface{} {
	split := strings.SplitN(key, ":", 2)

	if len(split) == 2 {
		partter := split[0]
		skey := split[1]

		if _, exists := this.data[partter][skey]; exists {
			return this.data[partter][skey]
		} else {
			return def
		}

	} else {
		if _, exists := this.data[0][key]; exists {
			return this.data[0][key]
		} else {
			return def
		}
	}
}

func (this *IniCfg) All(partter ...string) map[string]string {
	if partter == nil {
		return this.data[0]
	}
	return this.data[partter[0]]
}

func (this *IniCfg) Set(key string, value string) *IniCfg {
	split := strings.SplitN(key, ":", 2)
	if len(split) == 2 {
		this.data[split[0]][split[1]] = value
	} else {
		this.data[0][key] = value
	}
	return this
}

func (this *IniCfg) Save(file ...string) bool {
	fn := this.cfg
	if file != nil {
		fn = file[0]
	}

	save, _ := files.Put(fn, this.psToIni(), 0)
	return save
}

func psVal(vals []string) (string, string) {
	key := strings.TrimSpace(vals[0])
	value := strings.TrimSpace(vals[1])
	value = strings.Trim(value, `"`)
	return key, value
}

func (this *IniCfg) psToIni() string {
	text := psToIniString(this.data[0])

	for k, v := range this.data {
		if k == 0 {
			continue
		}

		partter := fmt.Sprintf("%s", k)
		text += "[" + partter + "]\n"
		text += psToIniString(v)
	}

	return text
}

func psToIniString(data map[string]string) string {
	text := ""
	for k, v := range data {
		text += k + " = " + v + "\n"
	}
	return text
}
