package router
import (
	"Schedule/handler"
	"Schedule/handler/user"
	//"Schedule/handler/schedule"
	"Schedule/handler/date"
	"Schedule/handler/door"
	"github.com/gin-gonic/gin"
)
func Router(r *gin.Engine){
	//handler里面应该有user，schedule，photo，date和door
	r.GET("/",handler.Index)

	r.POST("/login",user.Login)

	r.POST("/write",user.Write)//schedule
	r.PUT("/check/:content",user.Check)//schedule
	r.POST("/tear",user.Tear)

	//r.POST("/addPhoto",photo.addPhoto)//photo等我搞懂了图床再回来写
	
	r.POST("/searchdate",date.SearchDate)//查找时间
	r.PUT("/memory",date.Memory)

	r.POST("/preview",door.Preview)
}

/*还有什么要写：
2.config文件的配置
3.go routine（实现定期刷新）*/