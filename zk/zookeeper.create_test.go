package zk

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"strings"

	"github.com/qxnw/lib4go/utility"
)

/*
   192.168.0.159:2181
   192.168.0.154:2181
   做的集群，159为主
*/

var (
	masterAddress = []string{"192.168.0.159:2181", "192.168.0.154:2181"}
)

// TestCreatePersistentNode 测试监控一个节点的值是否发送变化
func TestCreatePersistentNode(t *testing.T) {
	zk, err := New(masterAddress, time.Second)
	expect(t, err, nil)
	err = zk.Connect()
	expect(t, err, nil)

	//创建永久节点
	path1 := fmt.Sprintf("/%s", utility.GetGUID())
	err = zk.CreatePersistentNode(path1, "data")
	expect(t, err, nil)

	//检查节点是否存在
	b, err := zk.Exists(path1)
	expect(t, err, nil)
	expect(t, b, true)
	zk.Close()

	//关闭连接
	b, err = zk.Exists(path1)
	expect(t, b, false)
	expect(t, err, ErrColientCouldNotConnect)

	//重新连接
	err = zk.Reconnect()
	expect(t, err, nil)

	//节点应该存在
	b, err = zk.Exists(path1)
	expect(t, err, nil)
	expect(t, b, true)

	//删除节点
	err = zk.Delete(path1)
	expect(t, err, nil)
}

// TestBindWatchValue 测试监控一个节点的值是否发送变化
func TestBindCreateTempNode(t *testing.T) {
	zk, err := New(masterAddress, time.Second)
	expect(t, err, nil)
	err = zk.Connect()
	expect(t, err, nil)

	//创建临时节点
	path1 := fmt.Sprintf("/T%s", utility.GetGUID())
	err = zk.CreateTempNode(path1, "data")
	expect(t, err, nil)

	//检查节点是否存在
	b, err := zk.Exists(path1)
	expect(t, err, nil)
	expect(t, b, true)
	zk.Close()

	//关闭连接
	b, err = zk.Exists(path1)
	expect(t, b, false)
	expect(t, err, ErrColientCouldNotConnect)

	//重新连接
	err = zk.Reconnect()
	expect(t, err, nil)

	//节点应该存在
	b, err = zk.Exists(path1)
	expect(t, err, nil)
	expect(t, b, false)
}

// TestBindWatchValue 测试监控一个节点的值是否发送变化
func TestBindCreateSeqNode(t *testing.T) {
	zk, err := New(masterAddress, time.Second)
	expect(t, err, nil)
	err = zk.Connect()
	expect(t, err, nil)

	//创建临时节点
	path1 := fmt.Sprintf("/S%s_", utility.GetGUID())
	rpath, err := zk.CreateSeqNode(path1, "data")
	expect(t, err, nil)
	expect(t, strings.HasPrefix(rpath, path1), true)
	//检查节点是否存在
	b, err := zk.Exists(rpath)
	expect(t, err, nil)
	expect(t, b, true)
	zk.Close()

	//关闭连接
	b, err = zk.Exists(rpath)
	expect(t, b, false)
	expect(t, err, ErrColientCouldNotConnect)

	//重新连接
	err = zk.Reconnect()
	expect(t, err, nil)

	//节点应该存在
	b, err = zk.Exists(rpath)
	expect(t, err, nil)
	expect(t, b, false)
}

// TestBindWatchValue 测试监控一个节点的值是否发送变化
func TestDelete(t *testing.T) {
	zk, err := New(masterAddress, time.Second)
	expect(t, err, nil)
	err = zk.Connect()
	expect(t, err, nil)

	//创建临时节点
	path1 := fmt.Sprintf("/S%s_", utility.GetGUID())
	rpath, err := zk.CreateSeqNode(path1, "data")
	expect(t, err, nil)
	expect(t, strings.HasPrefix(rpath, path1), true)
	//检查节点是否存在
	b, err := zk.Exists(rpath)
	expect(t, err, nil)
	expect(t, b, true)

	//删除节点
	err = zk.Delete(rpath)
	expect(t, err, nil)

	//检查是否存在
	b, err = zk.Exists(rpath)
	expect(t, err, nil)
	expect(t, b, false)
}
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
