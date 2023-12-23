package controllers

type BaseAction interface {
	GetActionName() string
	GetMethod() string
	GetHttpMethod() string
}
