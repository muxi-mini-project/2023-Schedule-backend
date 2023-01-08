package door
import(
	"github.com/gin-gonic/gin"
	"Schedule/model"
	"Schedule/service/token"
	"Schedule/handler/user"
)
func Preview (c *gin.Context){
	year :=c.PostForm("year")
	month:=c.PostForm("month")
	day:=c.PostForm("day")
	claim,_:=token.ParseToken(user.User.Token)
	userid:=claim.ID
	allschedhule:=model.LookSchededule(year,month,day,userid)
	c.JSON(200,allschedhule)
	//allphoto:=model.LookPhoto(year,month,day,userid)图片之后再写
}