/*************************************************************************
   > File Name: ini.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.22
************************************************************************/
package cfg

import (
	"github.com/keesely/go-extra/files"
	"strings"
)

type Cfg struct {
	multiData map[string]map[string]string
	data      map[string]string
}

func Ini(cfg string) *Cfg {
	cfgs, _ := files.Files(cfg)

	this := &Cfg{
		multiData: make(map[string]map[string]string),
		data:      make(map[string]string),
	}

	partter := ""
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

			if _, exists := this.multiData[partter]; exists == false {
				this.multiData[partter] = make(map[string]string)
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
			this.multiData[partter][key] = value
		} else {
			this.data[key] = value
		}
	}

	return this
}

func (this *Cfg) String(key string, def interface{}) interface{} {
	split := strings.SplitN(key, ":", 2)

	if len(split) == 2 {
		partter := split[0]
		skey := split[1]

		if _, exists := this.multiData[partter][skey]; exists {
			return this.multiData[partter][skey]
		} else {
			return def
		}

	} else {
		if _, exists := this.data[key]; exists {
			return this.data[key]
		} else {
			return def
		}
	}
}

func psVal(vals []string) (string, string) {
	key := strings.TrimSpace(vals[0])
	value := strings.TrimSpace(vals[1])
	return key, value
}
