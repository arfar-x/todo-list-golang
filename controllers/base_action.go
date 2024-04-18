package controllers

import (
	"gorm.io/gorm"
)

type BaseAction interface {
	GetActionName() string
	GetMethod() string
	GetHttpMethod() string
	Do(ArgumentBody) (map[string]any, map[string]any)
}

// ArgumentBody The type used to pass arguments through each server runner.
type ArgumentBody map[string]any

func GetActions(db gorm.DB) map[string]BaseAction {
	return map[string]BaseAction{
		"list":   NewList(db),
		"create": NewCreate(db),
		"update": NewUpdate(db),
		"delete": NewDelete(db),
	}
}
