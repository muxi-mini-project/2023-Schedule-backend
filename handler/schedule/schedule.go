package schedule
import(
	"fmt"
	"time"
	//"strconv"
	"Schedule/service/token"
	"Schedule/handler/user"
	muser "Schedule/model/user"
	"Schedule/model"
	"Schedule/handler/date"
	"Schedule/service/datechange"
	"github.com/gin-gonic/gin"
)

// @Summary 写日程
// @Tag schedule
// @Accept application/json
// @Produce application/json
// @Param content formData string true "日程内容"
// @Success 200 {string} string "{"status":"ok"}"
// @Failure 400 {string} string "{"message":"token err"}"
// @Router /calendar/write [post]
func Write(c *gin.Context){
	content:=c.PostForm("schedule")
	//接下来把这一块写到model里面
	err:=muser.WriteSchedule(date.Date.Year,date.Date.Month,date.Date.Day,content,user.User.Token)
	if err!=nil{
		c.JSON(400,err)
		return
	}else{
		c.JSON(200,gin.H{
			"status":"ok",
		})
	}
}

// @Summary 完成某日程
// @Description 点击实现完成或取消完成
// @Tag schedule
// @Accept application/json
// @Produce application/json
// @Param content path string true "某个日程"
// @Success 200 {string} string "{"status","ok"}"
// @Failure 400 {string} string "{"token",err}"
// @Router /calendar/check/{content} [put]
func Check(c *gin.Context){
	claim,err:=token.ParseToken(user.User.Token)
	if err != nil{
		fmt.Printf("token err:%v",err)
		c.JSON(400,gin.H{
			"token":err,
		})
		return 
	}else{
		userid:=claim.ID
		content:=c.Param("content")
		model.Check(userid,content)
		c.JSON(200,gin.H{"status":"ok"})//done:=
	}//这里改成status ok即可
}//完成一个任务或者否认它的完成

// @Summary 查看某天的任务
// @Tag schedule
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Schedule "{"schedules":schedules}"
// @Success 200 {object} []model.Photo "{"photos":photos}"
// @Failure 400 {string} string "{"message":"token err"}"
// @Router /calendar [get]
func Calendar(c *gin.Context){
	claim,err:=token.ParseToken(user.User.Token)
	if err != nil{
		c.JSON(400,gin.H{
			"token":err,
		})
		return 
	}else{
		userid:=claim.ID
		schedules:=model.LookSchedule(date.Date.Year,date.Date.Month,date.Date.Day,userid)
		photos:=model.LookPhoto(date.Date.Year,date.Date.Month,date.Date.Day,userid)
		c.JSON(200,gin.H{
			"schedules":schedules,
			"photos":photos,
		})
	}
}

// @Summary 回收站
// @Description 里面有近14天已完成的任务
// @Tag schedule
// @Accept application/json
// @Produce application/json
// @Success 200 {object} []model.Schedule "{"scheduel":schedule}"
// @Failure 400 {string} string "{"token":err}"
// @Router /crash [get]
func Crash(c *gin.Context){
	year:=time.Now().Year()
	month:=datechange.MonthTrans(time.Now().Month().String())//这里要把January string转化为1 int
	day:=time.Now().Day()
	claim,err:=token.ParseToken(user.User.Token)
	if err != nil{
		c.JSON(400,gin.H{
			"token":err,
		})
		return 
	}else{
		userid:=claim.ID
		y,m,d:=year,month,day
		//我还得实现14天以内的备忘，14天得怎么实现？
		for i:=1;i<=14;i++{
			y,m,d=datechange.Yesterday(y,m,d)
			schedule:=model.LookScheduleFinish(y,m,d,userid)
			c.JSON(200,gin.H{
				"schedule":schedule,
			})
		}//循环14次，找到两周内做完的的任务
	}
}//垃圾桶里面装的是近两周内做完的任务

// //撕去一天的日历
// func Tear(c *gin.Context){
// 	year:=date.Date.Year
// 	month:=date.Date.Month
// 	day:=date.Date.Day
// 	claim,err:=token.ParseToken(user.User.Token)
// 	if err != nil{
// 		c.JSON(400,gin.H{
// 			"token":err,
// 		})
// 		return
// 	}else{
// 		userid:=claim.ID
// 		ok:=model.LookAllScheduleDone(year,month,day,userid)//要去dateplus里面找
// 		c.JSON(200,gin.H{
// 			"finish?":ok,
// 		})
// 	}
// }//允不允许撕下//丢给前端干