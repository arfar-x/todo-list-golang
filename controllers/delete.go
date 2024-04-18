package controllers

import (
	"gorm.io/gorm"
	"strconv"
	"todo_list/models"
)

type Delete struct {
	ActionName string
	Method     string
	HttpMethod string
	DB         gorm.DB
}

func NewDelete(db gorm.DB) *Delete {
	return &Delete{
		ActionName: "Delete",
		Method:     "list/:id",
		HttpMethod: "DELETE",
		DB:         db,
	}
}

func (c *Delete) GetActionName() string {
	return c.ActionName
}

func (c *Delete) GetMethod() string {
	return c.Method
}

func (c *Delete) GetHttpMethod() string {
	return c.HttpMethod
}

func (c *Delete) Do(args ArgumentBody) (map[string]any, map[string]any) {

	id, _ := strconv.ParseUint(args["id"].(string), 10, 8)

	task := models.Task{ID: uint(id)}

	result := c.DB.Delete(&task)

	if result.Error != nil {
		return map[string]any{
			"status": 500,
			"data":   "Could not delete the task",
		}, nil
	}

	return map[string]any{
		"status": 204,
		"data":   make(map[string]any),
	}, nil
}
