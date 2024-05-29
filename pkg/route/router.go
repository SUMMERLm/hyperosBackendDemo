package route

// 路由包

import (
	"net/http"

	"github.com/SUMMERLm/serverless/pkg/api"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			//c.Header("Access-Control-Allow-Origin", "http://localhost:63344") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "http://localhost:63344")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, SCNID, HOST")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,"+
				" Cache-Control, Content-Language, Content-Type ")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(Cors())
	// get Test
	router.GET("/", api.GetIndex)
	// post Test
	router.POST("/test", api.PostTest)

	return router
}
