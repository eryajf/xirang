package system

import (
	"fmt"

	"github.com/eryajf/xirang/config"
	"github.com/eryajf/xirang/model/system"
	systemReq "github.com/eryajf/xirang/model/system/request"
	systemRsp "github.com/eryajf/xirang/model/system/response"
	"github.com/eryajf/xirang/public/tools"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type UserLogic struct{}

// Add 添加数据
func (l UserLogic) Add(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.UserAddReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	if userService.Exist(tools.H{"username": r.Username}) {
		return nil, tools.NewValidatorError(fmt.Errorf("用户名已存在,请勿重复添加"))
	}
	if userService.Exist(tools.H{"mobile": r.Mobile}) {
		return nil, tools.NewValidatorError(fmt.Errorf("手机号已存在,请勿重复添加"))
	}
	if userService.Exist(tools.H{"job_number": r.JobNumber}) {
		return nil, tools.NewValidatorError(fmt.Errorf("工号已存在,请勿重复添加"))
	}
	if userService.Exist(tools.H{"mail": r.Mail}) {
		return nil, tools.NewValidatorError(fmt.Errorf("邮箱已存在,请勿重复添加"))
	}

	// 密码通过RSA解密
	// 密码不为空就解密
	if r.Password != "" {
		decodeData, err := tools.RSADecrypt([]byte(r.Password), config.Conf.System.RSAPrivateBytes)
		if err != nil {
			return nil, tools.NewValidatorError(fmt.Errorf("密码解密失败"))
		}
		r.Password = string(decodeData)
		if len(r.Password) < 6 {
			return nil, tools.NewValidatorError(fmt.Errorf("密码长度至少为6位"))
		}
	} else {
		r.Password = "123456"
	}

	// 当前登陆用户角色排序最小值（最高等级角色）以及当前登陆的用户
	currentRoleSortMin, ctxUser, err := userService.GetCurrentUserMinRoleSort(c)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("获取当前登陆用户角色排序最小值失败"))
	}

	// 根据角色id获取角色
	if r.RoleIds == nil || len(r.RoleIds) == 0 {
		r.RoleIds = []uint{2} // 默认添加为普通用户角色
	}

	roles, err := roleService.GetRolesByIds(r.RoleIds)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败"))
	}

	var reqRoleSorts []int
	for _, role := range roles {
		reqRoleSorts = append(reqRoleSorts, int(role.Sort))
	}
	// 前端传来用户角色排序最小值（最高等级角色）
	reqRoleSortMin := uint(funk.MinInt(reqRoleSorts).(int))

	// 如果登录用户的角色ID为1，亦即为管理员，则直接放行，保障管理员拥有最大权限
	if currentRoleSortMin != 1 {
		// 当前用户的角色排序最小值 需要小于 前端传来的角色排序最小值（用户不能创建比自己等级高的或者相同等级的用户）
		if currentRoleSortMin >= reqRoleSortMin {
			return nil, tools.NewValidatorError(fmt.Errorf("用户不能创建比自己等级高的或者相同等级的用户"))
		}
	}
	user := system.User{
		Username:      r.Username,
		Password:      r.Password,
		Nickname:      r.Nickname,
		GivenName:     r.GivenName,
		Mail:          r.Mail,
		JobNumber:     r.JobNumber,
		Mobile:        r.Mobile,
		Avatar:        r.Avatar,
		PostalAddress: r.PostalAddress,
		Departments:   r.Departments,
		Position:      r.Position,
		Introduction:  r.Introduction,
		Status:        r.Status,
		Creator:       ctxUser.Username,
		DepartmentId:  tools.SliceToString(r.DepartmentId, ","),
		Source:        r.Source,
		Roles:         roles,
	}

	if user.Source == "" {
		user.Source = "platform"
	}

	// 获取用户将要添加的分组
	depts, err := deptService.GetDeptByIds(tools.StringToSlice(user.DepartmentId, ","))
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("根据部门ID获取部门信息失败" + err.Error()))
	}

	err = CommonAddUser(&user, depts)
	if err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("添加用户失败" + err.Error()))
	}
	return nil, nil
}

