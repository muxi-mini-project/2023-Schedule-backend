package getId
import(
	"github.com/gin-gonic/gin"
)
func GetId(c *gin.Context) string {
	ID, _ := c.Get("UserId")
	id, _ := ID.(string)
	return id
}