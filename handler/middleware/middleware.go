package middleware
import(
	"Schedule/service/token"
	"Schedule/service/setDate"
	"github.com/gin-gonic/gin"
)

// @Summary  身份验证
// @Description  获取token并验证，成功则将Set UserId
// @Tags middleware
// @Accept application/json
// @Produce application/json
// @Param token header string false "token"
// @Failure 401 {object} model.Fundmt
func TokenMiddle(c *gin.Context) {
	UserId := token.GetToken(c)//得到的就是id，不是结构体
	if UserId == "0" {
		c.JSON(401, gin.H{
			"code": 401, 
			"messsage": "权限不足",
		})
		c.Abort()
		return
		//到此一游
	} else {
		c.Set("UserId", UserId)
		c.Next()
		//根据 id 找到对应用户信息并返回
		//后期取id的时候使用id,err:=c.Get(UserId)取出id即可
	}
}
//再写一个设定时间的中间件//还得再写这个的接口

// @Summary  设置时间获得
// @Description  获取date并验证，成功则将Set
// @Tags middleware
// @Accept application/json
// @Produce application/json
// @Param year header string false "token"
// @Failure 401 {object} model.Fundmt
func DateMiddle(c *gin.Context){
	Year,Month,Day:=setDate.SetDate(c)//返回的是int，别记错了
	if Year == 0 {
		c.JSON(401, gin.H{
			"code": 401, 
			"messsage": "权限不足",
		})
		c.Abort()
		return
		//到此一游
	} else {
		c.Set("Year", Year)
		c.Set("Month", Month)
		c.Set("Day", Day)
		c.Next()
		//根据 id 找到对应用户信息并返回
		//后期取id的时候使用id,err:=c.Get(UserId)取出id即可
	}
}
