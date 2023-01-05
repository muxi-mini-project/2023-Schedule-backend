package punch
import(
	"time"
	"strconv"
	//"github.com/gin-gonic/gin"
	"Schedule/service/token"
	"Schedule/handler/user"
)
var tmp bool

func timeAndUser()(string,string,string,string){
	year:=strconv.Itoa(time.Now().Year())
	month:=time.Now().Month().String()
	day:=strconv.Itoa(time.Now().Day())
	claim,_:=token.ParseToken(user.User.Token)
	userid:=claim.ID
	return year,month,day,userid
}
func Memory(){
	// year,month,day,uid:=timeAndUser()
	// ok:=IfMemory(year,month,day,uid)
	
}//每日检测今天是否纪念

func Finish(){

}//每日检测今天的任务做完了没有

func OpenDoor(){
	
}//七天后开启任意门

func OutDate(){
	
}//定期删除超过两周的