// List 数据列表
func (l UserLogic) List(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.UserListReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	users, err := userService.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户列表失败：" + err.Error()))
	}

	rets := make([]system.User, 0)
	for _, user := range users {
		rets = append(rets, *user)
	}
	count, err := userService.ListCount(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取用户总数失败：" + err.Error()))
	}

	return systemRsp.UserListRsp{
		Total: int(count),
		Users: rets,
	}, nil
}

// Update 更新数据
func (l UserLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.UserUpdateReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	if !userService.Exist(tools.H{"id": r.ID}) {
		return nil, tools.NewMySqlError(fmt.Errorf("该记录不存在"))
	}

	// 获取当前登陆用户
	ctxUser, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户失败"))
	}

	// 获取当前登陆用户角色ID集合
	var currentRoleSorts []int
	for _, role := range ctxUser.Roles {
		currentRoleSorts = append(currentRoleSorts, int(role.Sort))
	}

	// 获取将要操作的用户角色ID集合
	var reqRoleSorts []int
	roles, _ := roleService.GetRolesByIds(r.RoleIds)
	if len(roles) == 0 {
		return nil, tools.NewValidatorError(fmt.Errorf("根据角色ID获取角色信息失败"))
	}
	for _, role := range roles {
		reqRoleSorts = append(reqRoleSorts, int(role.Sort))
	}

	// 当前登陆用户角色排序最小值（最高等级角色）
	currentRoleSortMin := funk.MinInt(currentRoleSorts).(int)
	// 前端传来用户角色排序最小值（最高等级角色）
	reqRoleSortMin := funk.MinInt(reqRoleSorts).(int)

	// 如果登录用户的角色ID为1，亦即为管理员，则直接放行，保障管理员拥有最大权限
	if currentRoleSortMin != 1 {
		// 判断是更新自己还是更新别人,如果操作的ID与登陆用户的ID一致，则说明操作的是自己
		if int(r.ID) == int(ctxUser.ID) {
			// 不能更改自己的角色
			reqDiff, currentDiff := funk.Difference(reqRoleSorts, currentRoleSorts)
			if len(reqDiff.([]int)) > 0 || len(currentDiff.([]int)) > 0 {
				return nil, tools.NewValidatorError(fmt.Errorf("不能更改自己的角色"))
			}
		}

		// 如果是更新别人，操作者不能更新比自己角色等级高的或者相同等级的用户
		minRoleSorts, err := userService.GetUserMinRoleSortsByIds([]uint{uint(r.ID)}) // 根据userIdID获取用户角色排序最小值
		if err != nil || len(minRoleSorts) == 0 {
			return nil, tools.NewValidatorError(fmt.Errorf("根据用户ID获取用户角色排序最小值失败"))
		}
		if currentRoleSortMin >= minRoleSorts[0] || currentRoleSortMin >= reqRoleSortMin {
			return nil, tools.NewValidatorError(fmt.Errorf("用户不能更新比自己角色等级高的或者相同等级的用户"))
		}
	}

	// 先获取用户信息
	oldData := new(system.User)
	err = userService.Find(tools.H{"id": r.ID}, oldData)
	if err != nil {
		return nil, tools.NewMySqlError(err)
	}

	// 拼装新的用户信息
	user := system.User{
		Model:         oldData.Model,
		Username:      r.Username,
		Nickname:      r.Nickname,
		GivenName:     r.GivenName,
		Mail:          r.Mail,
		JobNumber:     r.JobNumber,
		Mobile:        r.Mobile,
		Avatar:        r.Avatar,
		PostalAddress: r.PostalAddress,
		Departments:   r.Departments,
		Position:      r.Position,
		Introduction:  r.Introduction,
		Creator:       ctxUser.Username,
		DepartmentId:  tools.SliceToString(r.DepartmentId, ","),
		Source:        oldData.Source,
		Roles:         roles,
	}

	if err = CommonUpdateUser(oldData, &user, r.DepartmentId); err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("更新用户失败" + err.Error()))
	}

	return nil, nil
}

