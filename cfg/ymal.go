/*************************************************************************
   > File Name: ymal.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2018.05.23
************************************************************************/
package cfg

import (
	"github.com/keesely/go-extra/files"
	"gopkg.in/yaml.v2"
	//"reflect"
	"strings"
)

type YamlCfg struct {
	cfg  string
	data map[interface{}]interface{}
}

func Yaml(cfg string) (*YamlCfg, error) {
	cfgs, err := files.Get(cfg)

	if err != nil {
		return nil, err
	}

	var data map[interface{}]interface{}

	err = yaml.Unmarshal([]byte(cfgs), &data)

	ycfg := &YamlCfg{
		cfg:  cfg,
		data: data,
	}

	return ycfg, err
}

func (this *YamlCfg) All() interface{} {
	return this.data
}

func (this *YamlCfg) Get(key string, def ...interface{}) interface{} {
	split := strings.Split(key, ":")

	data := this.data

	val := getVal(split, data)
	if val == nil {
		if def != nil {
			return def[0]
		}
		return nil
	}
	return val
}

func (this *YamlCfg) Set(key string, value string) *YamlCfg {
	split := strings.Split(key, ":")
	data := this.data

	for _, sk := range split {
		split = split[1:]
		if _, exists := data[sk]; exists == false {
			data[sk] = make(map[interface{}]interface{})
		}
		if len(split) > 0 {
			tmp := data[sk]
			data = tmp.(map[interface{}]interface{})
		} else {
			data[sk] = value
		}
	}
	return this
}

func (this *YamlCfg) Save() []byte {
	t, _ := yaml.Marshal(this.data)
	return t
}

func getVal(key []string, data map[interface{}]interface{}) interface{} {
	sk := key[0]
	key = key[1:]
	if _, ext := data[sk]; ext {
		val := data[sk]
		switch val.(type) {
		case string:
			return val
		case int:
			return val
		case int32:
			return val
		case int64:
			return val
		default:
			if len(key) == 0 {
				return val
			} else {
				tt := val.(map[interface{}]interface{})
				return getVal(key, tt)
			}
		}
	} else {
		return nil
	}
}
