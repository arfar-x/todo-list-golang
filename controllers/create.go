package controllers

import (
	"gorm.io/gorm"
	"strconv"
	"todo_list/models"
)

type Create struct {
	ActionName string
	Method     string
	HttpMethod string
	DB         gorm.DB
}

func NewCreate(db gorm.DB) *Create {
	return &Create{
		ActionName: "Create",
		Method:     "list",
		HttpMethod: "POST",
		DB:         db,
	}
}

func (c *Create) GetActionName() string {
	return c.ActionName
}

func (c *Create) GetMethod() string {
	return c.Method
}

func (c *Create) GetHttpMethod() string {
	return c.HttpMethod
}

func (c *Create) Do(args ArgumentBody) (map[string]any, map[string]any) {

	var done bool
	if _, ok := args["done"].(bool); !ok {
		done, _ = strconv.ParseBool(args["done"].(string))
	}

	task := models.Task{
		Name: args["name"].(string),
		Done: done,
	}

	result := c.DB.Create(&task)

	if result.Error != nil {
		return map[string]any{
			"status": 500,
			"data":   "Could not create the task",
		}, nil
	}

	return map[string]any{
		"status": 201,
		"data":   task,
	}, nil
}
