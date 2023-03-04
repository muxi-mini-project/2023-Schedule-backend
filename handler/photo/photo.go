package photo
import(
	//"Schedule/handler/date"
	//"Schedule/handler/user"
	"Schedule/model"	
	//"Schedule/service/token"
	"Schedule/service/getId"
	"Schedule/service/photo"
	"Schedule/service/setDate"
	"github.com/gin-gonic/gin"
)

// @Summary 添加图片
// @Tags photo
// @Accept application/json
// @Produce application/json
// @Param photo formData file true "图片"
// @Success 200 {object} model.Url
// @Failure 401 {object} model.Fundmt
// @Router /calendar/addphoto [post]
func AddPhoto(c *gin.Context){
	userid:=getId.GetId(c)
	photo,err:=c.FormFile("photo")
	if err != nil{
		c.JSON(401,gin.H{
			"code":10010,
			"message":err.Error(),
		})
		return
	}
	code,url:= photos.UploadQiniu(photo)
	c.JSON(200, gin.H{
		"code": code,
		"message":"ok",
		"url":  url,
	})//告诉前端我们存好啦
	//把url存到数据库里
	y,m,d:=setDate.GetDate(c)
	model.PhotoUrlInsert(y,m,d,userid,url,1)//所以我们需要在加照片时，再加一个坐标。
}

// @Summary 添加图片2
// @Tags photo
// @Accept application/json
// @Produce application/json
// @Param photo formData file true "图片"
// @Success 200 {object} model.Url
// @Failure 401 {object} model.Fundmt
// @Router /calendar/addphoto2 [post]
func AddPhoto2(c *gin.Context){
	userid:=getId.GetId(c)
	photo,err:=c.FormFile("photo")
	if err != nil{
		c.JSON(401,gin.H{
			"code":10010,
			"message":err.Error(),
		})
		return
	}
	code,url:= photos.UploadQiniu(photo)
	c.JSON(200, gin.H{
		"code": code,
		"message":"ok",
		"url":  url,
	})//告诉前端我们存好啦
	//把url存到数据库里
	y,m,d:=setDate.GetDate(c)
	model.PhotoUrlInsert(y,m,d,userid,url,2)//所以我们需要在加照片时，再加一个坐标。
}

// @Summary 添加图片3
// @Tags photo
// @Accept application/json
// @Produce application/json
// @Param photo formData file true "图片"
// @Success 200 {object} model.Url
// @Failure 401 {object} model.Fundmt
// @Router /calendar/addphoto3 [post]
func AddPhoto3(c *gin.Context){
	userid:=getId.GetId(c)
	photo,err:=c.FormFile("photo")
	if err != nil{
		c.JSON(401,gin.H{
			"code":10010,
			"message":err.Error(),
		})
		return
	}
	code,url:= photos.UploadQiniu(photo)
	c.JSON(200, gin.H{
		"code": code,
		"message":"ok",
		"url":  url,
	})//告诉前端我们存好啦
	//把url存到数据库里
	y,m,d:=setDate.GetDate(c)
	model.PhotoUrlInsert(y,m,d,userid,url,3)//所以我们需要在加照片时，再加一个坐标。
}

// @Summary 添加图片4
// @Tags photo
// @Accept application/json
// @Produce application/json
// @Param photo formData file true "图片"
// @Success 200 {object} model.Url
// @Failure 401 {object} model.Fundmt
// @Router /calendar/addphoto4 [post]
func AddPhoto4(c *gin.Context){
	userid:=getId.GetId(c)
	photo,err:=c.FormFile("photo")
	if err != nil{
		c.JSON(401,gin.H{
			"code":10010,
			"message":err.Error(),
		})
		return
	}
	code,url:= photos.UploadQiniu(photo)
	c.JSON(200, gin.H{
		"code": code,
		"message":"ok",
		"url":  url,
	})//告诉前端我们存好啦
	//把url存到数据库里
	y,m,d:=setDate.GetDate(c)
	model.PhotoUrlInsert(y,m,d,userid,url,4)//所以我们需要在加照片时，再加一个坐标。
}