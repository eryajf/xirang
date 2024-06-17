package system

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/eryajf/xirang/model/system"
	systemReq "github.com/eryajf/xirang/model/system/request"
	systemRsp "github.com/eryajf/xirang/model/system/response"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
)

type DeptLogic struct{}

// Add 添加数据
func (l DeptLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.DeptAddReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	if deptService.Exist(tools.H{"name": r.Name}) {
		return nil, tools.NewValidatorError(fmt.Errorf("该部门已存在"))
	}

	// 获取当前用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
	}

	dept := system.Dept{
		ParentId: r.ParentId,
		Name:     r.Name,
		Remark:   r.Remark,
		Creator:  ctxUser.Username,
	}

	// 然后在数据库中创建组
	err = deptService.Add(&dept)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("向MySQL创建部门失败"))
	}

	return nil, nil
}

// List 数据列表
func (l DeptLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.DeptListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	// 获取数据列表
	depts, err := deptService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取部门列表失败: %s", err.Error()))
	}

	rets := make([]system.Dept, 0)
	for _, dept := range depts {
		rets = append(rets, *dept)
	}
	count, err := deptService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取部门总数失败"))
	}

	return systemRsp.DeptListRsp{
		Total: count,
		Depts: rets,
	}, nil
}

// GetTree 数据树
func (l DeptLogic) GetTree(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.DeptListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = r
	_ = c

	// var depts []*system.Dept
	// depts, err := deptService.ListTree(r)
	// if err != nil {
	// 	return nil, tools.NewMySqlError(fmt.Errorf("获取资源列表失败: " + err.Error()))
	// }

	// tree := genDeptTree(0, depts)

	return nil, nil
}

// Update 更新数据
func (l DeptLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.DeptUpdateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": int(r.ID)}
	if !deptService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("部门不存在"))
	}

	// 获取当前登陆用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户失败"))
	}

	oldDept := new(system.Dept)
	err = deptService.Find(filter, oldDept)
	if err != nil {
		return nil, tools.NewMySqlError(err)
	}

	newDept := system.Dept{
		Model:   oldDept.Model,
		Name:    r.Name,
		Remark:  r.Remark,
		Creator: ctxUser.Username,
	}

	err = deptService.Update(&newDept)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("向MySQL更新部门失败"))
	}
	return nil, nil
}

// Delete 删除数据
func (l DeptLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.DeptDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	for _, id := range r.DeptIds {
		filter := tools.H{"id": int(id)}
		if !deptService.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("有部门不存在"))
		}
	}

	depts, err := deptService.GetDeptByIds(r.DeptIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取部门列表失败: %s", err.Error()))
	}

	// 从MySQL中删除
	err = deptService.Delete(depts)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("删除接口失败: %s", err.Error()))
	}

	return nil, nil
}

// AddUser 添加用户到部门
func (l DeptLogic) AddUser(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.DeptAddUserReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.DeptID}

	if !deptService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("部门不存在"))
	}

	users, err := userService.GetUserByIds(r.UserIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户列表失败: %s", err.Error()))
	}

	dept := new(system.Dept)
	err = deptService.Find(filter, dept)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取部门失败: %s", err.Error()))
	}

	// 先添加到MySQL
	err = deptService.AddUserToDept(dept, users)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("添加用户到部门失败: %s", err.Error()))
	}

	return nil, nil
}

func (l DeptLogic) updataUser(newUser *system.User) error {
	err := userService.Update(newUser)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("在MySQL更新用户失败：" + err.Error()))
	}
	return nil
}

// RemoveUser 移除用户
func (l DeptLogic) RemoveUser(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.DeptRemoveUserReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	filter := tools.H{"id": r.DeptID}

	if !deptService.Exist(filter) {
		return nil, tools.NewMySqlError(fmt.Errorf("部门不存在"))
	}

	users, err := userService.GetUserByIds(r.UserIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户列表失败: %s", err.Error()))
	}

	dept := new(system.Dept)
	err = deptService.Find(filter, dept)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取部门失败: %s", err.Error()))
	}

	// 再操作MySQL
	err = deptService.RemoveUserFromDept(dept, users)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("将用户从MySQL移除失败: %s", err.Error()))
	}

	for _, user := range users {
		oldData := new(system.User)
		err = userService.Find(tools.H{"id": user.ID}, oldData)
		if err != nil {
			return nil, tools.NewMySqlError(err)
		}
		newData := oldData

		var newDepts []string
		var newDeptIds []string
		// 删掉移除的部门名字
		for _, v := range strings.Split(oldData.Departments, ",") {
			if v != dept.Name {
				newDepts = append(newDepts, v)
			}
		}
		// 删掉移除的部门id
		for _, v := range strings.Split(oldData.DepartmentId, ",") {
			if v != strconv.Itoa(int(r.DeptID)) {
				newDeptIds = append(newDeptIds, v)
			}
		}

		newData.Departments = strings.Join(newDepts, ",")
		newData.DepartmentId = strings.Join(newDeptIds, ",")
		err = l.updataUser(newData)
		if err != nil {
			return nil, tools.NewOperationError(fmt.Errorf("处理用户的部门数据失败:" + err.Error()))
		}
	}

	return nil, nil
}

// UserInDept 在部门内的用户
func (l DeptLogic) UserInDept(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	return nil, nil
}

// UserNoInDept 不在部门内的用户
func (l DeptLogic) UserNoInDept(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	return nil, nil
}
