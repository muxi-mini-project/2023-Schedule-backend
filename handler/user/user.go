package user
import(
	"fmt"
	"time"
	"Schedule/model"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var User model.User

func queryUser(uid string)string{
	var user model.User
	model.DB.Where("uid = ?",uid).First(&user)
	return user.Password
}

func Login(c *gin.Context){
	uid:=c.PostForm("uid")
	password:=c.PostForm("password")
	pwd:=queryUser(uid)
	if password != "" {
		if pwd == password{
			//c.SetCookie("name",username,3600,"/","localhost",false,true)
			//改cookie为token//还没跑
			var err error
			User.Token,err=generateToken(uid)
			if err!=nil{
				c.String(400,"登陆失败，请重试")
				fmt.Printf("token err:%v\n",err)
			}else{
				c.String(200,"登陆成功")
			}
			
		}else{
			c.String(200,"密码不正确")
		}
	}else{
		c.String(200,"该用户不存在")
	}
}

func generateToken(uid string)(string,error){
	nowTime := time.Now()
	expireTime := nowTime.Add(24*60*60*time.Second)
	issuer := uid
	claims := model.Claims{
		ID:		uid,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt:	expireTime.Unix(),
			Issuer:		issuer,
		},
	}//看起来，这似乎是个接口
	token,err := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims).SignedString([]byte("Hello"))
	return token,err
}

func Write(c *gin.Context){
	
}

func Check(c *gin.Context){
	
}

func Tear(c *gin.Context){
	
}