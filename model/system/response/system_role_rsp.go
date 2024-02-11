package response

import "github.com/eryajf/xirang/model/system"

type RoleListRsp struct {
	Total int64         `json:"total"`
	Roles []system.Role `json:"roles"`
}
