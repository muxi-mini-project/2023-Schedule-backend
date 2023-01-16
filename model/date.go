package model
import(
	//"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Memory(year int,month int,day int, uid string){
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

func IfMemory(year int,month int,day int, uid string)(bool){
	var dateplus DatePlus
	DB.Where("year=? and month=? and day=? and uid = ?",year,month,day,uid).Find(&dateplus)
	return dateplus.Memory
}

func QueryDate(year int,month int,day int, uid string)(bool){
	var date1 DatePlus
	DB.Where("year=? and month=? and day=? and uid = ?",year,month,day,uid).Find(&date1)
	if date1.UserId!=""{
		return true//有这条记录
	}else{
		return false//没有这条记录
	}
}

// func LookAllScheduleDone(year int,month int,day int,uid string)bool{
// 	var date1 DatePlus
// 	DB.Where("year=? and month=? and day=? and uid = ?",year,month,day,uid).Find(&date1)
// 	return date1.Done
// }//看看今日的任务写完了没（是看！）

// func FinishAllSchedule(year int,month int,day int,uid string,ok bool){
// 	var date1 DatePlus
// 	DB.Where("year=? and month=? and day=? and uid = ?",year,month,day,uid).Find(&date1)
// 	DB.Model(&date1).Update("Done",ok)
// }//修改今天的任务是否做完（是改）