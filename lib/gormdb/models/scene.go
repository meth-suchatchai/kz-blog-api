package dbmodels

import (
	"encoding/json"
	"github.com/meth-suchatchai/kurostatemachine"
	"gorm.io/gorm"
)

var SceneState = map[kurostatemachine.State][]kurostatemachine.State{
	"TODO":      {"READY", "DOING"},
	"READY":     {"DOING"},
	"DOING":     {"FAILED", "COMPLETED"},
	"FAILED":    {"DOING"},
	"COMPLETED": {},
}

var SM = kurostatemachine.NewStateMachine(SceneState)

type ApproveState string

const (
	APPROVE   ApproveState = "APPROVE"
	UNAPPROVE ApproveState = "UNAPPROVE"
)

func (as ApproveState) String() string {
	return string(as)
}

type Scene struct {
	gorm.Model
	MainCharacter string                 `gorm:"column:main_character;type:varchar(50)"`
	Description   string                 `gorm:"column:description;type:text"`
	Link          string                 `gorm:"column:link;type:text"`
	Medias        json.RawMessage        `gorm:"column:medias;type:jsonb;default([])"`
	Status        kurostatemachine.State `gorm:"column:status;type:varchar(10);default('TODO')"`
	ApproveStatus ApproveState           `gorm:"column:approve_status;type:varchar(10);default('UNAPPROVE')"`
}

func (m Scene) TableName() string {
	return "scenes"
}

func (m Scene) BeforeUpdate(tx *gorm.DB) error {
	var old Scene
	if err := tx.First(&old, m.ID).Error; err != nil {
		return err
	}

	item := &kurostatemachine.Item{
		TaskID: int(old.ID),
		State:  old.Status,
	}

	err := SM.Transition(item, m.Status)
	if err != nil {
		return err
	}

	if err = tx.Save(m).Error; err != nil {
		return err
	}

	return nil
}
