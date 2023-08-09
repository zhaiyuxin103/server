package v1

import (
	"github.com/spf13/cast"
	"net/http"
	"server/app/http/controllers/api"
	"server/app/models"
	"server/app/models/file"
	"server/app/policies"
	"server/app/requests"
	"server/pkg/auth"
	filepkg "server/pkg/file"
	"server/pkg/response"
	"server/pkg/str"

	"github.com/gin-gonic/gin"
)

type FilesController struct {
	api.BaseController
}

func (ctrl *FilesController) Index(c *gin.Context) {
	files := file.All()
	response.Data(c, files)
}

func (ctrl *FilesController) Show(c *gin.Context) {
	fileModel := file.Get(c.Param("id"), true)
	if fileModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, fileModel)
}

func (ctrl *FilesController) Store(c *gin.Context) {

	request := requests.FileRequest{}
	if ok := requests.Validate(c, &request, requests.StoreFile); !ok {
		return
	}

	path, err := filepkg.SaveUploadImage(c, str.Plural(request.Folder))
	if err != nil {
		response.Abort500(c, "上传头像失败，请稍后尝试~")
		return
	}

	fileModel := file.File{
		Type:   request.Type,
		UserID: cast.ToUint64(auth.CurrentUID(c)),
		Folder: str.Plural(request.Folder),
		Path:   path,
		CommonTimestampsField: models.CommonTimestampsField{
			State: cast.ToUint8(request.State),
			Order: request.Order,
		},
	}
	fileModel.Create()
	if fileModel.ID > 0 {
		response.Created(c, fileModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *FilesController) Update(c *gin.Context) {

	fileModel := file.Get(c.Param("id"), false)
	if fileModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyFile(c, fileModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.FileRequest{}
	if ok := requests.Validate(c, &request, requests.UpdateFile); !ok {
		return
	}

	rowsAffected := fileModel.Save()
	if rowsAffected > 0 {
		response.Data(c, fileModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *FilesController) Delete(c *gin.Context) {

	fileModel := file.Get(c.Param("id"), false)
	if fileModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyFile(c, fileModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := fileModel.Delete()
	if rowsAffected > 0 {
		response.Success(c, http.StatusOK, "删除成功！", gin.H{})
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
