package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetStudent(name string) string {
	return name
}
func GetName(ctx *gin.Context) {
	param := ctx.Query("name")
	if len(param) == 0 {
		ctx.String(http.StatusBadRequest, "lose param name")
		return
	}
	name := GetStudent(param)
	ctx.String(http.StatusOK, name)
	return
}
func main() {
	engine := gin.Default()
	engine.GET("/getName", GetName)
	err := engine.Run("localhost:2345")
	if err != nil {
		panic(err)
	}
}
