package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "<h1>hello</h1>")
}
func main() {
	//http.HandleFunc("/hello", hello)
	//err := http.ListenAndServe(":8080", nil)
	//if err != nil {
	//	fmt.Printf("err:%+v\n", err)
	//	return
	//}
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	r.GET("/ping", func(c *gin.Context) {
		name, err := c.GetQuery("name")
		if !err {
			fmt.Printf("query lost:%+v\n", err)
		}
		c.JSON(200, gin.H{
			"name": name,
		})
	})
	r.Run()
}
