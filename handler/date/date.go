package date
import(
	//"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	//"Schedule/service/token"
	"Schedule/service/getId"
	"Schedule/service/setDate"
	"Schedule/model"
	//"Schedule/handler/user"
)
//var Date model.Date//全局变量date

// @Summary 查日期
// @Description 输入年月日
// @Tags date
// @Accept application/json
// @Produce application/json
// @Param year formData string true "年"
// @Param month formData string true "月"
// @Param day formData string true "日"
// @Success 200 {object} model.SetDate
// @Router /board/searchdate [post]
func SearchDate(c *gin.Context){
	year,_:=strconv.Atoi(c.PostForm("year"))
	month,_:=strconv.Atoi(c.PostForm("month"))
	day,_:=strconv.Atoi(c.PostForm("day"))
	//claim,err:=token.ParseToken(user.User.Token)
	//userid,_:=c.Get("UserId")
	//Date.UserId=getId.GetId(c)
	c.JSON(200,gin.H{
		"code":200,
		"message":"ok",
		"year":year,
		"month":month,
		"day":day,
	})//之后放在header里面取出来就好
	
}//规定一个全局的date，后期查的时候专门从这里取出date//老旧的想法，狠狠的嘲讽过去的我

// @Summary 纪念某一天
// @Tags date
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Fundmt
// @Router /calendar/memory [put]
func Memory(c *gin.Context){
	year,month,day:=setDate.GetDate(c)
	userid:=getId.GetId(c)
	model.Memory(year,month,day,userid)
	c.JSON(200,gin.H{
		"code":200,
		"message":"ok",
	})
}

// @Summary 查看某一天的未完成计划和图片
// @Tags date
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.SchAndPh
// @Router /board [put]
func PhotoAndSchedule(c *gin.Context){
	y,m,d:=setDate.GetDate(c)
	uid:=getId.GetId(c)
	photo:=model.LookPhoto(y,m,d,uid)
	unfinishedSchedule:=model.LookScheduleUnfinish(y,m,d,uid)
	c.JSON(200,gin.H{
		"code":200,
		"message":"ok",
		"schedule":unfinishedSchedule,
		"photos":photo,
	})
}