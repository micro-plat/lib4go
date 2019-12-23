package lmq

import (
	"errors"
	"time"

	"github.com/micro-plat/lib4go/mq"
	"github.com/micro-plat/lib4go/queue/lmq"
)

//LMQProducer Producer
type LMQProducer struct {
	backupMsg chan *mq.ProcuderMessage
	client    *lmq.LMQClient
	closeCh   chan struct{}
	done      bool
}

//NewLMQProducer 创建新的producer
func NewLMQProducer(address string, opts ...mq.Option) (producer *LMQProducer, err error) {
	m, err := lmq.New([]string{address}, "")
	if err != nil {
		return nil, err
	}
	producer = &LMQProducer{client: m}
	producer.closeCh = make(chan struct{})
	return
}

//Connect  循环连接服务器
func (producer *LMQProducer) Connect() (err error) {
	return nil
}

//GetBackupMessage 获取备份数据
func (producer *LMQProducer) GetBackupMessage() chan *mq.ProcuderMessage {
	return producer.backupMsg
}

//Send 发送消息
func (producer *LMQProducer) Send(queue string, msg string, timeout time.Duration) (err error) {
	if producer.done {
		return errors.New("lmq producer 已关闭")
	}
	err = producer.client.Push(queue, msg)
	return
}

//Close 关闭当前连接
func (producer *LMQProducer) Close() {
	producer.done = true
	close(producer.closeCh)
	close(producer.backupMsg)
	producer.client.Close()
}

type LMQProducerResolver struct {
}

func (s *LMQProducerResolver) Resolve(address string, opts ...mq.Option) (mq.MQProducer, error) {
	return NewLMQProducer(address, opts...)
}
func init() {
	mq.RegisterProducer("lmq", &LMQProducerResolver{})
}
