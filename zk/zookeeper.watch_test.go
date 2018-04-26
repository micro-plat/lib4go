package zk

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/qxnw/lib4go/utility"
)

// TestWatchChildren 获取节点值检查是否符合预期
func TestWatchValueChange(t *testing.T) {

	zk, err := New(masterAddress, time.Second)
	expect(t, err, nil)
	err = zk.Connect()
	expect(t, err, nil)

	//创建临时节点
	path1 := fmt.Sprintf("/t%s", utility.GetGUID())
	err = zk.CreateTempNode(path1, "data")
	expect(t, err, nil)

	//获取节点值，检查是否正确
	buf, _, err := zk.GetValue(path1)
	expect(t, err, nil)
	expect(t, string(buf), "data")

	valueCh, err := zk.WatchValue(path1)
	expect(t, err, nil)

	//更新节点值
	value2 := utility.GetGUID()
	err = zk.Update(path1, value2)
	expect(t, err, nil)

	buf, _, err = zk.GetValue(path1)
	expect(t, err, nil)
	expect(t, string(buf), value2)

	//获取节点值，检查是否正确
	select {
	case <-time.After(TIMEOUT * 2):
		t.Error("获取节点值超时")
	case v := <-valueCh:
		expect(t, v.GetError(), nil)
		value, _ := v.GetValue()
		expect(t, string(value), value2)
	}
	zk.Close()
}

// TestWatchChildren 获取节点值检查是否符合预期
func TestWatchChildren(t *testing.T) {
	zk, err := New(masterAddress, time.Second)
	expect(t, err, nil)
	err = zk.Connect()
	expect(t, err, nil)

	//创建父节点
	root := fmt.Sprintf("/%s", utility.GetGUID())
	err = zk.CreatePersistentNode(root, "data")
	expect(t, err, nil)

	valueCh, err := zk.WatchChildren(root)
	expect(t, err, nil)

	child := utility.GetGUID()
	err = zk.CreateTempNode(fmt.Sprintf("%s/%s", root, child), "")
	expect(t, err, nil)

	//获取节点值，检查是否正确
	select {
	case <-time.After(TIMEOUT * 2):
		t.Error("获取节点值超时")
	case v := <-valueCh:
		expect(t, v.GetError(), nil)
		values, _ := v.GetValue()
		expect(t, strings.Join(values, ","), child)
	}

	zk.Close()

}
