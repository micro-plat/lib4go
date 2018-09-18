package xsync

import (
	"sync"
	"github.com/micro-plat/lib4go/utility"
)

//Queue 顺序队列
type Queue struct{
   s []*Ticket
   lk sync.Mutex
}

//NewQueue 创建队列
func NewQueue()*Queue{
	return &Queue{s:make([]*Ticket,0,1)}
}
//Get 获取一个门票
func(d *Queue)Get()*Ticket{
	d.lk.Lock()
	defer d.lk.Unlock()
	Ticket:=newTicket(d)
	d.s=append(d.s,Ticket)
	if len(d.s)==1{
		d.s[0].notify()
	}
	return Ticket
}
//quit 放弃排队
func(d *Queue)quit(s *Ticket){
	d.lk.Lock()
	defer d.lk.Unlock()
	for i,sub:=range d.s{
		if sub.id==s.id{
			nsubs:=make([]*Ticket,0,len(d.s))
			nsubs=append(nsubs,d.s[0:i]...)
			nsubs=append(nsubs,d.s[i+1:]...)
			d.s=nsubs
		}
	}
	if len(d.s)>0{
		d.s[0].notify()
	}
}

//Ticket 目标对象
type Ticket struct{
	id string
	msg chan int
	notified bool
	ch chan struct{}
    d *Queue
}
func newTicket(d *Queue)*Ticket{
	return &Ticket{
		id:utility.GetGUID(),
		d:d,
		msg:make(chan int,1),
		ch:make(chan struct{}),
	}
}
func(s *Ticket)notify(){
	select{
		case  s.msg<-1:
		default:
	}
}

//Wait 等待叫号
func(s *Ticket)Wait()bool{
	for{
		select{
		case <-s.msg:
			return true
		case <-s.ch:
			return false
		}
	}
}

//Done 任务完成
func (s *Ticket)Done(){
	s.d.quit(s)
	s.Quit()

}

//Quit 放弃排队
func(s *Ticket)Quit(){
	close(s.ch)
}
