package message

type KeyValue interface {
	ContainsKey(key string) (bool, error)
	GetDouble(key string) (float64, error)
	GetInt(key string) (int16, error)
	GetLong(key string) (int64, error)
	GetShort(key string) (int16, error)
	GetString(key string) (string, error)
	KeySet() ([]string, error)
	PutFloat64(key string, value float64) (KeyValue, error)
	PutInt(key string, value int) (KeyValue, error)
	PutInt16(key string, value int16) (KeyValue, error)
	PutInt64(key string, value int64) (KeyValue, error)
	PutString(key string, value string) (KeyValue, error)
}
