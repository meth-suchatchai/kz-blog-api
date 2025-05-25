package sceneservices

import (
	"github.com/gofiber/fiber/v2"
	scenemodels "github.com/meth-suchatchai/kz-blog-api/app/scene/models"
)

func (svc *defaultService) CreateScene(data scenemodels.CreateSceneData) (*scenemodels.Scene, *fiber.Error) {
	return svc.rp.CreateScene(data)
}
