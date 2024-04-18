package controllers

import (
	"gorm.io/gorm"
	"strconv"
	"todo_list/models"
)

type Update struct {
	ActionName string
	Method     string
	HttpMethod string
	DB         gorm.DB
}

func NewUpdate(db gorm.DB) *Update {
	return &Update{
		ActionName: "Update",
		Method:     "list/:id",
		HttpMethod: "PUT",
		DB:         db,
	}
}

func (c *Update) GetActionName() string {
	return c.ActionName
}

func (c *Update) GetMethod() string {
	return c.Method
}

func (c *Update) GetHttpMethod() string {
	return c.HttpMethod
}

func (c *Update) Do(args ArgumentBody) (map[string]any, map[string]any) {

	done, ok := args["done"].(bool)
	if !ok {
		done, _ = strconv.ParseBool(args["done"].(string))
	}

	id, _ := strconv.ParseUint(args["id"].(string), 10, 8)

	task := models.Task{
		ID:   uint(id),
		Name: args["name"].(string),
		Done: done,
	}

	result := c.DB.Model(task).Updates(task)

	if result.Error != nil {
		return map[string]any{
			"status": 500,
			"data":   "Could not update the task",
		}, nil
	}

	return map[string]any{
		"status": 200,
		"data":   task,
	}, nil
}
