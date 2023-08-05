package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/spf13/cast"
	"math/rand"
	"server/app/models"
	"server/app/models/topic"
)

func MakeTopics(count int) []topic.Topic {

	var objs []topic.Topic

	// 设置唯一性，如 Topic 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		topicModel := topic.Topic{
			Title:           faker.Sentence(),
			SubTitle:        faker.Sentence(),
			Content:         faker.Paragraph(),
			CategoryID:      1,
			UserID:          1,
			VoteCount:       cast.ToUint64(rand.Intn(100)),
			UnvoteCount:     cast.ToUint64(rand.Intn(100)),
			ViewCount:       cast.ToUint64(rand.Intn(100)),
			FavoriteCount:   cast.ToUint64(rand.Intn(100)),
			ShareCount:      cast.ToUint64(rand.Intn(100)),
			ReplyCount:      cast.ToUint64(rand.Intn(100)),
			LastReplyUserID: 1,
			CommonTimestampsField: models.CommonTimestampsField{
				State: cast.ToUint8(rand.Intn(100)),
				Order: cast.ToUint64(rand.Intn(1000)),
			},
		}
		objs = append(objs, topicModel)
	}

	return objs
}
