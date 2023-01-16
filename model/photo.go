package model
import(
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func PhotoUrlInsert(year int,month int,day int,
		uid string,url string){
	photo:=Photo{
		PhotoUrl:url,
	}
	photo.Year=year
	photo.Month=month
	photo.Day=day
	photo.UserId=uid
	DB.Create(&photo)
}//添加图片url

func LookPhoto(year int,month int,day int,
		uid string)([]Photo){
	var photo []Photo
	DB.Where("year=? and month=? and day=? and user_id = ?",year,month,day,uid).Find(&photo)
	return photo
}//查看图片，返回json，里面有url给前端