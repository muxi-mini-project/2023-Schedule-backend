package user
import (
	//"fmt"
	"Schedule/model"
	//"Schedule/service/token"
)
func WriteSchedule(year int,month int,day int,
		content string,uid string)(error){
	schedule:=model.Schedule{
		Content:	content,
		Done:		false,
	}
	schedule.Year=year
	schedule.Month=month
	schedule.Day=day
	schedule.UserId=uid
	err:=model.DB.Create(&schedule).Error
	return err
}
