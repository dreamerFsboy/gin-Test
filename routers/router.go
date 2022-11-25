package routers

import (
	v1 "gin-Test/routers/api/v1"

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
		apiv1.GET("/tags", v1.AddTag)
		apiv1.GET("/tags", v1.DeleteTags)
		apiv1.GET("/tags", v1.EditTags)
	}

	return r
}
