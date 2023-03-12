package router
import (
	"Schedule/handler"
	"Schedule/handler/user"
	"Schedule/handler/schedule"
	"Schedule/handler/date"
	"Schedule/handler/door"
	"Schedule/handler/photo"
	"Schedule/handler/middleware"
	"github.com/gin-gonic/gin"
)
func Router(r *gin.Engine){
	//handler里面应该有user，schedule，photo，date和door
	r.GET("/api/v1",handler.Index)	//首页
	r.POST("/api/v1/login",user.Login)//登录
	g:=r.Group("/api/v1")//登陆后首页
	g.Use(middleware.TokenMiddle)
	g.Use(middleware.DateMiddle)
	{
		g.GET("/index",handler.Index2)
		g.GET("/preview/:year/:month/:day",door.Preview)//任意门
		g.GET("/crash",schedule.Crash)//垃圾桶
	}
	
	g1:=r.Group("/api/v1/calendar")//日历本
	g1.Use(middleware.TokenMiddle)
	g1.Use(middleware.DateMiddle)
	{
		g1.GET("",schedule.Calendar)//去往日历本
		g1.POST("/write",schedule.Write)//写计划
		g1.PUT("/check/:content",schedule.Check)//完成计划
		g1.POST("/addPhoto",photo.AddPhoto)//photo等我搞懂了图床再回来写
		g1.POST("/addPhoto2",photo.AddPhoto2)
		g1.POST("/addPhoto3",photo.AddPhoto3)
		g1.POST("/addPhoto4",photo.AddPhoto4)
		g1.PUT("/memory",date.Memory)//值得纪念的一天
		//g1.POST("/tear",schedule.Tear)//撕下来//丢给前端
	}
	g2:=r.Group("/api/v1/board")
	g2.Use(middleware.TokenMiddle)
	{
		g2.POST("/searchdate",date.SearchDate)//查找时间
		g2.Use(middleware.DateMiddle)
		g2.GET("",date.PhotoAndSchedule)//背景板//得去找那天的图片和未做完的任务 not ok
		
	}
}//这里的路由得升级一下，改写成路由组

/*还有什么要写：
*/