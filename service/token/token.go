package token
import(
	"time"
	"Schedule/model"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)
//生成token
func GenerateToken(uid string)(string,error){
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
//检查token
func ParseToken(token string)(*model.Claims,error){
	tokenClaims,err:=jwt.ParseWithClaims(token,&model.Claims{},
		func(token *jwt.Token)(interface{},error){
			return []byte("Hello"),nil
	})
	if err != nil{
		return nil,err
	}
	if tokenClaims != nil{
		claims,ok := tokenClaims.Claims.(*model.Claims)
		if ok && tokenClaims.Valid{
			return claims,nil//在这里得到了claims，这里面有id
		}
	}
	return nil,err
}
//一个token先验//我也不知道我在讲什么
func GetToken(c *gin.Context)string{
	tokenString:=c.GetHeader("Authorization")                                        //啊啊啊为什么是authorization啊，没想明白
	if tokenString == "" {
		return "0"
	}
	claims, err := ParseToken(tokenString)
	if err != nil{
		return "0"
	}
	return claims.Id
}