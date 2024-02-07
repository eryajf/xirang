package request

// DomainListReq 获取资源列表结构体
type DomainListReq struct {
	DomainID string `json:"domainID" form:"domainID"`
	Name     string `json:"name" form:"name"`
	Remark   string `json:"remark" form:"remark"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// DomainAddReq 添加资源结构体
type DomainAddReq struct {
	DomainID string `json:"domainID" validate:"required,min=1,max=64"`
	Name     string `json:"name" validate:"required,min=1,max=64"`
	Remark   string `json:"remark" validate:"min=0,max=128"`
}

// DomainUpdateReq 更新资源结构体
type DomainUpdateReq struct {
	ID       uint   `json:"id" validate:"required"`
	DomainID string `json:"domainID" validate:"min=0,max=64"`
	Name     string `json:"name" validate:"min=0,max=64"`
	Remark   string `json:"remark" validate:"min=0,max=128"`
}

// DomainDeleteReq 删除资源结构体
type DomainDeleteReq struct {
	DomainIds []uint `json:"domainIds" validate:"required"`
}

// DomainGetTreeReq 获取资源树结构体
type DomainGetTreeReq struct {
}
