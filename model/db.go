package model
import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"github.com/spf13/viper"
)
var DB *gorm.DB
func Init(){
	DB = open()
}
func open()*gorm.DB{
	config:="root:yzx999.@/mini"
	db,err:=gorm.Open("mysql",config)
	if err!= nil {
		fmt.Printf("init database error:%v\n",err)
		panic(err)
	}else{
		fmt.Printf("init success!\n")
	}
	return db
}