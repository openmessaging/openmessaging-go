package message

type Message interface {
	GetData() ([]byte, error)
	SetData(data []byte) error
	Headers() (Headers, error)
	Properties() (KeyValue, error)
}
