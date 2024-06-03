package main

import (
	"fmt"
)
import "github.com/gin-gonic/gin"

var myMiddleWare = func(ctx *gin.Context) {
	fmt.Println("my middleware is running...")
}
var myMiddleWareGroup = func(ctx *gin.Context) {
	fmt.Println("my group middleware is running...")
}

func main() {
	engine := gin.Default()

	engine.Use(myMiddleWare)
	group := engine.Group(`/utest`)
	group.Use(myMiddleWareGroup)

	group.POST("/test", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("pong"))
	})

	engine.POST("/ping", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("pong"))
	})
	err := engine.Run(":8080")
	if err != nil {
		return
	}

}
