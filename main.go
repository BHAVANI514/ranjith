package main

import (
	"fmt"
	"time"
)

func main() {
	    router := gin.Default()
	router.GET("/",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H {
			"meassage" : "welocome to gin API!"	
		})	
	})

    // Start the server on port 8081
    router.Run(":8081")
}