// Delete 删除数据
func (l UserLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.UserDeleteReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c

	for _, id := range r.UserIds {
		filter := tools.H{"id": int(id)}
		if !userService.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("有用户不存在"))
		}
	}

	// 根据用户ID获取用户角色排序最小值
	roleMinSortList, err := userService.GetUserMinRoleSortsByIds(r.UserIds)
	if err != nil || len(roleMinSortList) == 0 {
		return nil, tools.NewValidatorError(fmt.Errorf("根据用户ID获取用户角色排序最小值失败"))
	}

	// 获取当前登陆用户角色排序最小值（最高等级角色）以及当前用户
	minSort, ctxUser, err := userService.GetCurrentUserMinRoleSort(c)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("获取当前登陆用户角色排序最小值失败"))
	}

	// 不能删除自己
	if funk.Contains(r.UserIds, ctxUser.ID) {
		return nil, tools.NewValidatorError(fmt.Errorf("用户不能删除自己"))
	}

	// 不能删除比自己(登陆用户)角色排序低(等级高)的用户
	for _, sort := range roleMinSortList {
		if int(minSort) > sort {
			return nil, tools.NewValidatorError(fmt.Errorf("用户不能删除比自己角色等级高的用户"))
		}
	}

	// 再将用户从MySQL中删除
	err = userService.Delete(r.UserIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL删除用户失败: " + err.Error()))
	}

	return nil, nil
}

// ChangePwd 修改密码
func (l UserLogic) ChangePwd(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.UserChangePwdReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 前端传来的密码是rsa加密的,先解密
	// 密码通过RSA解密
	decodeOldPassword, err := tools.RSADecrypt([]byte(r.OldPassword), config.Conf.System.RSAPrivateBytes)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("原密码解析失败"))
	}
	decodeNewPassword, err := tools.RSADecrypt([]byte(r.NewPassword), config.Conf.System.RSAPrivateBytes)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("新密码解析失败"))
	}
	r.OldPassword = string(decodeOldPassword)
	r.NewPassword = string(decodeNewPassword)
	// 获取当前用户
	user, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户失败"))
	}
	if tools.NewParPasswd(user.Password) != r.OldPassword {
		return nil, tools.NewValidatorError(fmt.Errorf("原密码错误"))
	}

	// 更新密码
	err = userService.ChangePwd(user.Username, tools.NewGenPasswd(r.NewPassword))
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL更新密码失败: " + err.Error()))
	}

	return nil, nil
}

// ChangeUserStatus 修改用户状态
func (l UserLogic) ChangeUserStatus(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.UserChangeUserStatusReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}
	_ = c
	// 校验工作
	filter := tools.H{"id": r.ID}
	if !userService.Exist(filter) {
		return nil, tools.NewValidatorError(fmt.Errorf("该用户不存在"))
	}
	user := new(system.User)
	err := userService.Find(filter, user)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL查询用户失败: " + err.Error()))
	}

	if r.Status == 1 && r.Status == user.Status {
		return nil, tools.NewValidatorError(fmt.Errorf("用户已经是在职状态"))
	}
	if r.Status == 2 && r.Status == user.Status {
		return nil, tools.NewValidatorError(fmt.Errorf("用户已经是离职状态"))
	}

	// 获取当前登录用户，只有管理员才能够将用户状态改变
	// 获取当前登陆用户角色排序最小值（最高等级角色）以及当前用户
	minSort, _, err := userService.GetCurrentUserMinRoleSort(c)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("获取当前登陆用户角色排序最小值失败"))
	}

	if int(minSort) != 1 {
		return nil, tools.NewValidatorError(fmt.Errorf("只有管理员才能更改用户状态"))
	}

	err = userService.ChangeStatus(int(r.ID), int(r.Status))
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL更新用户状态失败: " + err.Error()))
	}
	return nil, nil
}

