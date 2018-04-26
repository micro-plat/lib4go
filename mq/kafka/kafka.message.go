package kafka

/*
import (
	"github.com/Shopify/sarama"
)

type KafkaMessage struct {
	msg     *sarama.ConsumerMessage
	Message string
}

//Ack
func (m *KafkaMessage) Ack() error {
	return nil
	//m.s.conn.Ack(m.msg.Headers)
}
func (m *KafkaMessage) Nack() error {
	return nil
	//m.s.conn.Nack(m.msg.Headers)
}
func (m *KafkaMessage) GetMessage() string {
	return m.Message
}

//NewMessage
func NewKafkaMessage(msg *sarama.ConsumerMessage) *KafkaMessage {
	//return &StompMessage{s: s, msg: msg, Message: string(msg.Body)}
	return &KafkaMessage{msg: msg, Message: string(msg.Value)}
}
*/
