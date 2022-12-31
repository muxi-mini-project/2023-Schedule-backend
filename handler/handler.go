package handler
import(
	"fmt"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context){
	fmt.Printf("Welcome!\n")
	c.JSON(200,gin.H{
		"status":"Welcome!",
	})
}

func PostPhotos(c *gin.Context){
	
}