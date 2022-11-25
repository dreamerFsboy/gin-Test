package v_1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/astaxie/beego/validation"
	"github.com/unknwon/com"

	"github.com/dreamerFsboy/gin-Test/pkg/e"
	"github.com/dreamerFsboy/gin-blog/models"
	"github.com/dreamerFsboy/gin-blog/pkg/setting"
	"github.com/dreamerFsboy/gin-blog/pkg/util"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddTags(c *gin.Context) {

}

func EditTags(c *gin.Context) {

}

func DeleteTags(c *gin.Context) {

}
