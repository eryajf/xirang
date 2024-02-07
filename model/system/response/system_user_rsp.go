package response

import "github.com/eryajf/xirang/model/system"

type UserListRsp struct {
	Total int           `json:"total"`
	Users []system.User `json:"users"`
}
