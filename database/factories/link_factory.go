package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/spf13/cast"
	"math/rand"
	"server/app/models"
	"server/app/models/link"
)

func MakeLinks(count int) []link.Link {

	var objs []link.Link

	// 设置唯一性，如 Link 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		linkModel := link.Link{
			Name:        faker.Username(),
			Description: faker.Paragraph(),
			URI:         faker.URL(),
			CommonTimestampsField: models.CommonTimestampsField{
				State: cast.ToUint8(rand.Intn(100)),
				Order: cast.ToUint64(rand.Intn(1000)),
			},
		}
		objs = append(objs, linkModel)
	}

	return objs
}
