package openmessaging_go

import (
	. "git.ucloudadmin.com/paas_vppd/prj-wintersoldier/openmessaging-go/message"
)

type TransactionalContext interface {
	commit()
	rollback()
	unKnown()
}

type TransactionStateCheckListener interface {
	check(message *Message, context TransactionalContext)
}
