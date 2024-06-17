package request

// DeptListReq 获取资源列表结构体
type DeptListReq struct {
	Name     string `json:"name" form:"name"`
	Remark   string `json:"remark" form:"remark"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// DeptListAllReq 获取资源列表结构体，不分页
type DeptListAllReq struct {
	Name               string `json:"name" form:"name"`
	DeptType           string `json:"deptType" form:"deptType"`
	Remark             string `json:"remark" form:"remark"`
	Source             string `json:"source" form:"source"`
	SourceDeptId       string `json:"sourceDeptId"`
	SourceDeptParentId string `json:"SourceDeptParentId"`
}

// DeptAddReq 添加资源结构体
type DeptAddReq struct {
	Name string `json:"name" validate:"required,min=1,max=20"`
	//父级Id 大于等于0 必填
	ParentId uint   `json:"parentId" validate:"omitempty,min=0"`
	Remark   string `json:"remark" validate:"min=0,max=100"` // 分组的中文描述
}

// DeptUpdateReq 更新资源结构体
type DeptUpdateReq struct {
	ID     uint   `json:"id" form:"id" validate:"required"`
	Name   string `json:"name" validate:"required,min=1,max=20"`
	Remark string `json:"remark" validate:"min=0,max=100"` // 分组的中文描述
}

// DeptDeleteReq 删除资源结构体
type DeptDeleteReq struct {
	DeptIds []uint `json:"deptIds" validate:"required"`
}

// DeptGetTreeReq 获取资源树结构体
type DeptGetTreeReq struct {
	Name     string `json:"name" form:"name"`
	Remark   string `json:"remark" form:"remark"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

type DeptAddUserReq struct {
	DeptID  uint   `json:"deptId" validate:"required"`
	UserIds []uint `json:"userIds" validate:"required"`
}

type DeptRemoveUserReq struct {
	DeptID  uint   `json:"deptId" validate:"required"`
	UserIds []uint `json:"userIds" validate:"required"`
}

// DeptInDeptReq 在分组内的用户
type DeptInDeptReq struct {
	DeptID   uint   `json:"deptId" form:"deptId" validate:"required"`
	Nickname string `json:"nickname" form:"nickname"`
}

// UserNoInDeptReq 不在分组内的用户
type UserNoInDeptReq struct {
	DeptID   uint   `json:"deptId" form:"deptId" validate:"required"`
	Nickname string `json:"nickname" form:"nickname"`
}
