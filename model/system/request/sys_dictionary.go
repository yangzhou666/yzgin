package request

import (
	"yzgin/model/common/request"
	"yzgin/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
