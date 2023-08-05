package v1

import (
	"github.com/spf13/cast"
	"server/app/http/controllers/api"
	"server/app/models"
	"server/app/models/topic"
	"server/app/policies"
	"server/app/requests"
	"server/pkg/auth"
	"server/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	api.BaseController
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.StoreTopic); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		SubTitle:   request.SubTitle,
		Content:    request.Content,
		CategoryID: cast.ToUint64(request.CategoryID),
		UserID:     cast.ToUint64(auth.CurrentUID(c)),
		CommonTimestampsField: models.CommonTimestampsField{
			State: cast.ToUint8(request.State),
			Order: request.Order,
		},
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Update(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"), false)
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.UpdateTopic); !ok {
		return
	}

	if request.Title != "" {
		topicModel.Title = request.Title
	}
	if request.SubTitle != "" {
		topicModel.SubTitle = request.SubTitle
	}
	if request.CategoryID != "" {
		topicModel.CategoryID = cast.ToUint64(request.CategoryID)
	}
	if request.Content != "" {
		topicModel.Content = request.Content
	}
	rowsAffected := topicModel.Save()
	if rowsAffected > 0 {
		response.Data(c, topicModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Delete(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"), false)
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := topicModel.Delete()
	if rowsAffected > 0 {
		response.Ok(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
