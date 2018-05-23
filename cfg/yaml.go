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
	"reflect"
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

func (this *YamlCfg) ToString() string {
	t, _ := yaml.Marshal(this.data)
	return string(t)
}

func (this *YamlCfg) Get(key string, def ...interface{}) interface{} {
	split := strings.Split(key, ".")
	data := this.data

	defVal := func(val []interface{}) interface{} {
		if val != nil {
			return val[0]
		}
		return nil
	}

	for _, sk := range split {
		split = split[1:]
		if _, exists := data[sk]; exists {
			v := reflect.TypeOf(data[sk])
			if v.Kind() == reflect.Array || len(split) == 0 {
				return data[sk]
			}

			if v.Kind() == reflect.String {
				return defVal(def)
			}

			val := data[sk]
			data = val.(map[interface{}]interface{})
		} else {
			return defVal(def)
		}
	}

	return data

	// val := getVal(split, data)
	// if val == nil {
	//  return defVal(def)
	// }
	// return val
}

func (this *YamlCfg) Set(key string, value interface{}) *YamlCfg {
	split := strings.Split(key, ".")
	data := this.data

	for _, sk := range split {
		split = split[1:]
		if _, exists := data[sk]; exists == false {
			data[sk] = make(map[interface{}]interface{})
		}
		if len(split) > 0 {
			v := reflect.TypeOf(data[sk])
			if v.Kind() == reflect.String {
				data[sk] = make(map[interface{}]interface{})
			}
			tmp := data[sk]
			data = tmp.(map[interface{}]interface{})
		} else {
			data[sk] = value
		}
	}
	return this
}

func (this *YamlCfg) Save(file ...string) bool {
	t, _ := yaml.Marshal(this.data)
	fn := this.cfg
	if file != nil {
		fn = file[0]
	}

	save, _ := files.Put(fn, string(t), 0)
	return save
}

func getVal(key []string, data map[interface{}]interface{}) interface{} {
	sk := key[0]
	key = key[1:]
	if _, ext := data[sk]; ext {
		val := data[sk]
		switch val.(type) {
		case []interface{}:
			return val
		case interface{}:
			if len(key) == 0 {
				return val
			} else {
				tt := val.(map[interface{}]interface{})
				return getVal(key, tt)
			}
		default:
			return val
		}
	} else {
		return nil
	}
}
