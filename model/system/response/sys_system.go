package response

import "yzgin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
