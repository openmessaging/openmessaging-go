package common

import (
	"testing"

	"github.com/openmessaging-go/openmessaging"
	"github.com/openmessaging-go/openmessaging/consumer"
	"github.com/openmessaging-go/openmessaging/manager"
	"github.com/openmessaging-go/openmessaging/producer"
)

func TestCycleImportIssue(t *testing.T) {
	s := &sampleMessagingAccessPoint{}
	v, err := s.Version()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

type sampleMessagingAccessPoint struct {
}

func (sampleMessagingAccessPoint) Version() (string, error) {
	return "0.0.1", nil
}

func (sampleMessagingAccessPoint) Attributes() (openmessaging.KeyValue, error) {
	panic("implement me")
}

func (sampleMessagingAccessPoint) CreateProducer() (producer.Producer, error) {
	panic("implement me")
}

func (sampleMessagingAccessPoint) CreateTransactionProducer(transactionStateCheckListener producer.TransactionStateCheckListener) (producer.Producer, error) {
	panic("implement me")
}

func (sampleMessagingAccessPoint) CreateConsumer() (consumer.Consumer, error) {
	panic("implement me")
}

func (sampleMessagingAccessPoint) ResourceManager() (manager.ResourceManager, error) {
	panic("implement me")
}
