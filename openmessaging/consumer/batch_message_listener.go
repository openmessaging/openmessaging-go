package consumer

import (
	. "git.ucloudadmin.com/paas_vppd/prj-wintersoldier/openmessaging-go/message"
)

type Context interface {
	Success(messages MessageReceipt) error
	Ack() error
}

type BatchMessageListener interface {
	OnReceived(message Message, context Context) error
}
