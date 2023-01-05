package date
import(
	"fmt"
	"github.com/gin-gonic/gin"
	"Schedule/service/token"
	"Schedule/model"
	"Schedule/handler/user"
)
var Date model.Date

func SearchDate(c *gin.Context){
	Date.Year =c.PostForm("year")
	Date.Month=c.PostForm("month")
	Date.Day=c.PostForm("day")
	claim,_:=token.ParseToken(user.User.Token)
	Date.UserId=claim.ID
}//规定一个全局的date，后期查的时候专门从这里取出date

func Memory(c *gin.Context){
	year:=Date.Year
	month:=Date.Month
	day:=Date.Day
	claim,err:=token.ParseToken(user.User.Token)
	if err == nil{
		userid:=claim.ID
		model.Memory(year,month,day,userid)
		c.JSON(200,gin.H{
			"status":"ok",
		})
	}else{
		fmt.Printf("token err:%v",err)
		c.JSON(400,err)
	}
}