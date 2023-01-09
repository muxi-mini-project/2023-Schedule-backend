package punch
import(
	"time"
	"Schedule/service/token"
	"Schedule/handler/user"
	"Schedule/service/datechange"
	//"Schedule/handler/date"
	//"Schedule/model"
)
var tmp bool

func timeAndUser()(int,int,int,string){
	year:=time.Now().Year()
	month:=datechange.MonthTrans(time.Now().Month().String())
	day:=time.Now().Day()
	claim,_:=token.ParseToken(user.User.Token)
	userid:=claim.ID
	return year,month,day,userid
}//取得当前时间和user

// func Finish(){
// 	for {
// 		year:=date.Date.Year
// 		month:=date.Date.Month
// 		day:=date.Date.Day
// 		claim,_:=token.ParseToken(user.User.Token)
// 		userid:=claim.ID
// 		//time.Sleep(100000000*10)//睡十秒
// 		schedules:=model.LookSchedule(year,month,day,userid)//当前日期的所有任务
// 		ok:=true
// 		for _,schedule:=range schedules{
// 			if schedule.Done==false{
// 				ok=false
// 			}
// 		}
// 		model.FinishAllSchedule(year,month,day,userid,ok)
// 	}
// }//检测今天的任务做完了没有
//丢给前端干

func OpenDoor(){
	
}//七天后开启任意门//这是前端干的吗？毕竟感觉是前端负责渲染页面
//此处暂时搁置

//看起来这个package似乎没什么用了

