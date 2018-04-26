package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/qxnw/lib4go/concurrent/cmap"
	"github.com/qxnw/lib4go/jsons"
	"github.com/qxnw/lib4go/transform"
)

//JSONConf json配置文件
type JSONConf struct {
	data    map[string]interface{}
	cache   cmap.ConcurrentMap
	Content string
	version int32
	*transform.Transform
}

//NewJSONConfWithJson 根据JSON初始化conf对象
func NewJSONConfWithJson(c string, version int32, i interface{}) (r *JSONConf, err error) {
	m := make(map[string]interface{})
	err = json.Unmarshal([]byte(c), &m)
	if err != nil {
		return
	}
	return &JSONConf{
		Content:   c,
		data:      m,
		cache:     cmap.New(8),
		Transform: transform.NewMaps(m),
		version:   version,
	}, nil
}

//NewJSONConfWithEmpty 初始化空的conf对象
func NewJSONConfWithEmpty() *JSONConf {
	return NewJSONConfWithHandle(make(map[string]interface{}), 0)
}

//NewJSONConfWithHandle 根据map和动态获取函数构建
func NewJSONConfWithHandle(m map[string]interface{}, version int32) *JSONConf {

	return &JSONConf{
		data:      m,
		cache:     cmap.New(8),
		Transform: transform.NewMaps(m),
		version:   version,
	}
}

//GetData 获取原始数据
func (j *JSONConf) GetData() map[string]interface{} {
	r, _ := jsons.Unmarshal([]byte(j.Content))
	return r
}

//GetContent 获取输入JSON原串
func (j *JSONConf) GetContent() string {
	return j.Content
}

//Len 获取参数个数
func (j *JSONConf) Len() int {
	return len(j.data)
}

//GetVersion 获取当前配置的版本号
func (j *JSONConf) GetVersion() int32 {
	return j.version
}

//Set 设置参数值
func (j *JSONConf) Set(key string, value string) {
	if _, ok := j.data[key]; ok {
		return
	}
	j.Transform.Set(key, value)
}

//String 获取字符串，已缓存则从缓存中获取
func (j *JSONConf) String(key string, def ...string) (r string) {
	if v, err := j.Transform.Get(key); err == nil {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

//GetArray 获取数组列表
func (j *JSONConf) GetArray(key string) (r []interface{}, err error) {
	d, ok := j.data[key]
	if !ok {
		err = fmt.Errorf("不包含数据:%s", key)
		return
	}
	if r, ok := d.([]interface{}); ok {
		return r, nil
	}
	err = fmt.Errorf("不包含数据:%s", key)
	return
}

//Strings 获取字符串数组，原字符串以“;”号分隔
func (j *JSONConf) Strings(key string, def ...[]string) (r []string) {
	stringVal := j.String(key)
	if stringVal != "" {
		r = strings.Split(j.String(key), ";")
		return
	}
	if len(def) > 0 {
		return def[0]
	}
	return []string{}
}

//Each 循化每个节点值
func (j *JSONConf) Each(f func(key string)) {
	for k := range j.data {
		f(k)
	}
}

//GetSMap 获取map数据
func (j *JSONConf) GetSMap(section string) (map[string]string, error) {
	data := make(map[string]string)
	val := j.data[section]
	if val != nil {
		if v, ok := val.(map[string]interface{}); ok {
			for k, a := range v {
				data[k] = j.TranslateAll(fmt.Sprintf("%s", a), false)
			}
		}
	}
	return data, nil
}

//GetSectionString 获取section原始JSON串
func (j *JSONConf) GetSectionString(section string) (r string, err error) {
	//nkey := "_section_string_" + section
	//if value, ok := j.cache.Get(nkey); ok {
	//return value.(string), nil
	//}
	val := j.data[section]
	if val != nil {
		buffer, err := json.Marshal(val)
		if err != nil {
			return "", err
		}
		r = j.TranslateAll(jsons.Escape(string(buffer)), false)
		//	j.cache.Set(nkey, r)
		return r, nil
	}
	err = errors.New("not exist section:" + section)
	return
}
