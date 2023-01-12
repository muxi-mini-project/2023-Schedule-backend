package door
import(
	"strconv"
	"github.com/gin-gonic/gin"
	"Schedule/model"
	"Schedule/service/token"
	"Schedule/handler/user"
)

// @Summary 打开任意门
// @Description 查看当天的图片与做完的任务
// @Tag door
// @Accept application/json
// @Produce application/json
// @Param year path string true "年"
// @Param month path string true "月"
// @Param day path string true "日"
// @Success 200 {object} []model.Schedule "{"allschedules":allschedule}"
// @Success 200 {object} []model.Photo "{"allphotos":allphoto}"
// @Failure 400 {string} string "{"token":err}"
// @Router /preview/{year}/{month}/{day} [get]
func Preview (c *gin.Context){
	year,_ :=strconv.Atoi(c.Param("year"))
	month,_:=strconv.Atoi(c.Param("month"))
	day,_:=strconv.Atoi(c.Param("day"))
	claim,err:=token.ParseToken(user.User.Token)
	if err != nil{
		c.JSON(400,gin.H{
			"token":err,
		})
	}else{
		userid:=claim.ID
		allschedhule:=model.LookSchedule(year,month,day,userid)
		allphoto:=model.LookPhoto(year,month,day,userid)
		c.JSON(200,gin.H{
			"allSchedules":allschedhule,
			"allPhotos":allphoto,
		})
	}
}