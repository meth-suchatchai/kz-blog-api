package sceneservices

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kurostatemachine"
)

func (svc *defaultService) UpdateSceneStatus(id uint, status kurostatemachine.State) *fiber.Error {
	return svc.rp.UpdateSceneState(id, status)
}
