package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/spf13/cast"
	"math/rand"
	"server/app/models"
	"server/app/models/category"
)

func MakeCategories(count int) []category.Category {

	var objs []category.Category

	// 设置唯一性，如 Category 模型的某个字段需要唯一，即可取消注释
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		categoryModel := category.Category{
			Name:        faker.Name(),
			ParentID:    0,
			IsDirectory: false,
			Level:       0,
			Path:        "-",
			Description: faker.Sentence(),
			CommonTimestampsField: models.CommonTimestampsField{
				State: cast.ToUint8(rand.Intn(100)),
				Order: cast.ToUint64(rand.Intn(1000)),
			},
		}
		objs = append(objs, categoryModel)
	}

	return objs
}
