package user
import (
	"fmt"
	"Schedule/model"
	"Schedule/service/token"
)
func WriteSchedule(year int,month int,day int,
		content string,tokenn string)(error){
	schedule:=model.Schedule{
		Content:	content,
		Done:		false,
	}
	schedule.Year=year
	schedule.Month=month
	schedule.Day=day
	claim,err:=token.ParseToken(tokenn)
	if err == nil{
		schedule.UserId=claim.ID
		model.DB.Create(&schedule)
	}else{
		fmt.Printf("token err:%v",err)
	}
	return err
}
