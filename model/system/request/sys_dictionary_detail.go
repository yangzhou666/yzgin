package request

import (
	"yzgin/model/common/request"
	"yzgin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
