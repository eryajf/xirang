package response

import "github.com/eryajf/xirang/model"

type MenuListRsp struct {
	Total int64        `json:"total"`
	Menus []model.Menu `json:"menus"`
}
