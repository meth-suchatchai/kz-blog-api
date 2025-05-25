package scenerepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kurostatemachine"
	scenemodels "github.com/meth-suchatchai/kz-blog-api/app/scene/models"
)

type Repository interface {
	CreateScene(data scenemodels.CreateSceneData) (*scenemodels.Scene, *fiber.Error)
	UpdateSceneState(id uint, state kurostatemachine.State) *fiber.Error
}
