package sceneservices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kurostatemachine"
	scenemodels "github.com/meth-suchatchai/kz-blog-api/app/scene/models"
)

type Service interface {
	CreateScene(data scenemodels.CreateSceneData) (*scenemodels.Scene, *fiber.Error)
	UpdateSceneStatus(id uint, status kurostatemachine.State) *fiber.Error
}
