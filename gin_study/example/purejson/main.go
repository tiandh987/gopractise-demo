package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 通常，JSON 使用 unicode 替换特殊 HTML 字符，
// 例如 < 变为 \ u003c
// 如果要按字面对这些字符进行编码，则可以使用 PureJSON。

func main()  {
	r := gin.Default()

	// 提供 unicode 实体
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// 提供字面字符
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world</b>",
		})
	})

	r.Run(":8080")
}