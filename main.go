package main
import(
	//"fmt"
	"Schedule/config"
	"Schedule/model"
	"Schedule/router"
	//"Schedule/service/punch"
	"github.com/gin-gonic/gin"
)

// @title library API
// @version 1.0
// @description 我的日程本API
// @host localhost
// @BasePath /api/v1
func main(){
	err:=config.Init("")//这个是init config.yaml文件
	if err!=nil{
		panic(err)
	}
	config.LoadQiniu()

	model.Init()
	defer model.DB.Close()
	model.InitTables()

	r:=gin.Default()
	router.Router(r)
	r.Run(":114")
}