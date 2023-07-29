// Package factories 存放工厂方法
package factories

import (
	"github.com/go-faker/faker/v4"
	"github.com/golang-module/carbon/v2"
	"github.com/spf13/cast"
	"math/rand"
	"server/app/models"
	"server/app/models/user"
	"server/pkg/helpers"
)

func MakeUsers(times int) []user.User {

	var objs []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			LastName:  faker.LastName(),
			FirstName: faker.FirstName(),
			LastKana:  faker.LastName(),
			FirstKana: faker.FirstName(),
			// faker.Date() 生成的是 time.Time 类型，我们需要的是 carbon.Carbon 类型
			// 所以需要使用 carbon.Parse() 方法进行转换
			Birthday:     carbon.Date{Carbon: carbon.Parse(faker.Date())},
			Gender:       cast.ToUint8(rand.Intn(3) + 1),
			Email:        faker.Email(),
			Phone:        helpers.RandomNumber(11),
			Password:     "password",
			Introduction: faker.Paragraph(),
			CommonTimestampsField: models.CommonTimestampsField{
				State: cast.ToUint8(rand.Intn(100)),
				Order: cast.ToUint64(rand.Intn(1000)),
			},
		}
		objs = append(objs, model)
	}

	return objs
}
