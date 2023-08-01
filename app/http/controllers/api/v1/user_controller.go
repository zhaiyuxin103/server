package v1

import (
	"server/app/http/controllers/api"
	"server/app/models/user"
	"server/app/requests"
	"server/pkg/auth"
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
