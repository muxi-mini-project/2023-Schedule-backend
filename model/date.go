package model
import(
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Memory(year string,month string,day string, uid string){
	var dateplus DatePlus
	flag:=QueryDate(year,month,day,uid)
	if flag==false{//创建
		dateplus.Year =year
		dateplus.Month =month
		dateplus.Day =day
		dateplus.UserId =uid
		dateplus.Memory=true
		DB.Create(&dateplus)
	}else{//更新
		DB.Where("year=? and month=? and day=? and uid = ?",year,month,day,uid).Find(&dateplus)
		DB.Model(&dateplus).Update("Memory",!dateplus.Memory)
	}
}

func QueryDate(year string,month string,day string, uid string)(bool){
	var date1 DatePlus
	DB.Where("year=? and month=? and day=? and uid = ?",year,month,day,uid).Find(&date1)
	if date1.Year!=""{
		return true
	}else{
		return false
	}
}