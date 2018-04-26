package event

import "time"
import "fmt"

//DelayCallback 延迟回调，多次触发任务时延迟到一定时间后执行，且只执行一次
type DelayCallback struct {
	msg        chan []interface{}
	delayTime  time.Duration
	firstDelay time.Duration
	tkt        *time.Ticker
	callback   func(...interface{})
}

//NewDelayCallback 创建延迟回调对象
func NewDelayCallback(delayTime time.Duration, firstDelay time.Duration, callback func(...interface{})) (*DelayCallback, error) {
	if callback == nil {
		return nil, fmt.Errorf("回调函数不能为nil")
	}

	tc := &DelayCallback{delayTime: delayTime, callback: callback, firstDelay: firstDelay}
	tc.msg = make(chan []interface{}, 1)
	go tc.call()
	return tc, nil
}

//Push 添加触发事件
func (t *DelayCallback) Push(p ...interface{}) {
	select {
	case t.msg <- p:
	default:
	}
}

//call 回调执行
func (t *DelayCallback) call() {
	time.Sleep(t.firstDelay)
	select {
	case v := <-t.msg:
		// t.callback(v...)
		/*change by champly 2016年12月08日15:36:05*/
		if t.callback != nil {
			t.callback(v...)
		}
		/*end*/
	}
	t.tkt = time.NewTicker(t.delayTime)
START:
	for {
		select {
		case _, ok := <-t.tkt.C:
			{
				if !ok {
					break START
				}
				select {
				case v, ok := <-t.msg:
					if !ok {
						break START
					}
					// t.callback(v...)
					/*change by champly 2016年12月08日15:36:05*/
					if t.callback != nil {
						t.callback(v...)
					}
					/*end*/
				default:
				}
			}
		}
	}
}

//Close 关闭回调对象
func (t *DelayCallback) Close() {
	if t.tkt != nil {
		t.tkt.Stop()
	}
	close(t.msg)
	/*add by champly 2016年12月08日15:36:23*/
	// 直接关闭报错
	t.callback = nil
	/*end*/
}
