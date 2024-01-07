package controllers

import (
	"todo_list/models"
)

type List struct {
	ActionName string
	Method     string
	HttpMethod string
}

func NewList() *List {
	return &List{
		ActionName: "List",
		Method:     "list",
		HttpMethod: "GET",
	}
}

func (l *List) GetActionName() string {
	return l.GetActionName()
}

func (l *List) GetMethod() string {
	return l.Method
}

func (l *List) GetHttpMethod() string {
	return l.HttpMethod
}

func (l *List) Do(args []string) (map[string]any, map[string]any) {

	status := 200

	tasks := models.DB.Find([]string{})

	return map[string]any{
		"status": status,
		"data":   tasks,
	}, nil
}
