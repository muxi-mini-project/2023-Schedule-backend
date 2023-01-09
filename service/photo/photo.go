package photos
import(
	"context"
	//"github.com/gin-gonic/gin"
	"Schedule/config"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	_ "github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	_ "github.com/qiniu/go-sdk/v7/storage"
	//"github.com/spf13/viper"
	"mime/multipart"
	_ "mime/multipart"
	//"net/http"
	//"strings"
	"time"
)
func UploadQiniu(file *multipart.FileHeader) (int, string) {
	src, err := file.Open()//src才是真正的文件，file不是文件本身
	if err != nil {
		return 10011, err.Error()
	}//转义失败
	defer src.Close()
	putPolicy := storage.PutPolicy{
		Scope: config.Q.Bucket,
	}//存到这里
	mac := qbox.NewMac(config.Q.AccessKey, config.Q.SecretKey)
	upToken := putPolicy.UploadToken(mac)	// 获取上传凭证
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}	// 配置参数
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}        // 上传返回后的结果
	putExtra := storage.PutExtra{} // 额外参数
	key := "(" + time.Now().String() + ")" + file.Filename// 自定义文件名及后缀
	if err := formUploader.Put(context.Background(), &ret,
		upToken, key, src, file.Size, &putExtra); err != nil {
		return 501, err.Error()
	}
	return 0, config.Q.Domain + "/" + ret.Key
}
