package v1

import (
	"github.com/spf13/cast"
	"server/app/http/controllers/api"
	"server/app/models"
	"server/app/models/category"
	"server/app/requests"
	"server/pkg/response"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	api.BaseController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.StoreCategory); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		ParentID:    request.ParentID,
		IsDirectory: request.IsDirectory,
		Description: request.Description,
		CommonTimestampsField: models.CommonTimestampsField{
			State: cast.ToUint8(request.State),
			Order: request.Order,
		},
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}
