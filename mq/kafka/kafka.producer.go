package kafka

/*
import (
	"errors"
	"strings"
	"sync"
	"time"

	"fmt"

	"github.com/Shopify/sarama"
	"github.com/qxnw/lib4go/concurrent/cmap"
	"github.com/qxnw/lib4go/logger"
	"github.com/qxnw/lib4go/mq"
)

//KafkaProducer Producer
type KafkaProducer struct {
	address    string
	messages   chan *mq.ProcuderMessage
	backupMsg  chan *mq.ProcuderMessage
	queues     cmap.ConcurrentMap
	connecting bool
	closeCh    chan struct{}
	done       bool
	once       sync.Once
	lk         sync.Mutex
	header     []string
	*mq.OptionConf
}
type kafkaProducer struct {
	producer sarama.SyncProducer
	msgQueue chan *sarama.ProducerMessage
}

//NewKafkaProducer 创建新的producer
func NewKafkaProducer(address string, opts ...mq.Option) (producer *KafkaProducer, err error) {
	producer = &KafkaProducer{address: address}
	producer.queues = cmap.New(2)
	producer.OptionConf = &mq.OptionConf{}
	producer.messages = make(chan *mq.ProcuderMessage, 10000)
	producer.backupMsg = make(chan *mq.ProcuderMessage, 100)
	producer.closeCh = make(chan struct{})
	for _, opt := range opts {
		opt(producer.OptionConf)
	}
	if producer.Logger == nil {
		producer.Logger = logger.GetSession("mq.producer", logger.CreateSession())
	}
	return
}

//Connect  循环连接服务器
func (producer *KafkaProducer) Connect() error {
	go producer.sendLoop()
	return nil
}

//sendLoop 循环发送消息
func (producer *KafkaProducer) sendLoop() {
	if producer.done {
		producer.disconnect()
		return
	}
	if producer.Retry {
	Loop1:
		for {
			select {
			case msg, ok := <-producer.backupMsg:
				if !ok {
					break Loop1
				}
				pd, ok := producer.queues.Get(msg.Queue)
				if !ok {
					select {
					case producer.backupMsg <- msg:
					default:
						producer.Logger.Errorf("重试发送失败，备份队列已满无法放入队列(%s):%s", msg.Queue, msg.Data)
					}
					continue
				}
				producerConn := pd.(*kafkaProducer)
				_, _, err := producerConn.producer.SendMessage(&sarama.ProducerMessage{Topic: msg.Queue, Partition: 0, Value: sarama.StringEncoder(msg.Data)})
				if err != nil {
					select {
					case producer.backupMsg <- msg:
					default:
						producer.Logger.Errorf("发送失败，备份队列已满无法放入队列(%s):%s", msg.Queue, msg.Data)
					}
				}
			case msg, ok := <-producer.messages:
				if !ok {
					break Loop1
				}
				pd, ok := producer.queues.Get(msg.Queue)
				if !ok {
					select {
					case producer.backupMsg <- msg:
					default:
						producer.Logger.Errorf("消息无法放入备份队列(%s):%s", msg.Queue, msg.Data)
					}
					producer.Logger.Errorf("消息无法从缓存中获取producer:%s,%s", msg.Queue, msg.Data)
					continue
				}
				producerConn := pd.(*kafkaProducer)
				_, _, err := producerConn.producer.SendMessage(&sarama.ProducerMessage{Topic: msg.Queue, Partition: 0, Value: sarama.StringEncoder(msg.Data)})
				if err != nil {
					select {
					case producer.backupMsg <- msg:
					default:
						producer.Logger.Errorf("消息无法放入备份队列(%s):%s", msg.Queue, msg.Data)
					}
				}
			}
		}
	} else {
	Loop2:
		for {
			select {
			case msg, ok := <-producer.messages:
				fmt.Println("send.msg:", msg)
				if !ok {
					break Loop2
				}
				pd, ok := producer.queues.Get(msg.Queue)
				if !ok {
					select {
					case producer.backupMsg <- msg:
					default:
						producer.Logger.Errorf("消息无法放入备份队列(%s):%s", msg.Queue, msg.Data)
					}
					producer.Logger.Errorf("消息无法从缓存中获取producer:%s,%s", msg.Queue, msg.Data)
					continue
				}
				producerConn := pd.(*kafkaProducer)
				//, Timestamp: msg.Timeout
				_, _, err := producerConn.producer.SendMessage(&sarama.ProducerMessage{Topic: msg.Queue, Partition: 0, Value: sarama.StringEncoder(msg.Data)})
				if err != nil {
					select {
					case producer.backupMsg <- msg:
					default:
						producer.Logger.Errorf("消息无法放入备份队列(%s):%s", msg.Queue, msg.Data)
					}
				}
			}
		}
	}
	if producer.done { //关闭连接
		producer.disconnect()
		return
	}
}
func (producer *KafkaProducer) disconnect() {

}

//GetBackupMessage 获取备份数据
func (producer *KafkaProducer) GetBackupMessage() chan *mq.ProcuderMessage {
	return producer.backupMsg
}

//Send 发送消息
func (producer *KafkaProducer) Send(queue string, msg string, timeout time.Duration) (err error) {
	if producer.done {
		return errors.New("mq producer 已关闭")
	}
	producer.queues.SetIfAbsentCb(queue, func(i ...interface{}) (interface{}, error) {
		var err error
		c := &kafkaProducer{}
		config := sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Partitioner = sarama.NewManualPartitioner

		c.producer, err = sarama.NewSyncProducer(strings.Split(producer.address, ","), config)
		c.msgQueue = make(chan *sarama.ProducerMessage, 10)

		//&sarama.ProducerMessage{Topic: *topic, Partition: int32(*partition)}
		return c, err
	})

	pm := &mq.ProcuderMessage{Queue: queue, Data: msg, Timeout: timeout}
	select {
	case producer.messages <- pm:
		return nil
	default:
		return errors.New("producer无法连接到MQ服务器，消息队列已满无法发送")
	}
}

//Close 关闭当前连接
func (producer *KafkaProducer) Close() {
	producer.done = true
	producer.once.Do(func() {
		close(producer.closeCh)
		close(producer.messages)
		close(producer.backupMsg)
	})

}

type kafkaProducerResolver struct {
}

func (s *kafkaProducerResolver) Resolve(address string, opts ...mq.Option) (mq.MQProducer, error) {
	return NewKafkaProducer(address, opts...)
}
func init() {
	mq.RegisterProducer("kafka", &kafkaProducerResolver{})
}
*/
