package extension

type Extension interface {
	GetQueueMetaData(queueName string) (QueueMetaData, error)
}
