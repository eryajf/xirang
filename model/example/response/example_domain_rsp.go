package response

import "github.com/eryajf/xirang/model/example"

type DomainListRsp struct {
	Total   int64            `json:"total"`
	Domains []example.Domain `json:"domains"`
}
