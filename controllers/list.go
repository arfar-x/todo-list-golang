package controllers

import (
	"gorm.io/gorm"
	"todo_list/models"
)

type List struct {
	ActionName string
	Method     string
	HttpMethod string
	DB         gorm.DB
}

func NewList(db gorm.DB) *List {
	return &List{
		ActionName: "List",
		Method:     "list",
		HttpMethod: "GET",
		DB:         db,
	}
}

func (l *List) GetActionName() string {
	return l.ActionName
}

func (l *List) GetMethod() string {
	return l.Method
}

func (l *List) GetHttpMethod() string {
	return l.HttpMethod
}

func (l *List) Do(_ ArgumentBody) (map[string]any, map[string]any) {
	status := 200

	var tasks []models.Task

	l.DB.Find(&tasks)

	return map[string]any{
		"status": status,
		"data":   tasks,
	}, nil
}
