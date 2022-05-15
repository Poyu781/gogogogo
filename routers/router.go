package routers

import (
	"github.com/gin-gonic/gin"
	"github.trendmicro.com/jimmy-c-chiu/app-crypto-wallet/routers/api/v1"
	"github.trendmicro.com/jimmy-c-chiu/app-crypto-wallet/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")

	{
		apiv1.GET("/addresses", v1.GetAddress)
        apiv1.GET("/addresses/:address", v1.GetAddress)

        apiv1.POST("/addresses", v1.AddAddress)

        apiv1.DELETE("/addresses/:id", v1.DeleteAddress)
    }

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
