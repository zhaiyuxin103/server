package v1

import (
	"server/app/http/controllers/api"
	"server/pkg/auth"
	"server/pkg/response"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	api.BaseController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}
