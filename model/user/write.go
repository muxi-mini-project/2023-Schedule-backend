package user
import (
	"fmt"
	"time"
	"strconv"
	"Schedule/model"
	"Schedule/service/token"
)
func WriteSchedule(content string,tokenn string)(error){
	schedule:=model.Schedule{
		Content:	content,
		Done:		false,
	}
	schedule.Year=strconv.Itoa(time.Now().Year())
	schedule.Month=time.Now().Month().String()
	schedule.Day=strconv.Itoa(time.Now().Day())
	claim,err:=token.ParseToken(tokenn)
	if err == nil{
		schedule.UserId=claim.ID
		model.DB.Create(&schedule)
	}else{
		fmt.Printf("token err:%v",err)
	}
	return err
}
