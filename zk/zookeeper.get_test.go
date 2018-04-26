package zk

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/qxnw/lib4go/utility"
)

// TestGetValue 获取节点值检查是否符合预期
func TestGetValueVersion(t *testing.T) {
	zk, err := New(masterAddress, time.Second)
	expect(t, err, nil)
	err = zk.Connect()
	expect(t, err, nil)

	//创建临时节点
	path1 := fmt.Sprintf("/S%s_", utility.GetGUID())
	rpath, err := zk.CreateSeqNode(path1, "data")
	expect(t, err, nil)
	//获取节点值，检查是否正确
	_, version, err := zk.GetValue(rpath)
	expect(t, err, nil)
	expect(t, version, int32(0))

	zk.Update(rpath, "abceee")
	_, version, err = zk.GetValue(rpath)
	expect(t, err, nil)
	expect(t, version, int32(1))
}

// TestGetValue 获取节点值检查是否符合预期
func TestGetValue(t *testing.T) {
	zk, err := New(masterAddress, time.Second)
	expect(t, err, nil)
	err = zk.Connect()
	expect(t, err, nil)

	//创建临时节点
	path1 := fmt.Sprintf("/S%s_", utility.GetGUID())
	rpath, err := zk.CreateSeqNode(path1, "data")
	expect(t, err, nil)
	expect(t, strings.HasPrefix(rpath, path1), true)
	//获取节点值，检查是否正确
	buf, _, err := zk.GetValue(rpath)
	expect(t, err, nil)
	expect(t, string(buf), "data")

	//更新节点值
	err = zk.Update(rpath, "data2")
	expect(t, err, nil)

	//重新获取检查值是否变化
	buf, _, err = zk.GetValue(rpath)
	expect(t, err, nil)
	expect(t, string(buf), "data2")

	//删除节点值
	err = zk.Delete(rpath)
	expect(t, err, nil)

	//节点值是否存在
	b, err := zk.Exists(rpath)
	expect(t, err, nil)
	expect(t, b, false)
}
