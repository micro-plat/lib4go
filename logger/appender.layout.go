package logger

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/micro-plat/lib4go/file"
)

const loggerPath = "../conf/logger.json"

//Layout 输出器
type Layout struct {
	Type   string `json:"type" toml:"type"`
	Level  string `json:"level" toml:"level"`
	Path   string `json:"path,omitempty" toml:"path"`
	Layout string `json:"layout" toml:"layout"`
}

func newDefLayouts() []*Layout {
	layouts := make([]*Layout, 0, 2)
	fileLayout := &Layout{Type: "file", Level: SLevel_ALL}
	fileLayout.Path, _ = file.GetAbs("../logs/%date.log")
	fileLayout.Layout = "[%datetime.%ms][%l][%session] %content%n"
	layouts = append(layouts, fileLayout)

	stdLayout := &Layout{Type: "stdout", Level: SLevel_ALL}
	stdLayout.Layout = "[%datetime.%ms][%l][%session]%content"
	layouts = append(layouts, stdLayout)
	return layouts
}

//Encode 将当前配置内容保存到文件中
func Encode(path string) error {
	if _, err := os.Stat(path); err == nil || os.IsExist(err) {
		return nil
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("无法打开文件:%s %w", path, err)
	}
	encoder := toml.NewEncoder(f)
	err = encoder.Encode(newDefLayouts())
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

//Decode 从配置文件中读取配置信息
func Decode(f string) ([]*Layout, error) {
	data := make([]*Layout, 0, 1)
	if _, err := toml.DecodeFile(f, &data); err != nil {
		return nil, err
	}
	return data, nil
}

//进行日志配置文件初始化
func init() {
	if err := Encode(loggerPath); err != nil {
		sysLog.Errorf("创建日志配置文件失败 %v", err)
	}
	layouts, err := Decode(loggerPath)
	if err != nil {
		sysLog.Errorf("读取配置文件失败 %v", err)
	}
	AddLayout(layouts...)
}
