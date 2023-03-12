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

//var User model.User//全局变量User位于handler

// @Summary 登录
// @Description 使用一站式
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param user formData model.User true "登录的用户信息"
// @Success 200 {object} model.Token
// @Failure 402 {object} model.Fundmt
// @Failure 401 {object} model.Fundmt
// @Failure 400 {object} model.Fundmt
// @Router /login [post]
func Login(c *gin.Context){
	var userr model.User
	err1:=c.Bind(&userr)
	if err1 != nil{
		c.JSON(402,gin.H{
			"code":402,
			"message":"输入信息格式有误",
		})
		return
	}
	if userr.UID == ""{
		c.JSON(402,gin.H{
			"code":402,
			"message":"用户名不为空",
		})
		return
	}
	pwd:=userr.Password//此时userr.Password还是用户输进来的内容，pwd也是
	//但在下一行if执行完后，userr.Password就已经是数据库里的东西了（已经变成哈希的形状了）
	if resu := model.DB.Where("uid = ?", userr.UID).First(&userr); resu.Error != nil {
		_, err := model.GetUserInfoFormOne(userr.UID, pwd)//这里如果是user.UID的话，会让编译器理解不了这个user到底是变量名还是包名……呵呵
		if err != nil {//用户第一次登录
			// c.Abort()
			c.JSON(401, gin.H{
				"code":401,
				"message":"Password or account wrong.",
			})
			return
		}
		userr.Password = user.GeneratePasswordHash(pwd) //加密//记得写加密
		model.DB.Create(&userr)//存数据库
	}else{//用户不是第一次登录
		//pwd时用户输进来的密码
		if  !user.CheckPassword(pwd, userr.Password){
			c.JSON(401, gin.H{
				"code":401,
				"message":"Password or account wrong.",
			})
			return//注册过，但是密码错了
		}
	}
	//账号和密码正确后生成一个token
	tokenStr,err:=token.GenerateToken(userr.UID)//生成token和一个错误
	if err!=nil{
		c.JSON(400,gin.H{
			"code":400,
			"message":"登陆失败，请重试",
		})
		fmt.Printf("token err:%v\n",err)
	}else{
		c.JSON(200,gin.H{
			"code":200,
			"message":"登陆成功",
			"token":tokenStr,
		})
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