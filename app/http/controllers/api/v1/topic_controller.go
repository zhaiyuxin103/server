package v1

import (
	"github.com/spf13/cast"
	"server/app/http/controllers/api"
	"server/app/models"
	"server/app/models/topic"
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
