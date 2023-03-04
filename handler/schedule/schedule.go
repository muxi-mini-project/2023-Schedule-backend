package schedule
import(
	//"fmt"
	"time"
	//"strconv"
	//"Schedule/service/token"
	//"Schedule/handler/user"
	muser "Schedule/model/user"
	"Schedule/model"
	//"Schedule/handler/date"
	"Schedule/service/datechange"
	"Schedule/service/getId"
	//"Schedule/service/setDate"
	"Schedule/service/setDate"
	"github.com/gin-gonic/gin"
)

// @Summary 写日程
// @Tags schedule
// @Accept application/json
// @Produce application/json
// @Param content formData string true "日程内容"
// @Success 200 {object} model.Fundmt
// @Failure 400 {object} model.Fundmt
// @Router /calendar/write [post]
func Write(c *gin.Context){
	content:=c.PostForm("schedule")
	//接下来把这一块写到model里面
	userid:=getId.GetId(c)
	y,m,d:=setDate.GetDate(c)
	err:=muser.WriteSchedule(y,m,d,content,userid)
	if err!=nil{
		c.JSON(400,gin.H{
			"code":400,
			"message":"创建失败，请重试",
		})
		return
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"message":"创建记录成功",
		})
	}
}

// @Summary 完成某日程
// @Description 点击实现完成或取消完成
// @Tags schedule
// @Accept application/json
// @Produce application/json
// @Param content path string true "某个日程"
// @Success 200 {object} model.Fundmt
// @Router /calendar/check/{content} [put]
func Check(c *gin.Context){
	userid:=getId.GetId(c)
	content:=c.Param("content")
	model.Check(userid,content)
	c.JSON(200,gin.H{
		"code":200,
		"message":"ok",
	})//done:=
//这里改成status ok即可
}//完成一个任务或者否认它的完成

// @Summary 查看某天的任务
// @Tags schedule
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.SchAndPh
// @Router /calendar [get]
func Calendar(c *gin.Context){
	userid:=getId.GetId(c)
	y,m,d:=setDate.GetDate(c)
	schedules:=model.LookSchedule(y,m,d,userid)
	photos:=model.LookPhoto(y,m,d,userid)
	c.JSON(200,gin.H{
		"code":200,
		"message":"ok",
		"schedule":schedules,
		"photo":photos,
	})
}

// @Summary 回收站
// @Description 里面有近14天已完成的任务
// @Tags schedule
// @Accept application/json
// @Produce application/json
// @Success 200 {object} model.Sch
// @Router /crash [get]
func Crash(c *gin.Context){
	year:=time.Now().Year()
	month:=datechange.MonthTrans(time.Now().Month().String())//这里要把January string转化为1 int
	day:=time.Now().Day()
	userid:=getId.GetId(c)
	y,m,d:=year,month,day
	//我还得实现14天以内的备忘，14天得怎么实现？
	for i:=1;i<=14;i++{
		y,m,d=datechange.Yesterday(y,m,d)
		schedule:=model.LookScheduleFinish(y,m,d,userid)
		c.JSON(200,gin.H{
			"code":200,
			"message":"ok",
			"schedule":schedule,
		})
	}//循环14次，找到两周内做完的的任务
}//垃圾桶里面装的是近两周内做完的任务