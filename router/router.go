package router
import (
	"Schedule/handler"
	"Schedule/handler/user"
	"Schedule/handler/schedule"
	"Schedule/handler/date"
	"Schedule/handler/door"
	"github.com/gin-gonic/gin"
)
func Router(r *gin.Engine){
	r.GET("/",handler.Index)

	r.POST("/login",user.Login)
	r.POST("/write",user.Write)
	r.POST("/check",user.Check)
	r.POST("/tear",user.Tear)

	r.POST("/searchdate",schedule.SearchDate)

	r.POST("/show",date.Show)
	r.POST("/insert",date.Insert)

	r.POST("/preview",door.Preview)
}