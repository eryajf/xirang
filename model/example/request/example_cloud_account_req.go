package request

// CloudAccountListReq 获取资源列表结构体
type CloudAccountListReq struct {
	CloudName string `json:"cloudName" form:"cloudName"`
	CloudType string `json:"cloudType" form:"cloudType"`
	SecretId  string `json:"secretId" form:"secretId"`
	PageNum   int    `json:"pageNum" form:"pageNum"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
}

// CloudAccountAddReq 添加资源结构体
type CloudAccountAddReq struct {
	CloudName string `json:"cloudName" validate:"required,min=1,max=64"`
	CloudType string `json:"cloudType" validate:"required,min=1,max=64"`
	SecretId  string `json:"secretId" validate:"required,min=1,max=64"`
	SecretKey string `json:"secretKey" validate:"required,min=1,max=64"`
	Remark    string `json:"remark" validate:"min=0,max=128"`
}

// CloudAccountUpdateReq 更新资源结构体
type CloudAccountUpdateReq struct {
	ID        uint   `json:"id" validate:"required"`
	CloudName string `json:"cloudName" validate:"min=0,max=64"`
	CloudType string `json:"cloudType" validate:"min=0,max=64"`
	SecretId  string `json:"secretId" validate:"min=0,max=64"`
	SecretKey string `json:"secretKey" validate:"min=0,max=64"`
	Remark    string `json:"remark" validate:"min=0,max=128"`
}

// CloudAccountDeleteReq 删除资源结构体
type CloudAccountDeleteReq struct {
	CloudAccountIds []uint `json:"cloudAccountIds" validate:"required"`
}

// CloudAccountGetTreeReq 获取资源树结构体
type CloudAccountGetTreeReq struct {
}
