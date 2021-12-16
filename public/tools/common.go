package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Msg   string      `json:"msg"`
	Error string      `json:"error,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
	// AlreadyExistErr 已存在
	AlreadyExistErr = 40002
	// NotExistErr 不存在
	NotExistErr = 40004
	// CreatedSuccess 创建成功
	CreatedSuccess = "创建成功！"
	// UpdatedSuccess 更新成功
	UpdatedSuccess = "更新成功！"
	// DeletedSuccess 删除成功
	DeletedSuccess = "删除成功！"
	// GetSuccess 查询成功
	GetSuccess = "查询成功！"

	// CreatedFail 创建失败
	CreatedFail = "创建失败！"
	// UpdatedFail 更新失败
	UpdatedFail = "更新失败！"
	// DeletedFail 删除失败
	DeletedFail = "删除失败！"
	// NotFound 查询失败
	NotFound = "未找到相关内容或者数据为空！"

	// ExecSuccess 执行成功
	ExecSuccess = "执行成功"
	// ExecFail 执行失败
	ExecFail = "执行失败"

	// SSH pass
	SSHPassWd = "root"
)

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "未登录",
	}
}

// HasError 通用错误处理
func HasError(errCode int, msg string, err error) Response {
	// 如果msg为空，则使用错误信息填充
	if msg == "" {
		msg = err.Error()
	}
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// 兼容函数
func Custum(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, data)
}

// DataList 分页时数据信息
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// BuildListResponse 列表构建器
func BuildListResponse(code int, items interface{}, total uint) Response {
	return Response{
		Code: code,
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
