package handler
import(
	"fmt"
	"time"
	"Schedule/handler/user"
	"Schedule/model"
	"Schedule/service/datechange"
	"Schedule/service/token"
	"github.com/gin-gonic/gin"
)
// @Summary 首页
// @Description 登录前
// @Tag handler
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"status":"Welcome!Click there to login!"}"
// @Router / [get]
func Index(c *gin.Context){
	fmt.Printf("Welcome!\n")
	c.JSON(200,gin.H{
		"status":"Welcome!Click there to login!",
	})
}//登陆前界面

// @Summary 首页2
// @Description 登陆后
// @Tag handler
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"status":"Welcome!"}"
// @Success 200 {boolean} bool "{"IfMemory?":ok}"
// @Router /index [get]
func Index2(c *gin.Context){
	c.JSON(200,gin.H{
		"status":"Welcome!",
	})
	year:=time.Now().Year()
	month:=datechange.MonthTrans(time.Now().Month().String())
	day:=time.Now().Day()
	claim,_:=token.ParseToken(user.User.Token)
	userid:=claim.ID
	ok:=model.IfMemory(year,month,day,userid)
	c.JSON(200,gin.H{
		"IfMemory?":ok,
	})
}//登录后页面