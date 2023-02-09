package initialize

import (
	_ "yzgin/source/example"
	_ "yzgin/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
