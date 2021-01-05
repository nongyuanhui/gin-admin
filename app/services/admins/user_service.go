package admins

import (
	"gin-admin/app/models"
	"gin-admin/db"
	"gin-admin/pkg/utils"
)

type UserService struct {
}

func (userService *UserService) GetUserList(page, limit int) *utils.Paginate {
	var userList []models.User
	var count int64
	db.DB.Offset((page - 1) * limit).Limit(limit).Find(&userList).Offset(0).Count(&count)
	return utils.NewPaginate(count, page, limit, userList)
}
