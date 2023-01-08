package model
import(
	"github.com/gin-gonic/gin"
)
func LookSchededule(year string,month string,day string, uid string)([]Schedule){
	var schedule []Schedule
	DB.Where("year=? and month=? and day=? and uid = ?",year,month,day,uid).Find(&schedule)
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

