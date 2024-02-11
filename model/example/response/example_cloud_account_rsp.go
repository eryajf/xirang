package response

import "github.com/eryajf/xirang/model/example"

type CloudAccountListRsp struct {
	Total         int64                  `json:"total"`
	CloudAccounts []example.CloudAccount `json:"cloudAccounts"`
}
