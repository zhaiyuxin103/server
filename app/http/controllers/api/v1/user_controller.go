package v1

import (
	"github.com/golang-module/carbon/v2"
	"server/app/http/controllers/api"
	"server/app/models/user"
	"server/app/requests"
	"server/pkg/auth"
	"server/pkg/helpers"
	"server/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	api.BaseController
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

func (ctrl *UsersController) Update(c *gin.Context) {

	request := requests.UserRequest{}
	if ok := requests.Validate(c, &request, requests.UpdateUser); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	if !helpers.Empty(request.LastName) {
		currentUser.LastName = request.LastName
	}
	if !helpers.Empty(request.FirstName) {
		currentUser.FirstName = request.FirstName
	}
	if !helpers.Empty(request.LastKana) {
		currentUser.LastKana = request.LastKana
	}
	if !helpers.Empty(request.FirstKana) {
		currentUser.FirstKana = request.FirstKana
	}
	if !helpers.Empty(request.Birthday) {
		currentUser.Birthday = carbon.Date{Carbon: carbon.Time2Carbon(request.Birthday).SetTimezone(carbon.PRC)}
	}
	if !helpers.Empty(request.AvatarID) {
		currentUser.AvatarID = request.AvatarID
	}
	if !helpers.Empty(request.Gender) {
		currentUser.Gender = request.Gender
	}
	if !helpers.Empty(request.Phone) {
		currentUser.Phone = request.Phone
	}
	if !helpers.Empty(request.Province) {
		currentUser.Province = request.Province
	}
	if !helpers.Empty(request.City) {
		currentUser.City = request.City
	}
	if !helpers.Empty(request.District) {
		currentUser.District = request.District
	}
	if !helpers.Empty(request.Address) {
		currentUser.Address = request.Address
	}
	if !helpers.Empty(request.Introduction) {
		currentUser.Introduction = request.Introduction
	}
	if _, ok := c.GetPostForm("state"); ok {
		currentUser.State = request.State
	}
	if _, ok := c.GetPostForm("order"); ok {
		currentUser.Order = request.Order
	}
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdateEmail(c *gin.Context) {

	request := requests.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, requests.UpdateUserEmail); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Email = request.Email
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Ok(c)
	} else {
		// 失败，显示错误提示
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePhone(c *gin.Context) {

	request := requests.UserUpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePhone); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Phone = request.Phone
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Ok(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}
