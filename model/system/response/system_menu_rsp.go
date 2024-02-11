package response

import "github.com/eryajf/xirang/model/system"

type MenuListRsp struct {
	Total int64         `json:"total"`
	Menus []system.Menu `json:"menus"`
}
