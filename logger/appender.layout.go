package logger

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/micro-plat/lib4go/file"
)

const loggerPath = "../conf/logger.toml"

//Layout 输出器
type Layout struct {
	Type   string `json:"type"  toml:"type"`
	Level  string `json:"level" valid:"in(Off|Info|Warn|Error|Fatal|Debug|All)" toml:"level"`
	Path   string `json:"path,omitempty" toml:"path"`
	Layout string `json:"layout" toml:"layout"`
}
type layoutSetting struct {
	Layouts []*Layout `json:"layouts" toml:"layouts"`
}

func newDefLayouts() *layoutSetting {
	setting := &layoutSetting{Layouts: make([]*Layout, 0, 2)}

	fileLayout := &Layout{Type: "file", Level: SLevel_ALL}
	fileLayout.Path, _ = file.GetAbs("../logs/%date.log")
	fileLayout.Layout = "[%datetime.%ms][%l][%session] %content%n"
	setting.Layouts = append(setting.Layouts, fileLayout)

	stdLayout := &Layout{Type: "stdout", Level: SLevel_ALL}
	stdLayout.Layout = "[%datetime.%ms][%l][%session]%content"
	setting.Layouts = append(setting.Layouts, stdLayout)

	return setting
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
func Decode(f string) (*layoutSetting, error) {
	l := &layoutSetting{}
	if _, err := toml.DecodeFile(f, &l); err != nil {
		return nil, err
	}
	return l, nil
}

//进行日志配置文件初始化
func init() {
	AddAppender("file", NewFileAppender())
	AddAppender("stdout", NewStudoutAppender())

	if err := Encode(loggerPath); err != nil {
		SysLog.Errorf("创建日志配置文件失败 %v", err)
		return
	}
	layouts, err := Decode(loggerPath)
	if err != nil {
		SysLog.Errorf("读取配置文件失败 %v", err)
		return
	}
	AddLayout(layouts.Layouts...)
}
