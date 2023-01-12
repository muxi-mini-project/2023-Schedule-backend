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
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var User model.User//全局变量User位于handler

// @Summary 登录
// @Description 使用一站式
// @Tag user
// @Accept application/json
// @Produce application/json
// @Param user formData model.User true "登录的用户信息"
// @Success 200 {string} string "登陆成功"
// @Failure 402 {string} string "{"message":"输入信息格式有误"} or {"message":"学号不为空"}"
// @Failure 401 {string} string "{"message":"Password or account wrong."}"
// @Failure 400 {string} string "{"message":"token err"}"
// @Router /login [post]
func Login(c *gin.Context){
	err1:=c.Bind(&User)
	if err1 != nil{
		c.JSON(402,gin.H{
			"message":"输入信息格式有误",
		})
		return
	}
	if User.UID == ""{
		c.JSON(402,gin.H{
			"message":"用户名不为空",
		})
		return
	}
	pwd:=User.Password
	if resu := model.DB.Where("uid = ?", User.UID).First(&User); resu.Error != nil {
		_, err := model.GetUserInfoFormOne(User.UID, pwd)
		if err != nil {
			// c.Abort()
			c.JSON(401, "Password or account wrong.")
			return
		}
		//User.Password = model.GeneratePasswordHash(pwd) //加密
		model.DB.Create(&User)//存数据库
	}else{
		pwd := user.QueryUserPwd(User.UID)
		if  pwd != User.Password{
			c.JSON(401, "Password or account wrong.")
			return//注册过，但是密码错了
		}
	}
	//账号和密码正确后生成一个token
	var err error
	User.Token,err=token.GenerateToken(User.UID)//生成token和一个错误
	if err!=nil{
		c.JSON(400,"登陆失败，请重试")
		fmt.Printf("token err:%v\n",err)
	}else{
		c.JSON(200,"登陆成功")
	}
}
//这里都是一些旧的代码
// func Register(c *gin.Context){
// 	var user model.User
// 	uid:=c.PostForm("uid")
// 	password:=c.PostForm("password")
// 	if uid != "" && password != ""{
// 		user.UID = uid
// 		user.Password = password
// 		u:=model.User{UID:user.UID,Password:password}
// 		model.DB.Create(&u)
// 		c.String(200,"注册成功")
// 	}else{
// 		c.String(200,"用户名与密码不能为空")
// 	}
// }

// func Login(c *gin.Context){
// 	uid:=c.PostForm("uid")
// 	password:=c.PostForm("password")
// 	pwd:=user.QueryUserPwd(uid)
// 	if pwd != "" {
// 		if pwd == password{
// 			var err error
// 			User.Token,err=token.GenerateToken(uid)//生成token和一个错误
// 			if err!=nil{
// 				c.String(400,"登陆失败，请重试")
// 				fmt.Printf("token err:%v\n",err)
// 			}else{
// 				c.String(200,"登陆成功")
// 			}
// 		}else{
// 		 	c.String(200,"密码不正确")
// 		}
// 	}else{
// 		c.String(200,"该用户不存在")
// 	}
// }