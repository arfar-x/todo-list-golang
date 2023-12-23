package runner

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

type Http struct {
	Factory
}

func (h *Http) GetName() string {
	return "HTTP Runner"
}

func (h *Http) Run([]string) {
	fmt.Println("HTTP Runner is running")
}
