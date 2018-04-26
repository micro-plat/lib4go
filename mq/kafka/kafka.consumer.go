package kafka

/*
import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/qxnw/lib4go/concurrent/cmap"
	"github.com/qxnw/lib4go/mq"
)

//KafkaConsumer kafka consumer
type KafkaConsumer struct {
	address   string
	consumers cmap.ConcurrentMap
	quitChan  chan struct{}
	*mq.OptionConf
}
type kafkaConsumer struct {
	consumer sarama.Consumer
	msgQueue chan *sarama.ConsumerMessage
}

//NewKafkaConsumer 初始化kafka Consumer
func NewKafkaConsumer(address string, opts ...mq.Option) (kafka *KafkaConsumer, err error) {
	kafka = &KafkaConsumer{address: address, quitChan: make(chan struct{}, 0)}
	kafka.OptionConf = &mq.OptionConf{}
	kafka.consumers = cmap.New(2)
	for _, opt := range opts {
		opt(kafka.OptionConf)
	}
	return
}

//Connect 连接到服务器
func (k *KafkaConsumer) Connect() error {
	return nil
}

//Consume 订阅消息
func (k *KafkaConsumer) Consume(queue string, call func(mq.IMessage)) (err error) {
	fmt.Println("启动接受消息")
	_, cnsmr, _ := k.consumers.SetIfAbsentCb(queue, func(i ...interface{}) (interface{}, error) {
		c := &kafkaConsumer{}
		//c.consumer = kafka.NewBrokerConsumer(k.address, queue, 0, 0, 1048576)
		//c.msgQueue = make(chan *kafka.Message, 10000)
		c.consumer, err = sarama.NewConsumer(strings.Split(k.address, ","), nil)
		c.msgQueue = make(chan *sarama.ConsumerMessage, 10000)
		return c, nil
	})
	consumer := cnsmr.(*kafkaConsumer)

	var (
		chanmsg = make(chan *sarama.ConsumerMessage, 10000)
		closing = make(chan struct{})
	)

	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Kill, os.Interrupt)
		<-signals
		fmt.Println("Initiating shutdown of consumer...")
		close(closing)
	}()

	pc, err := consumer.consumer.ConsumePartition(queue, 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("Failed to start consumer for partition %d: %s", 0, err)
	}

	go func(pc sarama.PartitionConsumer) {
		<-closing
		pc.AsyncClose()
	}(pc)

	fmt.Println("22222222222")
	go func(pc sarama.PartitionConsumer) {
		for {
			select {
			case message := <-pc.Messages():
				chanmsg <- message
			}
		}
	}(pc)

	go func() {
		fmt.Println("接受消息，使用对应函数")
	LOOP:
		for {
			select {
			case msg, ok := <-chanmsg:
				if ok {
					go call(NewKafkaMessage(msg))
				} else {
					break LOOP
				}
			}
		}
	}()

	//close(chanmsg)

	//if err := consumer.consumer.Close(); err != nil {
	//fmt.Println("Failed to close consumer: ", err)
	//}
	return nil
}

//UnConsume 取消注册消费
func (k *KafkaConsumer) UnConsume(queue string) {
	if c, ok := k.consumers.Get(queue); ok {
		consumer := c.(*kafkaConsumer)
		close(consumer.msgQueue)
	}
}

//Close 关闭当前连接
func (k *KafkaConsumer) Close() {
	close(k.quitChan)
	k.consumers.IterCb(func(key string, value interface{}) bool {
		consumer := value.(*kafkaConsumer)
		close(consumer.msgQueue)
		return true
	})
}

type kafkaConsumerResolver struct {
}

func (s *kafkaConsumerResolver) Resolve(address string, opts ...mq.Option) (mq.MQConsumer, error) {
	return NewKafkaConsumer(address, opts...)
}
func init() {
	mq.RegisterCosnumer("kafka", &kafkaConsumerResolver{})
}
*/
