package setDate
import(
	"strconv"
	"github.com/gin-gonic/gin"
)
func SetDate(c *gin.Context)(int,int,int){
	year,_:=strconv.Atoi(c.GetHeader("Year"))
	month,_:=strconv.Atoi(c.GetHeader("Month"))
	day,_:=strconv.Atoi(c.GetHeader("Day"))
	return year,month,day
}
func GetDate(c *gin.Context)(int,int,int){
	year,_:=c.Get("Year")
	month,_:=c.Get("Month")
	day,_:=c.Get("Day")
	y,_:=year.(int)
	m,_:=month.(int)
	d,_:=day.(int)
	return y,m,d
}