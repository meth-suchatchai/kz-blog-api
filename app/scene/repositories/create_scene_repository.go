package scenerepositories

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"

	scenemodels "github.com/meth-suchatchai/kz-blog-api/app/scene/models"
	dbmodels "github.com/meth-suchatchai/kz-blog-api/lib/gormdb/models"
)

func (r *defaultRepository) CreateScene(data scenemodels.CreateSceneData) (*scenemodels.Scene, *fiber.Error) {

	dataMedias, err := json.Marshal(data.Medias)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	dao, err := r.orm.CreateScene(&dbmodels.Scene{
		MainCharacter: data.MainCharacter,
		Description:   data.Description,
		Medias:        dataMedias,
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	var medias []map[string]interface{}
	b, err := dao.Medias.MarshalJSON()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	err = json.Unmarshal(b, &medias)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return &scenemodels.Scene{
		MainCharacter: dao.MainCharacter,
		Description:   dao.Description,
		Link:          dao.Link,
		Medias:        medias,
		Status:        dao.Status,
		ApproveStatus: dao.ApproveStatus.String(),
	}, nil
}
