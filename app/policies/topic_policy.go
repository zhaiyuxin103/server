// Package policies 用户授权
package policies

import (
	"github.com/spf13/cast"
	"server/app/models/topic"
	"server/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return cast.ToUint64(auth.CurrentUID(c)) == _topic.UserID
}
