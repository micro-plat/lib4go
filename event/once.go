package event

import "sync"
import "fmt"

//Once 一次执行锁,相同事件多次完成时只释放一次锁
type Once struct {
	sync    *sync.WaitGroup
	lk      *sync.Mutex
	actions map[string]int
}

//NewOnce 构建一次执行(防止相同的事件多次进入)锁
func NewOnce(wait int) (sn Once, err error) {
	/*add by champly 2016年12月08日14:17:57*/
	if wait <= 0 {
		err = fmt.Errorf("wait must greater than zero")
		return
	}
	/*end*/

	sn = Once{}
	sn.lk = &sync.Mutex{}
	sn.sync = &sync.WaitGroup{}
	sn.actions = make(map[string]int)
	sn.sync.Add(wait)
	return
}

//Wait 等待所有执行完成
func (s Once) Wait() {
	s.sync.Wait()
}

//WaitAndAdd 等待所有执行完成并添加新的任务
func (s Once) WaitAndAdd(delta int) {
	s.sync.Wait()
	s.sync.Add(delta)
}

//AddStep 添加执行步骤
func (s Once) AddStep(delta int) {
	s.sync.Add(delta)
}

//Done 完成任务
func (s Once) Done(action string) {
	s.lk.Lock()
	defer s.lk.Unlock()
	if _, ok := s.actions[action]; !ok {
		s.actions[action] = 1
		s.sync.Done()
	}
}

//Close 关闭当前锁
func (s Once) Close() {
	for i := range s.actions {
		delete(s.actions, i)
	}
}
