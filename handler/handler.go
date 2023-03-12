package handler
import(
	"fmt"
	"time"
	//"Schedule/handler/user"
	"Schedule/model"
	"Schedule/service/datechange"
	//"Schedule/service/token"
	"Schedule/service/getId"
	"github.com/gin-gonic/gin"
)
// @Summary 首页
// @Description 登录前
// @Tags handler
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Fundmt
// @Router / [get]
func Index(c *gin.Context){
	fmt.Printf("Welcome!\n")
	c.JSON(200,gin.H{
		"code":200,
		"message":"Welcome!Click there to login!",
	})
}//登陆前界面

// @Summary 首页2
// @Description 登陆后
// @Tags handler
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Mem
// @Router /index [get]
func Index2(c *gin.Context){
	year:=time.Now().Year()
	month:=datechange.MonthTrans(time.Now().Month().String())
	day:=time.Now().Day()
	userid:=getId.GetId(c)
	ok:=model.IfMemory(year,month,day,userid)
	c.JSON(200,gin.H{
		"code":200,
		"message":"Welcome!",
		"IfMemory?":ok,
	})
}//登录后页面