// GetUserInfo 获取用户信息
func (l UserLogic) GetUserInfo(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*systemReq.UserGetUserInfoReq)
	if !ok {
		return nil, tools.ReqAssertErr
	}

	_ = c
	_ = r

	user, err := userService.GetCurrentLoginUser(c)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取当前用户信息失败: " + err.Error()))
	}
	return user, nil
}

// ============ user logic ============

// CommonAddDept 标准创建分组
func CommonAddDept(dept *system.Dept) error {
	// 在数据库中创建组
	err := deptService.Add(dept)
	if err != nil {
		return err
	}

	// 默认创建分组之后，需要将admin添加到分组中
	adminInfo := new(system.User)
	err = userService.Find(tools.H{"id": 1}, adminInfo)
	if err != nil {
		return err
	}

	err = deptService.AddUserToDept(dept, []system.User{*adminInfo})
	if err != nil {
		return err
	}

	return nil
}

// CommonUpdateDept 标准更新分组
func CommonUpdateDept(oldDept, newDept *system.Dept) error {
	err := deptService.Update(newDept)
	if err != nil {
		return err
	}
	return nil
}

// CommonAddUser 标准创建用户
func CommonAddUser(user *system.User, depts []*system.Dept) error {
	// 用户信息的预置处理
	if user.Nickname == "" {
		user.Nickname = "佚名"
	}
	if user.GivenName == "" {
		user.GivenName = user.Nickname
	}
	if user.Introduction == "" {
		user.Introduction = user.Nickname
	}
	if user.Mail == "" {
		user.Mail = "该用户邮箱为空"
	}
	if user.JobNumber == "" {
		user.JobNumber = "该用户工号为空"
	}
	if user.Departments == "" {
		user.Departments = "默认:研发中心"
	}
	if user.Position == "" {
		user.Position = "默认:技术"
	}
	if user.PostalAddress == "" {
		user.PostalAddress = "默认:地球"
	}
	if user.Mobile == "" {
		user.Mobile = "emptyMobile"
	}

	// 先将用户添加到MySQL
	err := userService.Add(user)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("向MySQL创建用户失败：" + err.Error()))
	}

	// 处理用户归属的组
	for _, dept := range depts {
		// 先将用户和部门信息维护到MySQL
		err := deptService.AddUserToDept(dept, []system.User{*user})
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("向MySQL添加用户到分组关系失败：" + err.Error()))
		}
	}
	return nil
}

// CommonUpdateUser 标准更新用户
func CommonUpdateUser(oldUser, newUser *system.User, deptId []uint) error {
	err := userService.Update(newUser)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("在MySQL更新用户失败：" + err.Error()))
	}

	//判断部门信息是否有变化有变化则更新相应的数据库
	oldDeptIds := tools.StringToSlice(oldUser.DepartmentId, ",")
	addDeptIds, removeDeptIds := tools.ArrUintCmp(oldDeptIds, deptId)

	// 先处理添加的部门
	adddepts, err := deptService.GetDeptByIds(addDeptIds)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("根据部门ID获取部门信息失败" + err.Error()))
	}
	for _, dept := range adddepts {
		// 先将用户和部门信息维护到MySQL
		err := deptService.AddUserToDept(dept, []system.User{*newUser})
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("向MySQL添加用户到分组关系失败：" + err.Error()))
		}
	}

	// 再处理删除的部门
	removedepts, err := deptService.GetDeptByIds(removeDeptIds)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("根据部门ID获取部门信息失败" + err.Error()))
	}
	for _, dept := range removedepts {
		err := deptService.RemoveUserFromDept(dept, []system.User{*newUser})
		if err != nil {
			return tools.NewMySqlError(fmt.Errorf("在MySQL将用户从分组移除失败：" + err.Error()))
		}
	}
	return nil
}
