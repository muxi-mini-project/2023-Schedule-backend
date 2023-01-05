package main
import(
	//"fmt"
	"Schedule/model"
	"Schedule/router"
	"Schedule/service/punch"
	"github.com/gin-gonic/gin"
)
func main(){
	model.Init()
	go routine()
	defer model.DB.Close()
	//model.InitTables()

	r:=gin.Default()
	router.Router(r)
	r.Run(":114")
}
func routine(){
	//go punch.Memory()
	go punch.Finish()
	go punch.OpenDoor()
	go punch.OutDate()
}