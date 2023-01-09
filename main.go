package main
import(
	//"fmt"
	"Schedule/config"
	"Schedule/model"
	"Schedule/router"
	//"Schedule/service/punch"
	"github.com/gin-gonic/gin"
)
func main(){
	err:=config.Init("")//这个是init config.yaml文件
	if err!=nil{
		panic(err)
	}
	config.LoadQiniu()

	model.Init()
	//go routine()
	defer model.DB.Close()
	//model.InitTables()

	r:=gin.Default()
	router.Router(r)
	r.Run(":114")
}
// func routine(){
// 	//go punch.Memory()
// 	//punch.Finish()
// 	//go punch.OpenDoor()
// }
/*我要写什么？

*/