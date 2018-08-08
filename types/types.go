package types

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

//GetString 从对象中获取数据值，如果不是字符串则返回空
func GetString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

//GetInt 从对象中获取数据值，如果不是字符串则返回0
func GetInt(v interface{}, def ...int) int {
	if value, err := strconv.Atoi(fmt.Sprintf("%v", v)); err == nil {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

//GetInt64 从对象中获取数据值，如果不是字符串则返回0
func GetInt64(v interface{}, def ...int64) int64 {
	if value, err := strconv.ParseInt(fmt.Sprintf("%v", v), 10, 64); err == nil {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

//GetFloat32 从对象中获取数据值，如果不是字符串则返回0
func GetFloat32(v interface{}, def ...float32) float32 {
	if value, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 32); err == nil {
		return float32(value)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

//GetFloat64 从对象中获取数据值，如果不是字符串则返回0
func GetFloat64(v interface{}, def ...float64) float64 {
	if value, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64); err == nil {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

//GetBool 从对象中获取bool类型值，表示为true的值有：1, t, T, true, TRUE, True, YES, yes, Yes, Y, y, ON, on, On
func GetBool(v interface{}, def ...bool) bool {
	if value, err := parseBool(v); err == nil {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return false
}

//GetDatatime 获取时间字段
func GetDatatime(v interface{}, format ...string) (time.Time, error) {
	t, b := MustString(v)
	if !b {
		return time.Now(), errors.New("值不能为空")
	}
	f := "2006/01/02 15:04:05"
	if len(format) > 0 {
		f = format[0]
	}
	return time.ParseInLocation(f, t, time.Local)
}

//MustString 从对象中获取数据值，如果不是字符串则返回空
func MustString(v interface{}) (string, bool) {
	if value, ok := v.(string); ok {
		return value, true
	}
	return "", false
}

//MustInt 从对象中获取数据值，如果不是字符串则返回0
func MustInt(v interface{}) (int, bool) {
	if value, err := strconv.Atoi(fmt.Sprintf("%v", v)); err == nil {
		return value, true
	}
	return 0, false
}

//MustFloat32 从对象中获取数据值，如果不是字符串则返回0
func MustFloat32(v interface{}) (float32, bool) {
	if value, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 32); err == nil {
		return float32(value), true
	}
	return 0, false
}

//MustFloat64 从对象中获取数据值，如果不是字符串则返回0
func MustFloat64(v interface{}) (float64, bool) {
	if value, err := strconv.ParseFloat(fmt.Sprintf("%v", v), 64); err == nil {
		return value, true
	}
	return 0, false
}

//IsEmpty 当前对像是否是字符串空
func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	if t, ok := v.(string); ok && len(t) == 0 {
		return true
	}
	if t, ok := v.([]interface{}); ok && len(t) == 0 {
		return true
	}
	return false
}

//IntContains int数组中是否包含指定值
func IntContains(input []int, v int) bool {
	for _, i := range input {
		if i == v {
			return true
		}
	}
	return false
}
