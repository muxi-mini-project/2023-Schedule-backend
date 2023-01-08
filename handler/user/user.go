package user
import(
	"fmt"
	//"time"
	//"strconv"
	"Schedule/model"
	"Schedule/service/token"
	"Schedule/model/user"
	"github.com/gin-gonic/gin"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

var User model.User//全局变量User位于handler

func Login(c *gin.Context){
	uid:=c.PostForm("uid")
	password:=c.PostForm("password")
	//pwd:=queryUserPwd(uid)
	if password != "" {
		//if pwd == password{
			var err error
			User.Token,err=token.GenerateToken(uid)//生成token和一个错误
			if err!=nil{
				c.String(400,"登陆失败，请重试")
				fmt.Printf("token err:%v\n",err)
			}else{
				c.String(200,"登陆成功,token:%s",User.Token)
			}
		// }else{
		// 	c.String(200,"密码不正确")
		// }
	}else{
		c.String(200,"该用户不存在")
	}
}
//写入一条记录
func Write(c *gin.Context){
	content:=c.PostForm("schedule")
	//接下来把这一块写到model里面
	err:=user.WriteSchedule(content,User.Token)
	if err!=nil{
		c.JSON(400,err)
	}else{
		c.JSON(200,gin.H{
			"status":"ok",
		})
	}
}
//完成一个任务或者否认它的完成
func Check(c *gin.Context){
	claim,err:=token.ParseToken(User.Token)
	if err == nil{
		userid:=claim.ID
		content:=c.Param("content")
		done:=model.Check(userid,content)
		if done == true{
			c.JSON(200,gin.H{"status":"true"})
		}else{
			c.JSON(200,gin.H{"status":"true"})
		}
	}else{
		fmt.Printf("token err:%v",err)
		c.JSON(400,err)
	}
}
//撕去一天的日历
func Tear(c *gin.Context){
	
}