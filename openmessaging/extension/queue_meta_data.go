package extension

type Partition interface {
	PartitionId() (int, error)
	PartitonHost() (string, error)
}
type QueueMetaData interface {
	QueueName() (string, error)
	Partitions() (list Partition, err error)
}
