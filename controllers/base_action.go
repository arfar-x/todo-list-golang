package controllers

type BaseAction interface {
	GetActionName() string
	GetMethod() string
	GetHttpMethod() string
	Do([]string) (map[string]any, map[string]any)
}

func GetActions() []BaseAction {
	return []BaseAction{
		NewList(),
	}
}
