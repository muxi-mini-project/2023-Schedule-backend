package photo
import(
	"Schedule/handler/date"
	"Schedule/handler/user"
	"Schedule/model"	
	"Schedule/service/token"
	"Schedule/service/photo"
	"github.com/gin-gonic/gin"
)
func AddPhoto(c *gin.Context){
	//加一个验token的东西
	claim,err1:=token.ParseToken(user.User.Token)
	if err1 != nil{
		c.JSON(400,gin.H{
			"token":err1,
		})
		return
	}else{
		userid:=claim.ID
		photo,err:=c.FormFile("photo")
		if err != nil{
			c.JSON(400,gin.H{
				"code":10010,
				"error":err.Error(),
			})
			return
		}
		code,url:= photos.UploadQiniu(photo)
		c.JSON(200, gin.H{
			"code": code,
			"url":  url,
		})//告诉前端我们存好啦
		//把url存到数据库里
		model.PhotoUrlInsert(date.Date.Year,date.Date.Month,
			date.Date.Day,userid,url)
	}
}