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

func (ctrl *CategoriesController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := category.Paginate(c, 10)
	response.Data(c, gin.H{
		"data":  data,
		"pager": pager,
	})
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

func (ctrl *CategoriesController) Update(c *gin.Context) {

	// 验证 url 参数 id 是否正确
	categoryModel := category.Get(c.Param("id"), false)
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	// 表单验证
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.UpdateCategory); !ok {
		return
	}

	// 保存数据
	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	categoryModel.CommonTimestampsField.State = cast.ToUint8(request.State)
	categoryModel.CommonTimestampsField.Order = request.Order
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c)
	}
}
