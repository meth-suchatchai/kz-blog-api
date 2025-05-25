package scenerepositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meth-suchatchai/kurostatemachine"
)

func (r *defaultRepository) UpdateSceneState(id uint, state kurostatemachine.State) *fiber.Error {
	err := r.orm.UpdateStatusScene(id, state)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return nil
}
