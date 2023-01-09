package user
import(
	"fmt"
	//"time"
	//"strconv"
	"Schedule/model"
	"Schedule/service/token"
	"Schedule/model/user"
	"github.com/gin-gonic/gin"
	//"Schedule/handler/date"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
)

var User model.User//全局变量User位于handler

func Login(c *gin.Context){
	uid:=c.PostForm("uid")
	password:=c.PostForm("password")
	pwd:=user.QueryUserPwd(uid)
	if password != "" {
		if pwd == password{
			var err error
			User.Token,err=token.GenerateToken(uid)//生成token和一个错误
			if err!=nil{
				c.String(400,"登陆失败，请重试")
				fmt.Printf("token err:%v\n",err)
			}else{
				c.String(200,"登陆成功,token:%s",User.Token)
			}
		}else{
		 	c.String(200,"密码不正确")
		}
	}else{
		c.String(200,"该用户不存在")
	}
}

