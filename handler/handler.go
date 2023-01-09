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

func Index(c *gin.Context){
	fmt.Printf("Welcome!\n")
	c.JSON(200,gin.H{
		"status":"Welcome!Click there to login!",
	})
}//登陆前界面

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