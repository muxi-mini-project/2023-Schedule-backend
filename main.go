package main
import(
	//"fmt"
	"Schedule/model"
	"Schedule/router"
	"github.com/gin-gonic/gin"
)
func main(){
	model.Init()
	defer model.DB.Close()

	r:=gin.Default()
	router.Router(r)
	r.Run(":114")
}