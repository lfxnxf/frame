package proxy

import (
	"context"
	"fmt"

	"github.com/lfxnxf/frame/BackendPlatform/golang/kafka"
	"github.com/lfxnxf/frame/logic/inits"
)

type KafkaProducer struct {
	name string
}

type KafkaSyncProducer struct {
	name string
}

type KafkaConsumer struct {
	name string
}

func Topic(ctx context.Context, topic string) string {
	d, _ := inits.FromContext(ctx)
	if d != nil {
		return fmt.Sprintf("%s_%s", topic, d.Namespace)
	}
	return topic
}

func ContextFromMessage(m *kafka.ConsumerMessage) context.Context {
	key, ok := kafka.NSKey(m.Context())
	if !ok {
		return m.Context()
	}
	return inits.WithAPPKey(m.Context(), key)
}

func InitKafkaProducer(name string) *KafkaProducer {
	return &KafkaProducer{name}
}

func InitKafkaSyncProducer(name string) *KafkaSyncProducer {
	return &KafkaSyncProducer{name}
}

func InitKafkaConsumer(name string) *KafkaConsumer {
	return &KafkaConsumer{name}
}

func (k *KafkaSyncProducer) Send(ctx context.Context, message *kafka.ProducerMessage) (int32, int64, error) {
	d, _ := inits.FromContext(ctx)
	if d != nil {
		ctx = kafka.WithNSKey(ctx, d.Namespace)
	}
	return inits.SyncProducerClient(ctx, k.name).Send(ctx, message)
}

func (k *KafkaSyncProducer) SendSyncMsg(ctx context.Context, topic string, key string, msg []byte) (int32, int64, error) {
	m := &kafka.ProducerMessage{}
	m.Key = key
	m.Topic = Topic(ctx, topic)
	m.Value = msg
	return k.Send(ctx, m)
}

func (k *KafkaProducer) Send(ctx context.Context, message *kafka.ProducerMessage) (int32, int64, error) {
	d, _ := inits.FromContext(ctx)
	if d != nil {
		ctx = kafka.WithNSKey(ctx, d.Namespace)
	}
	return inits.KafkaProducerClient(ctx, k.name).Send(ctx, message)
}

func (k *KafkaProducer) SendKeyMsg(ctx context.Context, topic string, key string, msg []byte) error {
	m := &kafka.ProducerMessage{}
	m.Key = key
	m.Topic = Topic(ctx, topic)
	m.Value = msg
	_, _, err := k.Send(ctx, m)
	return err
}

// nolint:staticcheck
func (k *KafkaProducer) SendMsg(ctx context.Context, topic string, msg []byte) error {
	d, _ := inits.FromContext(ctx)
	if d != nil {
		ctx = kafka.WithNSKey(ctx, d.Namespace)
	}
	m := &kafka.ProducerMessage{
		Topic: topic,
		Key:   "",
		Value: msg,
	}
	_, _, err := inits.KafkaProducerClient(ctx, k.name).Send(ctx, m)
	return err
}

func (k *KafkaProducer) Errors(ctx context.Context) <-chan *kafka.ProducerError {
	d, _ := inits.FromContext(ctx)
	if d != nil {
		ctx = kafka.WithNSKey(ctx, d.Namespace)
	}
	return inits.KafkaProducerClient(ctx, k.name).Errors()
}

func (k *KafkaProducer) Success(ctx context.Context) <-chan *kafka.ProducerMessage {
	d, _ := inits.FromContext(ctx)
	if d != nil {
		ctx = kafka.WithNSKey(ctx, d.Namespace)
	}
	return inits.KafkaProducerClient(ctx, k.name).Success()
}

func (k *KafkaConsumer) GetMessages(ctx context.Context) <-chan *kafka.ConsumerMessage {
	return inits.KafkaConsumeClient(ctx, k.name).GetMessages()
}

func (k *KafkaConsumer) CommitUpto(ctx context.Context, message *kafka.ConsumerMessage) {
	inits.KafkaConsumeClient(ctx, k.name).CommitUpto(message)
}

func (k *KafkaConsumer) Messages(ctx context.Context, closeChan chan bool, maxQueueSize int) chan []byte {
	return inits.KafkaConsumeClient(ctx, k.name).Messages(closeChan, maxQueueSize)
}
