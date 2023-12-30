package runner

import (
	"os"
	"todo_list/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	err := godotenv.Load(".env")
	if err != nil {
		panic("Env variables could not be injected.")
	}
}

type Http struct {
	Factory
}

func (h *Http) GetName() string {
	return "HTTP Runner"
}

func (h *Http) Run(args []string) {

	for _, action := range controllers.GetActions() {
		if _, ok := action.(controllers.BaseAction); ok {
			// TODO What to do for routes that receive `:id`s ?
			router.RouterGroup.Handle("/"+action.GetHttpMethod(), action.GetMethod(), func(c *gin.Context) {
				result, err := action.Do(args)
				if err != nil {
					c.AbortWithStatusJSON(err["status"].(int), err)
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
