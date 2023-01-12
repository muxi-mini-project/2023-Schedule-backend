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

// @Summary 查日期
// @Description 输入年月日
// @Tag date
// @Accept application/json
// @Produce application/json
// @Param year formData string true "年"
// @Param month formData string true "月"
// @Param day formData string true "日"
// @Success 200 {string} string "{"status":"ok"}"
// @Failure 400 {string} string "{"token err":err}"
// @Router /board/searchdate [post]
func SearchDate(c *gin.Context){
	Date.Year,_=strconv.Atoi(c.PostForm("year"))
	Date.Month,_=strconv.Atoi(c.PostForm("month"))
	Date.Day,_=strconv.Atoi(c.PostForm("day"))
	claim,err:=token.ParseToken(user.User.Token)
	if err!= nil{
		c.JSON(400,gin.H{
			"token err":err,
		})
	}else{
		Date.UserId=claim.ID
		c.JSON(200,gin.H{
			"status":"ok",
		})
	}
	
}//规定一个全局的date，后期查的时候专门从这里取出date

// @Summary 纪念某一天
// @Tag date
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"status":"ok"}"
// @Failure 400 {string} string "{"token",err}"
// @Router /calendar/memory [put]
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
// @Summary 查看某一天的未完成计划和图片
// @Tag date
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Photo "{"photos":photo}"
// @Success 200 {object} []model.Schedule "{"unfinishedSchedule":unfinishedSchedule}"
// @Failure 400 {string} string "{"token":err}"
// @Router /board [put]
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