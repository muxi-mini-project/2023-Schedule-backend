package door
import(
	"strconv"
	"github.com/gin-gonic/gin"
	"Schedule/model"
	//"Schedule/service/token"
	//"Schedule/handler/user"
	"Schedule/service/getId"
)

// @Summary 打开任意门
// @Description 查看当天的图片与做完的任务
// @Tags door
// @Accept application/json
// @Produce application/json
// @Param year path string true "年"
// @Param month path string true "月"
// @Param day path string true "日"
// @Success 200 {object} model.SchAndPh
// @Router /preview/{year}/{month}/{day} [get]
func Preview (c *gin.Context){
	year,_ :=strconv.Atoi(c.Param("year"))
	month,_:=strconv.Atoi(c.Param("month"))
	day,_:=strconv.Atoi(c.Param("day"))
	userid:=getId.GetId(c)
	allschedhule:=model.LookSchedule(year,month,day,userid)
	allphoto:=model.LookPhoto(year,month,day,userid)
	c.JSON(200,gin.H{
		"code":200,
		"message":"ok",
		"schedule":allschedhule,
		"photo":allphoto,
	})
}