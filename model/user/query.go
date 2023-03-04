package user
import(
	//"github.com/jinzhu/gorm"
	// "time"
	// "strconv"
	"Schedule/model"
	"github.com/ShiinaOrez/GoSecurity/security"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func QueryUserPwd(uid string)string{
	var user1 model.User
	model.DB.Where("uid = ?",uid).First(&user1)
	return user1.Password
}//查密码

func GeneratePasswordHash(password string) string {
	return security.GeneratePasswordHash(password)
}//生成哈希加密后的密码

func CheckPassword(password, hashPwd string) bool {
	return security.CheckPasswordHash(password, hashPwd)
}//比较输入的密码和存储的、哈希加密过后的密码
