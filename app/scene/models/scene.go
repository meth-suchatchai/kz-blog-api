package scenemodels

import (
	"github.com/meth-suchatchai/kurostatemachine"
	"time"
)

type CreateSceneRequest struct {
	CreateSceneData
}

type CreateSceneResponse struct {
	Scene
}

type UpdateSceneStatusRequest struct {
	Status kurostatemachine.State `form:"status" json:"status" binding:"required"`
}
type UpdateSceneResponse struct{}
type CreateSceneData struct {
	MainCharacter string                   `json:"main_character"`
	Description   string                   `json:"description"`
	Medias        []map[string]interface{} `json:"medias"`
}

type Scene struct {
	ID            uint                     `json:"id"`
	CreatedAt     time.Time                `json:"created_at"`
	UpdatedAt     time.Time                `json:"updated_at"`
	MainCharacter string                   `json:"main_character"`
	Description   string                   `json:"description"`
	Link          string                   `json:"link"`
	Medias        []map[string]interface{} `json:"medias"`
	Status        kurostatemachine.State   `json:"status"`
	ApproveStatus string                   `json:"approve_status"`
}
