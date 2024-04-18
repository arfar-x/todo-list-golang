package runner

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
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

	if err := h.ValidateCliArgs(args); err != nil {
		panic("Could not validate HTTP arguments.")
	}

	for _, originalAction := range controllers.GetActions(db) {
		// The closure `func(c *gin.Context)` captures the action variable from the lexical scope of the loop.
		// This means that when the closure is executed later (e.g., when the callback function is called),
		// it uses the current value of the captured action variable, which may have changed since
		// the closure was created.
		action := originalAction
		if _, ok := action.(controllers.BaseAction); ok {
			router.Handle(action.GetHttpMethod(), "/"+action.GetMethod(), func(c *gin.Context) {

				body := h.getJsonBody(c)

				result, err := action.Do(body)
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

func (_ *Http) ValidateCliArgs(args []string) error {
	if len(args) > 0 {
		panic("Http server cannot accept more than 1 CLI arguments.")
	}
	return error(nil)
}

func (_ *Http) getJsonBody(c *gin.Context) controllers.ArgumentBody {
	var body controllers.ArgumentBody
	data, _ := ioutil.ReadAll(c.Request.Body)

	if len(data) == 0 {
		return body
	}

	if err := json.Unmarshal(data, &body); err != nil {
		panic("Could not extract Json data from request body.")
	}

	if len(c.Params) != 0 {
		for _, param := range c.Params {
			body[param.Key] = param.Value
		}
	}

	return body
}
