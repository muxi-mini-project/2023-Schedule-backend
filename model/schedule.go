package model
import(
	//"Schedule/service/datechange"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func LookSchedule(year int,month int,day int,
		uid string)([]Schedule){
	var schedule []Schedule
	DB.Where("year=? and month=? and day=? and user_id = ?",year,month,day,uid).Find(&schedule)
	return schedule
}//查看该用户在某天的所有任务

func Check(userid string,content string)(bool){
	var schedule Schedule
	DB.Where("user_id=? and content=?",userid,content).Find(&schedule)
	if schedule.Done == false{
		DB.Model(&schedule).Update(gin.H{"Done":true})
	}else {
		DB.Model(&schedule).Update(gin.H{"Done":false})
	}
	return schedule.Done
}//完成某一任务

func LookScheduleFinish(year int,month int,day int,uid string)([]Schedule){
	var schedule []Schedule
	DB.Where("year=? and month=? and day=? and user_id = ? and Done = ?",year,month,day,uid,true).Find(&schedule)
	return schedule
}//已做的任务
//所以要在hanler里面进行循环，model只负责查找

func LookScheduleUnfinish(year int,month int,day int,uid string)([]Schedule){
	var schedule []Schedule
	DB.Where("year=? and month=? and day=? and user_id = ? and Done = ?",year,month,day,uid,false).Find(&schedule)
	return schedule
}//还没做的任务