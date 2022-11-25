package routers

import (
	v1 "github.com/dreamerFsboy/gin-Test/routers/api/v_1"

	"github.com/dreamerFsboy/gin-Test/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTags)
		apiv1.PUT("/tags", v1.EditTags)
		apiv1.DELETE("/tags", v1.DeleteTags)
	}

	return r
}
