package runner

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"todo_list/controllers"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

type Http struct {
	Factory
	DB gorm.DB
}

func (h *Http) GetName() string {
	return "HTTP Runner"
}

func (h *Http) Run(args []string, db gorm.DB) {

	for _, action := range controllers.GetActions(db) {
		if _, ok := action.(controllers.BaseAction); ok {
			// TODO What to do for routes that receive `:id`s ?
			router.RouterGroup.Handle(action.GetHttpMethod(), "/"+action.GetMethod(), func(c *gin.Context) {
				result, err := action.Do(args)
				if err != nil {
					c.AbortWithStatusJSON(result["status"].(int), err)
				}
				c.JSON(result["status"].(int), result)
			})
		}
	}

	err := router.Run(os.Getenv("HTTP_IP") + ":" + os.Getenv("HTTP_PORT"))
	if err != nil {
		return
	}
}
