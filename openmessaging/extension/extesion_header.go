package extension

type ExtensionHeader interface {
	SetPartition(partition int) (ExtensionHeader, error)

	SetOffset(offset int64) (ExtensionHeader, error)

	SetCorrelationId(correlationId string) (ExtensionHeader, error)

	SetTransactionId(transactionId string) (ExtensionHeader, error)

	SetStoreTimestamp(storeTimestamp int64) (ExtensionHeader, error)

	SetStoreHost(storeHost string) (ExtensionHeader, error)

	SetMessageKey(messageKey string) (ExtensionHeader, error)

	SetTraceId(traceId string) (ExtensionHeader, error)

	SetDelayTime(delayTime int64) (ExtensionHeader, error)

	SetExpireTime(expireTime int64) (ExtensionHeader, error)

	GetPartiton() (int, error)

	GetOffset() (int64, error)

	GetCorrelationId() (string, error)

	GetTransactionId() (string, error)

	GetStoreTimestamp() (int64, error)

	GetStoreHost() (string, error)

	GetDelayTime() (int64, error)

	GetExpireTime() (int64, error)

	GetMessageKey() (string, error)

	GetTraceId() (string, error)
}
