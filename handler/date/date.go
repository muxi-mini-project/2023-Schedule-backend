package date
import(
	"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"Schedule/service/token"
	"Schedule/model"
	"Schedule/handler/user"
)
var Date model.Date//全局变量date

func SearchDate(c *gin.Context){
	Date.Year,_=strconv.Atoi(c.PostForm("year"))
	Date.Month,_=strconv.Atoi(c.PostForm("month"))
	Date.Day,_=strconv.Atoi(c.PostForm("day"))
	claim,_:=token.ParseToken(user.User.Token)
	Date.UserId=claim.ID
}//规定一个全局的date，后期查的时候专门从这里取出date

func Memory(c *gin.Context){
	year:=Date.Year
	month:=Date.Month
	day:=Date.Day
	claim,err:=token.ParseToken(user.User.Token)
	if err != nil{
		fmt.Printf("token err:%v",err)
		c.JSON(400,err)
		return 
	}else{
		userid:=claim.ID
		model.Memory(year,month,day,userid)
		c.JSON(200,gin.H{
			"status":"ok",
		})
	}
}

func PhotoAndSchedule(c *gin.Context){
	y:=Date.Year
	m:=Date.Month
	d:=Date.Day
	claim,err:=token.ParseToken(user.User.Token)
	if err != nil{
		fmt.Printf("token err:%v",err)
		c.JSON(400,err)
		return 
	}else{
		uid:=claim.ID
		photo:=model.LookPhoto(y,m,d,uid)
		unfinishedSchedule:=model.LookScheduleUnfinish(y,m,d,uid)
		c.JSON(200,gin.H{
			"photos":photo,
			"unfinishedSchedule":unfinishedSchedule,
		})
	}
}