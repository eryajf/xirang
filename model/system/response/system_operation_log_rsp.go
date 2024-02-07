package response

import "github.com/eryajf/xirang/model/system"

type LogListRsp struct {
	Total int64                 `json:"total"`
	Logs  []system.OperationLog `json:"logs"`
}
