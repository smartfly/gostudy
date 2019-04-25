package main

import "github.com/gin-gonic/gin"

/**
 * @author taohu
 * @date 19-4-25
 * @time 下午7:12
 * @desc
 * To change this template use File | Settings | File Templates | Includes | File Header
**/

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
