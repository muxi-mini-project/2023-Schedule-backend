package router
import (
	"Schedule/handler"
	"Schedule/handler/user"
	"Schedule/handler/schedule"
	"Schedule/handler/date"
	"Schedule/handler/door"
	"Schedule/handler/photo"
	"github.com/gin-gonic/gin"
)
func Router(r *gin.Engine){
	//handler里面应该有user，schedule，photo，date和door
	r.GET("",handler.Index)
	//首页
	r.GET("/index",handler.Index2)
	
	r.GET("/preview/:year/:month/:day",door.Preview)//任意门
	
	r.POST("/login",user.Login)

	g1:=r.Group("/calender")//日历本
		g1.GET("/",schedule.Calendar)//去往日历本
		g1.POST("/write",schedule.Write)//写计划
		g1.PUT("/check/:content",schedule.Check)//完成计划
		g1.POST("/addPhoto",photo.AddPhoto)//photo等我搞懂了图床再回来写
		g1.PUT("/memory",date.Memory)//值得纪念的一天
		//g1.POST("/tear",schedule.Tear)//撕下来//丢给前端
	
	g2:=r.Group("/board")
		g2.GET("/",date.PhotoAndSchedule)//背景板//得去找那天的图片和未做完的任务 not ok
		g2.POST("/searchdate",date.SearchDate)//查找时间
	
	r.GET("/crash",schedule.Crash)//垃圾桶
}//这里的路由得升级一下，改写成路由组

/*还有什么要写：
*/