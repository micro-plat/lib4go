package types

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type Input struct {
	Name  string `m2s:"name"`
	Value int    `m2s:"value"`
}

func TestT(t *testing.T) {
	i := NewXMapByMap(
		map[string]interface{}{
			"name":  "abcef",
			"value": 100,
		})

	var input Input
	err := i.ToStruct(&input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", input)

	// t.Error("abc")

}
func TestAppend(t *testing.T) {
	i := NewXMaps()
	i.Append(map[string]interface{}{
		"name":  "abcef",
		"value": 100,
	})

	var input []*Input
	err := i.ToStructs(&input)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("1. %+v\n", i)
	for _, v := range input {
		fmt.Printf("2.%+v\n", v)
	}

	// t.Error("abc")

}

func TestXMap_Keys(t *testing.T) {
	tests := []struct {
		name string
		q    XMap
		want map[string]string
	}{
		{name: "XMap存在单个数据", q: XMap{"key1": "1"}, want: map[string]string{"key1": "key1"}},
		{name: "XMap存在多个数据", q: XMap{"key1": "1", "key2": "2", "key3": "3"}, want: map[string]string{"key1": "key1", "key2": "key2", "key3": "key3"}},
		{name: "XMap不存在数据", q: XMap{}, want: map[string]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.q.Keys()
			if len(got) != len(tt.want) {
				t.Errorf("XMap.Keys() = %v, want %v", got, tt.want)
			}
			for _, item := range got {
				if _, ok := tt.want[item]; !ok {
					t.Errorf("XMap.Keys() = %v, want %v", got, tt.want)
				} else {
					delete(tt.want, item)
				}
			}
			if len(tt.want) > 0 {
				t.Errorf("XMap.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMap_GetString(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		q    XMap
		args args
		want string
	}{
		{name: "对象没有数据", q: XMap{}, args: args{name: "tkey"}, want: ""},
		{name: "数据不存在", q: XMap{"key1": "1"}, args: args{name: "tkey"}, want: ""},
		{name: "数据存在,类型不正确int", q: XMap{"key1": 1}, args: args{name: "key1"}, want: "1"},
		{name: "数据存在,类型不正确float", q: XMap{"key1": float32(10.1)}, args: args{name: "key1"}, want: "10.1"},
		//@todo 用例再讨论
		//{name: "数据存在,类型不正确float大数", q: XMap{"key1": float32(10012212425742542454.1)}, args: args{name: "key1"}, want: "10012212425742542454.1"},
		{name: "数据存在,类型不正确nil", q: XMap{"key1": nil}, args: args{name: "key1"}, want: ""},
		{name: "数据存在,类型不正确负数", q: XMap{"key1": -100}, args: args{name: "key1"}, want: "-100"},
		{name: "数据存在,类型正确", q: XMap{"key1": "1"}, args: args{name: "key1"}, want: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.GetString(tt.args.name); got != tt.want {
				t.Errorf("XMap.GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMap_GetInt(t *testing.T) {
	type args struct {
		name string
		def  []int
	}
	tests := []struct {
		name string
		q    XMap
		args args
		want int
	}{
		{name: "对象为空,无默认", q: XMap{}, args: args{name: "xx", def: []int{}}, want: 0},
		{name: "对象为空,有默认", q: XMap{}, args: args{name: "xx", def: []int{1}}, want: 1},
		{name: "数据不存在,无默认", q: XMap{"yy": 12}, args: args{name: "xx", def: []int{}}, want: 0},
		{name: "数据不存在,有默认", q: XMap{"yy": 12}, args: args{name: "xx", def: []int{1}}, want: 1},
		{name: "数据存在,类型是string字符,无默认", q: XMap{"yy": "as"}, args: args{name: "yy", def: []int{}}, want: 0},
		{name: "数据存在,类型是string字符,有默认", q: XMap{"yy": "as"}, args: args{name: "yy", def: []int{1}}, want: 1},
		{name: "数据存在,类型是string数字,无默认", q: XMap{"yy": "12"}, args: args{name: "yy", def: []int{}}, want: 12},
		{name: "数据存在,类型是string数字,有默认", q: XMap{"yy": "12"}, args: args{name: "yy", def: []int{1}}, want: 12},
		{name: "数据存在,类型是string大数字,无默认", q: XMap{"yy": "1212222222222222222222222222222222"}, args: args{name: "yy", def: []int{}}, want: 0},
		{name: "数据存在,类型是string大数字,有默认", q: XMap{"yy": "1212222222222222222222222222222222"}, args: args{name: "yy", def: []int{1}}, want: 1},

		{name: "数据存在,类型是float整数,无默认", q: XMap{"yy": float32(12)}, args: args{name: "yy", def: []int{}}, want: 12},
		{name: "数据存在,类型是float整数,有默认", q: XMap{"yy": float32(12)}, args: args{name: "yy", def: []int{1}}, want: 12},
		{name: "数据存在,类型是float小数,无默认", q: XMap{"yy": float32(12.1)}, args: args{name: "yy", def: []int{}}, want: 0},
		{name: "数据存在,类型是float小数,有默认", q: XMap{"yy": float32(12.1)}, args: args{name: "yy", def: []int{1}}, want: 1},
		//@todo 待讨论
		//{name: "数据存在,类型是float大数,无默认", q: XMap{"yy": float64(1212222222222222222222222222222222)}, args: args{name: "yy", def: []int{}}, want: 0},
		//{name: "数据存在,类型是float大数,有默认", q: XMap{"yy": float64(1212222222222222222222222222222222)}, args: args{name: "yy", def: []int{1}}, want: 1},
		{name: "数据存在,类型是int,无默认", q: XMap{"yy": 12}, args: args{name: "yy", def: []int{}}, want: 12},
		{name: "数据存在,类型是int,有默认", q: XMap{"yy": 12}, args: args{name: "yy", def: []int{1}}, want: 12},
		{name: "数据存在,类型是int大数,无默认", q: XMap{"yy": 6666666666666666666}, args: args{name: "yy", def: []int{}}, want: 6666666666666666666},
		{name: "数据存在,类型是int大数,有默认", q: XMap{"yy": 6666666666666666666}, args: args{name: "yy", def: []int{1}}, want: 6666666666666666666},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.GetInt(tt.args.name, tt.args.def...); got != tt.want {
				t.Errorf("XMap.GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMap_GetInt64(t *testing.T) {
	type args struct {
		name string
		def  []int64
	}
	tests := []struct {
		name string
		q    XMap
		args args
		want int64
	}{
		{name: "对象为空,无默认", q: XMap{}, args: args{name: "xx", def: []int64{}}, want: 0},
		{name: "对象为空,有默认", q: XMap{}, args: args{name: "xx", def: []int64{1}}, want: 1},
		{name: "数据不存在,无默认", q: XMap{"yy": 12}, args: args{name: "xx", def: []int64{}}, want: 0},
		{name: "数据不存在,有默认", q: XMap{"yy": 12}, args: args{name: "xx", def: []int64{1}}, want: 1},
		{name: "数据存在,类型是string字符,无默认", q: XMap{"yy": "as"}, args: args{name: "yy", def: []int64{}}, want: 0},
		{name: "数据存在,类型是string字符,有默认", q: XMap{"yy": "as"}, args: args{name: "yy", def: []int64{1}}, want: 1},
		{name: "数据存在,类型是string数字,无默认", q: XMap{"yy": "12"}, args: args{name: "yy", def: []int64{}}, want: 12},
		{name: "数据存在,类型是string数字,有默认", q: XMap{"yy": "12"}, args: args{name: "yy", def: []int64{1}}, want: 12},
		{name: "数据存在,类型是string大数字,无默认", q: XMap{"yy": "1212222222222222222222222222222222"}, args: args{name: "yy", def: []int64{}}, want: 0},
		{name: "数据存在,类型是string大数字,有默认", q: XMap{"yy": "1212222222222222222222222222222222"}, args: args{name: "yy", def: []int64{1}}, want: 1},
		{name: "数据存在,类型是float整数,无默认", q: XMap{"yy": float32(12)}, args: args{name: "yy", def: []int64{}}, want: 12},
		{name: "数据存在,类型是float整数,有默认", q: XMap{"yy": float32(12)}, args: args{name: "yy", def: []int64{1}}, want: 12},
		{name: "数据存在,类型是float小数,无默认", q: XMap{"yy": float32(12.1)}, args: args{name: "yy", def: []int64{}}, want: 0},
		{name: "数据存在,类型是float小数,有默认", q: XMap{"yy": float32(12.1)}, args: args{name: "yy", def: []int64{1}}, want: 1},
		//@todo 待讨论
		//{name: "数据存在,类型是float大数,无默认", q: XMap{"yy": float64(1212222222222222222222222222222222)}, args: args{name: "yy", def: []int64{}}, want: 0},
		//{name: "数据存在,类型是float大数,有默认", q: XMap{"yy": float64(1212222222222222222222222222222222)}, args: args{name: "yy", def: []int64{1}}, want: 1},
		{name: "数据存在,类型是int,无默认", q: XMap{"yy": 12}, args: args{name: "yy", def: []int64{}}, want: 12},
		{name: "数据存在,类型是int,有默认", q: XMap{"yy": 12}, args: args{name: "yy", def: []int64{1}}, want: 12},
		{name: "数据存在,类型是int大数,无默认", q: XMap{"yy": 6666666666666666666}, args: args{name: "yy", def: []int64{}}, want: 6666666666666666666},
		{name: "数据存在,类型是int大数,有默认", q: XMap{"yy": 6666666666666666666}, args: args{name: "yy", def: []int64{1}}, want: 6666666666666666666},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.GetInt64(tt.args.name, tt.args.def...); got != tt.want {
				t.Errorf("XMap.GetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMap_GetFloat32(t *testing.T) {
	type args struct {
		name string
		def  []float32
	}
	tests := []struct {
		name string
		q    XMap
		args args
		want float32
	}{
		{name: "对象为空,无默认", q: XMap{}, args: args{name: "xx", def: []float32{}}, want: 0},
		{name: "对象为空,有默认", q: XMap{}, args: args{name: "xx", def: []float32{1.1}}, want: 1.1},
		{name: "数据不存在,无默认", q: XMap{"yy": 12.1}, args: args{name: "xx", def: []float32{}}, want: 0},
		{name: "数据不存在,有默认", q: XMap{"yy": 12.1}, args: args{name: "xx", def: []float32{1.1}}, want: 1.1},
		{name: "数据存在,类型是string字符,无默认", q: XMap{"yy": "as"}, args: args{name: "yy", def: []float32{}}, want: 0},
		{name: "数据存在,类型是string字符,有默认", q: XMap{"yy": "as"}, args: args{name: "yy", def: []float32{1.1}}, want: 1.1},
		{name: "数据存在,类型是string数字,无默认", q: XMap{"yy": "12.1"}, args: args{name: "yy", def: []float32{}}, want: 12.1},
		{name: "数据存在,类型是string数字,有默认", q: XMap{"yy": "12.1"}, args: args{name: "yy", def: []float32{1}}, want: 12.1},
		{name: "数据存在,类型是string大数字,无默认", q: XMap{"yy": "12122222222222222222222222222.22222"}, args: args{name: "yy", def: []float32{}}, want: 12122222222222222222222222222.22222},
		{name: "数据存在,类型是string大数字,有默认", q: XMap{"yy": "121222222222222222222222222222.2222"}, args: args{name: "yy", def: []float32{1.1}}, want: 121222222222222222222222222222.2222},
		{name: "数据存在,类型是float整数,无默认", q: XMap{"yy": float32(12)}, args: args{name: "yy", def: []float32{}}, want: 12},
		{name: "数据存在,类型是float整数,有默认", q: XMap{"yy": float32(12)}, args: args{name: "yy", def: []float32{1}}, want: 12},
		{name: "数据存在,类型是float小数,无默认", q: XMap{"yy": float32(12.1)}, args: args{name: "yy", def: []float32{}}, want: 12.1},
		{name: "数据存在,类型是float小数,有默认", q: XMap{"yy": float32(12.1)}, args: args{name: "yy", def: []float32{1}}, want: 12.1},
		{name: "数据存在,类型是float大数,无默认", q: XMap{"yy": float64(1212222222222222222222222222222222222222222222222222222)}, args: args{name: "yy", def: []float32{}}, want: 0},
		{name: "数据存在,类型是float大数,有默认", q: XMap{"yy": float64(1212222222222222222222222222222222222222222222222222222)}, args: args{name: "yy", def: []float32{1}}, want: 1},
		{name: "数据存在,类型是int,无默认", q: XMap{"yy": 12}, args: args{name: "yy", def: []float32{}}, want: 12},
		{name: "数据存在,类型是int,有默认", q: XMap{"yy": 12}, args: args{name: "yy", def: []float32{1}}, want: 12},
		{name: "数据存在,类型是int大数,无默认", q: XMap{"yy": 6666666666666666666}, args: args{name: "yy", def: []float32{}}, want: 6666666666666666666},
		{name: "数据存在,类型是int大数,有默认", q: XMap{"yy": 6666666666666666666}, args: args{name: "yy", def: []float32{1}}, want: 6666666666666666666},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.GetFloat32(tt.args.name, tt.args.def...); got != tt.want {
				t.Errorf("XMap.GetFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMap_GetFloat64(t *testing.T) {
	type args struct {
		name string
		def  []float64
	}
	tests := []struct {
		name string
		q    XMap
		args args
		want float64
	}{
		{name: "对象为空,无默认", q: XMap{}, args: args{name: "xx", def: []float64{}}, want: 0},
		{name: "对象为空,有默认", q: XMap{}, args: args{name: "xx", def: []float64{1.1}}, want: 1.1},
		{name: "数据不存在,无默认", q: XMap{"yy": 12.1}, args: args{name: "xx", def: []float64{}}, want: 0},
		{name: "数据不存在,有默认", q: XMap{"yy": 12.1}, args: args{name: "xx", def: []float64{1.1}}, want: 1.1},
		{name: "数据存在,类型是string字符,无默认", q: XMap{"yy": "as"}, args: args{name: "yy", def: []float64{}}, want: 0},
		{name: "数据存在,类型是string字符,有默认", q: XMap{"yy": "as"}, args: args{name: "yy", def: []float64{1.1}}, want: 1.1},
		{name: "数据存在,类型是string数字,无默认", q: XMap{"yy": "12.1"}, args: args{name: "yy", def: []float64{}}, want: 12.1},
		{name: "数据存在,类型是string数字,有默认", q: XMap{"yy": "12.1"}, args: args{name: "yy", def: []float64{1}}, want: 12.1},
		{name: "数据存在,类型是string大数字,无默认", q: XMap{"yy": "999999999999999999999999999999999999999999999999999999999999999999912222222222222222223323232323333333333333333333333333333333333333333344444444444444444444444444444444444444444444444444444444444444444444444444443222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222.22222"}, args: args{name: "yy", def: []float64{}}, want: 0},
		{name: "数据存在,类型是string大数字,有默认", q: XMap{"yy": "9999999999999999999999999999999999999999999999999999999999999999999122222222222222222233232323233333333333333333333333333333333333333333444444444444444444444444444444444444444444444444444444444444444444444444444432222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222.2222"}, args: args{name: "yy", def: []float64{1.1}}, want: 1.1},
		{name: "数据存在,类型是float整数,无默认", q: XMap{"yy": float64(12)}, args: args{name: "yy", def: []float64{}}, want: 12},
		{name: "数据存在,类型是float整数,有默认", q: XMap{"yy": float64(12)}, args: args{name: "yy", def: []float64{1}}, want: 12},
		{name: "数据存在,类型是float小数,无默认", q: XMap{"yy": float64(12.1)}, args: args{name: "yy", def: []float64{}}, want: 12.1},
		{name: "数据存在,类型是float小数,有默认", q: XMap{"yy": float64(12.1)}, args: args{name: "yy", def: []float64{1}}, want: 12.1},
		{name: "数据存在,类型是float大数,无默认", q: XMap{"yy": float64(1212222222222222222222222222222222)}, args: args{name: "yy", def: []float64{}}, want: 1212222222222222222222222222222222},
		{name: "数据存在,类型是float大数,有默认", q: XMap{"yy": float64(1212222222222222222222222222222222)}, args: args{name: "yy", def: []float64{1}}, want: 1212222222222222222222222222222222},
		{name: "数据存在,类型是int,无默认", q: XMap{"yy": 12}, args: args{name: "yy", def: []float64{}}, want: 12},
		{name: "数据存在,类型是int,有默认", q: XMap{"yy": 12}, args: args{name: "yy", def: []float64{1}}, want: 12},
		{name: "数据存在,类型是int大数,无默认", q: XMap{"yy": 6666666666666666666}, args: args{name: "yy", def: []float64{}}, want: 6666666666666666666},
		{name: "数据存在,类型是int大数,有默认", q: XMap{"yy": 6666666666666666666}, args: args{name: "yy", def: []float64{1}}, want: 6666666666666666666},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.GetFloat64(tt.args.name, tt.args.def...); got != tt.want {
				t.Errorf("XMap.GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMap_GetBool(t *testing.T) {
	type args struct {
		name string
		def  []bool
	}
	tests := []struct {
		name string
		q    XMap
		args args
		want bool
	}{
		{name: "对象为空,无默认", q: XMap{}, args: args{name: "xx", def: []bool{}}, want: false},
		{name: "对象为空,有默认", q: XMap{}, args: args{name: "xx", def: []bool{true}}, want: true},
		{name: "数据不存在,无默认", q: XMap{"yy": true}, args: args{name: "xx", def: []bool{}}, want: false},
		{name: "数据不存在,有默认", q: XMap{"yy": false}, args: args{name: "xx", def: []bool{true}}, want: true},
		{name: "数据存在,值为string-1", q: XMap{"yy": "1"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为int-1", q: XMap{"yy": 1}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为float-1", q: XMap{"yy": float32(1)}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为float-1.0", q: XMap{"yy": float32(1.0)}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为float64-1", q: XMap{"yy": float64(1)}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为float64-1.0", q: XMap{"yy": float64(1.0)}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为rune-1", q: XMap{"yy": rune(1)}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为byte-1", q: XMap{"yy": byte(1)}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,值为string-t", q: XMap{"yy": "t"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为rune-t", q: XMap{"yy": rune('t')}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为byte-t", q: XMap{"yy": byte('t')}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,值为string-T", q: XMap{"yy": "T"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为rune-T", q: XMap{"yy": rune('T')}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为byte-T", q: XMap{"yy": byte('T')}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,值为string-Y", q: XMap{"yy": "Y"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为rune-Y", q: XMap{"yy": rune('Y')}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为byte-Y", q: XMap{"yy": byte('Y')}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,值为string-y", q: XMap{"yy": "y"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为rune-y", q: XMap{"yy": rune('y')}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为byte-y", q: XMap{"yy": byte('y')}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,值为string-true", q: XMap{"yy": "true"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-True", q: XMap{"yy": "True"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-TRUE", q: XMap{"yy": "TRUE"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-tRue", q: XMap{"yy": "tRue"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-yes", q: XMap{"yy": "yes"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-Yes", q: XMap{"yy": "Yes"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-YES", q: XMap{"yy": "YES"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-yEs", q: XMap{"yy": "yEs"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-on", q: XMap{"yy": "on"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-On", q: XMap{"yy": "On"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-ON", q: XMap{"yy": "ON"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,值为string-oN", q: XMap{"yy": "oN"}, args: args{name: "yy", def: []bool{}}, want: true},
		{name: "数据存在,false-f", q: XMap{"yy": "f"}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,false-yy", q: XMap{"yy": "yy"}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,false-tture", q: XMap{"yy": "tture"}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,false-1.0", q: XMap{"yy": "1.0"}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,false-yess", q: XMap{"yy": "yess"}, args: args{name: "yy", def: []bool{}}, want: false},
		{name: "数据存在,false-oon", q: XMap{"yy": "oon"}, args: args{name: "yy", def: []bool{}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.GetBool(tt.args.name, tt.args.def...); got != tt.want {
				t.Errorf("XMap.GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMap_GetDatetime(t *testing.T) {
	yesTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-01-02 15:16:17", time.Local)
	type args struct {
		name   string
		format []string
	}
	tests := []struct {
		name    string
		q       XMap
		args    args
		want    time.Time
		wantErr bool
	}{
		{name: "对象为空", q: XMap{}, args: args{name: "xx", format: []string{}}, want: time.Time{}, wantErr: true},
		{name: "数据不存在", q: XMap{"yy": "2021-01-02 15:16:17"}, args: args{name: "xx", format: []string{}}, want: time.Time{}, wantErr: true},

		{name: "正确数据存在-,默认格式化", q: XMap{"yy": "2021-01-02 15:16:17"}, args: args{name: "yy", format: []string{}}, want: time.Time{}, wantErr: true},
		{name: "正确数据存在-,自定义正确格式化", q: XMap{"yy": "2021-01-02 15:16:17"}, args: args{name: "yy", format: []string{"2006-01-02 15:04:05"}}, want: yesTime, wantErr: false},
		{name: "正确数据存在-,自定义错误格式化", q: XMap{"yy": "2021-01-02 15:16:17"}, args: args{name: "yy", format: []string{"2006-01-02 15:04:07"}}, want: time.Time{}, wantErr: true},

		{name: "错误数据存在-,默认格式化", q: XMap{"yy": "2021-13-02 15:16:17"}, args: args{name: "yy", format: []string{}}, want: time.Time{}, wantErr: true},
		{name: "错误数据存在-,自定义正确格式化", q: XMap{"yy": "2021-13-02 15:16:17"}, args: args{name: "yy", format: []string{"2006-01-02 15:04:05"}}, want: time.Time{}, wantErr: true},
		{name: "错误数据存在-,自定义错误格式化", q: XMap{"yy": "2021-13-02 15:16:17"}, args: args{name: "yy", format: []string{"2006-01-02 15:04:07"}}, want: time.Time{}, wantErr: true},

		{name: "正确数据存在/,默认格式化", q: XMap{"yy": "2021/01/02 15:16:17"}, args: args{name: "yy", format: []string{}}, want: yesTime, wantErr: false},
		{name: "错误数据存在/", q: XMap{"yy": "2021/13/02 15:16:17"}, args: args{name: "yy", format: []string{}}, want: time.Time{}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.GetDatetime(tt.args.name, tt.args.format...)
			if (err != nil) != tt.wantErr {
				t.Errorf("XMap.GetDatetime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMap.GetDatetime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMap_Set(t *testing.T) {
	type args struct {
		name  string
		value interface{}
	}
	tests := []struct {
		name string
		q    XMap
		args args
	}{
		{name: "新增数据", q: XMap{}, args: args{name: "test1", value: "1"}},
		{name: "更新数据,相同类型", q: XMap{"test1": "2"}, args: args{name: "test1", value: "1"}},
		{name: "更新数据,不相同类型,相同值", q: XMap{"test1": 1}, args: args{name: "test1", value: "1"}},
		{name: "更新数据,不相同类型,不同值", q: XMap{"test1": 2}, args: args{name: "test1", value: "1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.SetValue(tt.args.name, tt.args.value)
			if !reflect.DeepEqual(tt.q[tt.args.name], tt.args.value) {
				t.Errorf("XMap.Set() fail")
			}
		})
	}
}

func TestXMap_Has(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		q    XMap
		args args
		want bool
	}{
		{name: "空对象,数据不存在", q: XMap{}, args: args{name: "test1"}, want: false},
		{name: "数据不存在", q: XMap{"test2": "2"}, args: args{name: "test1"}, want: false},
		{name: "数据存在,值为空", q: XMap{"test1": nil}, args: args{name: "test1"}, want: true},  //?
		{name: "数据存在,值为字符空", q: XMap{"test1": ""}, args: args{name: "test1"}, want: true}, //?
		{name: "数据存在,值不为空", q: XMap{"test1": 2}, args: args{name: "test1"}, want: true}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.Has(tt.args.name); got != tt.want {
				t.Errorf("XMap.Has() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestXMap_MustString(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		q     XMap
		args  args
		want  string
		want1 bool
	}{
		{name: "空对象,数据不存在", q: XMap{}, args: args{name: "test1"}, want: "", want1: false},
		{name: "数据不存在", q: XMap{"test2": "2"}, args: args{name: "test1"}, want: "", want1: false},
		{name: "数据存在,值为空", q: XMap{"test1": nil}, args: args{name: "test1"}, want: "", want1: false},
		{name: "数据存在,值为字符空", q: XMap{"test1": ""}, args: args{name: "test1"}, want: "", want1: true},
		{name: "数据存在,值不为空-float", q: XMap{"test1": float32(2.0)}, args: args{name: "test1"}, want: "", want1: false},
		{name: "数据存在,值不为空-int", q: XMap{"test1": 2}, args: args{name: "test1"}, want: "", want1: false},
		{name: "数据存在,值不为空-rune", q: XMap{"test1": rune(2)}, args: args{name: "test1"}, want: "", want1: false},
		{name: "数据存在,值不为空-byte", q: XMap{"test1": byte(1)}, args: args{name: "test1"}, want: "", want1: false},
		{name: "数据存在,值不为空-string", q: XMap{"test1": "123456"}, args: args{name: "test1"}, want: "123456", want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.MustString(tt.args.name)
			if !(got1 == tt.want1 && got == tt.want) {
				t.Errorf("XMap.MustString() got  = %v got1 = %v, want %v,want1%v", got, got1, tt.want, tt.want1)
			}
		})
	}
}

//有强制转换的左右   类型不同只要能转换成功  都是正确的
func TestXMap_MustInt(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		q     XMap
		args  args
		want  int
		want1 bool
	}{
		{name: "空对象,数据不存在", q: XMap{}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据不存在", q: XMap{"test2": "2"}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值为空", q: XMap{"test1": nil}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值为0", q: XMap{"test1": 0}, args: args{name: "test1"}, want: 0, want1: true},
		{name: "数据存在,值不为空负数", q: XMap{"test1": -2}, args: args{name: "test1"}, want: -2, want1: true},
		{name: "数据存在,值不为空正数", q: XMap{"test1": 2}, args: args{name: "test1"}, want: 2, want1: true},

		{name: "数据存在,值不为空-float", q: XMap{"test1": float32(2)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-float.0", q: XMap{"test1": float32(2.0)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-float.1", q: XMap{"test1": float32(2.1)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-rune", q: XMap{"test1": rune(2)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-byte", q: XMap{"test1": byte(1)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-string", q: XMap{"test1": "123456"}, args: args{name: "test1"}, want: 0, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.MustInt(tt.args.name)
			if !(got1 == tt.want1 && got == tt.want) {
				t.Errorf("XMap.MustInt() got  = %v got1 = %v, want %v,want1%v", got, got1, tt.want, tt.want1)
			}
		})
	}
}

//有强制转换的左右   类型不同只要能转换成功  都是正确的 (float返回科学计数法的问题)
func TestXMap_MustFloat32(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		q     XMap
		args  args
		want  float32
		want1 bool
	}{
		{name: "空对象,数据不存在", q: XMap{}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据不存在", q: XMap{"test2": "2"}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值为空", q: XMap{"test1": nil}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值为0", q: XMap{"test1": 0}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空负数", q: XMap{"test1": -2.1}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空正数", q: XMap{"test1": 2.1}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-float", q: XMap{"test1": float32(2)}, args: args{name: "test1"}, want: 2, want1: true},
		{name: "数据存在,值不为空-float.0", q: XMap{"test1": float32(2.0)}, args: args{name: "test1"}, want: 2, want1: true},
		{name: "数据存在,值不为空-float.1", q: XMap{"test1": float32(2.1)}, args: args{name: "test1"}, want: 2.1, want1: true},
		{name: "数据存在,值不为空-float32", q: XMap{"test1": float32(22222222222222222222222222222)}, args: args{name: "test1"}, want: 22222222222222222222222222222, want1: true},
		{name: "数据存在,值不为空-float64,范围内", q: XMap{"test1": float64(22222222222222222222222222222)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-float64", q: XMap{"test1": float64(22222222222222222222222222222222222222222222)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-rune", q: XMap{"test1": rune(2)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-byte", q: XMap{"test1": byte(1)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-string", q: XMap{"test1": "1.32"}, args: args{name: "test1"}, want: 0, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.MustFloat32(tt.args.name)
			if !(got1 == tt.want1 && got == tt.want) {
				t.Errorf("XMap.MustFloat32() got  = %v got1 = %v, want %v,want1%v", got, got1, tt.want, tt.want1)
			}
		})
	}
}

//有强制转换的左右   类型不同只要能转换成功  都是正确的 (32位精度问题)
func TestXMap_MustFloat64(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name  string
		q     XMap
		args  args
		want  float64
		want1 bool
	}{
		{name: "空对象,数据不存在", q: XMap{}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据不存在", q: XMap{"test2": "2"}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值为空", q: XMap{"test1": nil}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值为0", q: XMap{"test1": 0}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空负数", q: XMap{"test1": -2.1}, args: args{name: "test1"}, want: -2.1, want1: true},
		{name: "数据存在,值不为空正数", q: XMap{"test1": 2.1}, args: args{name: "test1"}, want: 2.1, want1: true},
		{name: "数据存在,值不为空-float", q: XMap{"test1": float32(2)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-float.0", q: XMap{"test1": float32(2.0)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-float.1", q: XMap{"test1": float32(2.1)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-float32", q: XMap{"test1": float32(22222222222222222222222222222)}, args: args{name: "test1"}, want: 0, want1: false},
		{name: "数据存在,值不为空-float64", q: XMap{"test1": float64(22222222222222222222222222222222222222222222)}, args: args{name: "test1"}, want: 22222222222222222222222222222222222222222222, want1: true},
		//@todo
		//{name: "数据存在,值不为空-rune", q: XMap{"test1": rune("2")}, args: args{name: "test1"}, want: 2, want1: true},
		//{name: "数据存在,值不为空-byte", q: XMap{"test1": byte("1")}, args: args{name: "test1"}, want: 1, want1: true},
		//{name: "数据存在,值不为空-string", q: XMap{"test1": "1.32"}, args: args{name: "test1"}, want: 1.32, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.q.MustFloat64(tt.args.name)
			if !(got1 == tt.want1 && got == tt.want) {
				t.Errorf("XMap.MustFloat64() got  = %v got1 = %v, want %v,want1%v", got, got1, tt.want, tt.want1)
			}

		})
	}
}

func TestXMap_ToStruct(t *testing.T) {
	type outS struct {
		Intt      int     `json:"intt" m2s:"intt"`
		Float32st float32 `json:"float32t" m2s:"float32t"`
		Stringt   string  `json:"stringt" m2s:"stringt"`
	}

	type outS1 struct {
		intt      int     `json:"intt" m2s:"intt"`
		float32st float32 `json:"float32t" m2s:"float32t"`
		stringt   string  `json:"stringt" m2s:"stringt"`
	}
	type args struct {
		o interface{}
	}
	tests := []struct {
		name    string
		q       XMap
		args    args
		wantErr bool
	}{
		{name: "入参不是指针-string", q: XMap{"intt": 11, "float32t": 1.1, "stringt": "123"}, args: args{o: "sd"}, wantErr: true},
		{name: "入参不是指针-int", q: XMap{}, args: args{o: 123}, wantErr: true},
		{name: "入参不是指针-float", q: XMap{}, args: args{o: 2.01}, wantErr: true},
		{name: "入参不是指针-byte", q: XMap{}, args: args{o: byte(1)}, wantErr: true},
		{name: "入参不是指针-struct", q: XMap{}, args: args{o: outS{}}, wantErr: true},
		{name: "入参不是指针,有数据-struct", q: XMap{"intt": 11, "float32t": 1.1, "stringt": "123"}, args: args{o: outS{}}, wantErr: true},
		{name: "空数据,正确入参,共有变量", q: XMap{"intt": 11, "float32t": 1.1, "stringt": "123"}, args: args{o: &outS{}}, wantErr: false},
		{name: "空数据,正确入参1,私有变量", q: XMap{"intt": 11, "float32t": 1.1, "stringt": "123"}, args: args{o: &outS1{}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.ToStruct(tt.args.o); (err != nil) != tt.wantErr {
				t.Errorf("XMap.ToStruct() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				if v, ok := tt.args.o.(*outS1); ok {
					if v.intt != 0 || v.float32st != 0 || v.stringt != "" {
						t.Errorf("XMap.ToStruct()1  out %v", tt.args.o)
					}
				}

				if v, ok := tt.args.o.(*outS); ok {
					if v.Intt != tt.q.GetInt("intt") || v.Float32st != tt.q.GetFloat32("float32t") || v.Stringt != tt.q.GetString("stringt") {
						t.Errorf("XMap.ToStruct()2 tt.q:%v,  out %v", tt.q, v)
					}
				}
			}
		})
	}
}

func TestXMap_ToSMap(t *testing.T) {
	type st struct {
		Test1 string `json:"test1"`
	}

	tests := []struct {
		name string
		q    XMap
		want map[string]string
	}{
		{name: "空对象转换", q: XMap{}, want: map[string]string{}},
		{name: "float小数转换", q: XMap{"t": 1.1}, want: map[string]string{"t": "1.1"}},
		//@todo
		//{name: "float大数转换", q: XMap{"t": 2222222222222222222222222222222.2}, want: map[string]string{"t": "2222222222222222222222222222222.2"}},
		{name: "int转换", q: XMap{"t": 12}, want: map[string]string{"t": "12"}},
		{name: "int大数转换", q: XMap{"t": 122222222222222222}, want: map[string]string{"t": "122222222222222222"}},
		{name: "rune转换", q: XMap{"t": rune(1)}, want: map[string]string{"t": "1"}},
		{name: "string转换", q: XMap{"t": "12222"}, want: map[string]string{"t": "12222"}},
		{name: "空map转换", q: XMap{"t": map[string]interface{}{}}, want: map[string]string{"t": "{}"}},
		{name: "map-i转换", q: XMap{"t": map[string]interface{}{"11": "11"}}, want: map[string]string{"t": `{"11":"11"}`}},
		{name: "map-s转换", q: XMap{"t": map[string]string{"11": "11"}}, want: map[string]string{"t": `{"11":"11"}`}},
		{name: "空struct转换", q: XMap{"t": st{}}, want: map[string]string{"t": `{"test1":""}`}},
		{name: "struct转换", q: XMap{"t": st{Test1: "11"}}, want: map[string]string{"t": `{"test1":"11"}`}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.q.ToSMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMap.ToSMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMap_MergeMap(t *testing.T) {
	type args struct {
		anr map[string]interface{}
	}
	tests := []struct {
		name string
		q    XMap
		args args
		want map[string]interface{}
	}{
		{name: "空+空对象合并", q: XMap{}, args: args{anr: map[string]interface{}{}}, want: map[string]interface{}{}},
		{name: "空+实体对象合并", q: XMap{}, args: args{anr: map[string]interface{}{"t1": 1}}, want: map[string]interface{}{"t1": 1}},
		{name: "同key不同值合并", q: XMap{"t1": 2}, args: args{anr: map[string]interface{}{"t1": 1}}, want: map[string]interface{}{"t1": 1}},
		{name: "不同key合并", q: XMap{"t1": 2}, args: args{anr: map[string]interface{}{"t2": 1}}, want: map[string]interface{}{"t2": 1, "t1": 2}},
		{name: "多key复合合并", q: XMap{"t1": 2, "t3": "2323"}, args: args{anr: map[string]interface{}{"t2": 1, "t3": 11, "t4": 1.52}}, want: map[string]interface{}{"t1": 2, "t2": 1, "t3": 11, "t4": 1.52}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.q.MergeMap(tt.args.anr)
			if len(tt.q) != len(tt.want) {
				t.Errorf("XMap.MergeMap()1  tt.q= %v, want %v", tt.q, tt.want)
			} else {
				for k, v := range tt.q {
					if _, ok := tt.want[k]; ok && v == tt.want[k] {
						delete(tt.want, k)
					}
				}

				if len(tt.want) > 0 {
					t.Errorf("XMap.MergeMap()2  tt.q= %v, want %v", tt.q, tt.want)
				}
			}
		})
	}
}
