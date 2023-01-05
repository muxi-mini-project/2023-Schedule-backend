package user
import(
	//"github.com/jinzhu/gorm"
	"Schedule/model"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func QueryUserPwd(uid string)string{
	var user model.User
	model.DB.Where("uid = ?",uid).First(&user)
	return user.Password
}