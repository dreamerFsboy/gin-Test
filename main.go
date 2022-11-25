package main

import (
	"fmt"
	"gin-Test/routers"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dreamerFsboy/gin-Test/pkg/setting"
	"github.com/dreamerFsboy/gin-Test/routers"
)

func main() {
	router := routers.InitRouter()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "test",
		})
	})

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
