package controllers

import "gorm.io/gorm"

type BaseAction interface {
	GetActionName() string
	GetMethod() string
	GetHttpMethod() string
	Do([]string) (map[string]any, map[string]any)
}

func GetActions(db gorm.DB) map[string]BaseAction {
	return map[string]BaseAction{
		"list": NewList(db),
	}
}
