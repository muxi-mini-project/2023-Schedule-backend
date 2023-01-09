package door
import(
	"strconv"
	"github.com/gin-gonic/gin"
	"Schedule/model"
	"Schedule/service/token"
	"Schedule/handler/user"
)
func Preview (c *gin.Context){
	year,_ :=strconv.Atoi(c.Param("year"))
	month,_:=strconv.Atoi(c.Param("month"))
	day,_:=strconv.Atoi(c.Param("day"))
	claim,_:=token.ParseToken(user.User.Token)
	userid:=claim.ID
	allschedhule:=model.LookSchedule(year,month,day,userid)
	c.JSON(200,allschedhule)
	allphoto:=model.LookPhoto(year,month,day,userid)
	c.JSON(200,allphoto)
}