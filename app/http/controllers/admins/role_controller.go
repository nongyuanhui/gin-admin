package admins

import (
	"gin-admin/app/http"
	"gin-admin/app/services/admins"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	roleService = admins.RoleService{}
)

type RoleController struct {
	http.BaseController
}

// @Summary 角色列表
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |roleName|角色名称||
// @Tags  admin
// @version 1.0
// @Param page query int false "页码" default(1)
// @success 200 {object} utils.JSONResult{data=[]utils.Paginate}
// @Router /admin/v1/role/list [GET]
func (r *RoleController) GetRoleList(ctx *gin.Context) {
	page := 1
	if p, err := strconv.Atoi(ctx.DefaultQuery("page", "1")); err == nil {
		page = p
	}
	roleList := roleService.GetRoleList(page, 15)
	r.Success(ctx, roleList)
}

// @Summary 角色详情
// @Security ApiKeyAuth
// @Description | 参数 | 说明 |备注|
// @Description | :-----: | :----: | :----: |
// @Description |roleName|角色名称||
// @Description |permissionList|权限id||
// @Description |id|权限列表|permissionList列表下面的|
// @Description |apiPath|API地址||
// @Description |rule|匹配规则||
// @Description |method|请求方法||
// @Description |title|权限名称||
// @Description |parentId|parentId||
// @Description |isMenu|是否是菜单|1是0不是|
// @Description |path|路由地址||
// @Param roleId path int true "角色id" minimum(1)
// @Tags  admin
// @version 1.0
// @success 200 {object} utils.JSONResult{data=models.Role}
// @Router /admin/v1/role/detail/{roleId} [GET]
func (r *RoleController) GetRoleDetail(ctx *gin.Context) {
	form := struct {
		RoleId uint `json:"roleId" uri:"roleId" binding:"required,gte=1"`
	}{}
	if err := ctx.ShouldBindUri(&form); err != nil {
		lang := map[string]string{}
		lang["roleId"] = "角色id"
		err = r.Translate(err, lang)
		if err != nil {
			r.Error(ctx, err.Error())
			return
		}
		r.Error(ctx, "")
		return
	}
	role := roleService.GetRoleById(form.RoleId)
	if role.ID == 0 {
		r.Error(ctx, "角色不存在")
		return
	}
	rolePermissionList, err := roleService.GetRolePermission(role)
	if err != nil {
		r.Error(ctx, err.Error())
	}
	role.PermissionList = rolePermissionList
	r.Success(ctx, role)
}
