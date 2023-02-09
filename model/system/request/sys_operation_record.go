package request

import (
	"yzgin/model/common/request"
	"yzgin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
