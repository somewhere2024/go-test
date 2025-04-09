package api

import (
	"fmt"
	"gin--/internal/models"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"reflect"
)

func GetHeader(c *gin.Context) {
	header := c.Request.Header
	fmt.Println(header)

	c.JSON(200, gin.H{"message": "test", "data": header})
}

func GetPath(c *gin.Context) {
	user := &models.UserTest{}
	err := c.ShouldBindUri(user)
	if err != nil {
		c.JSON(200, gin.H{"message": "test", "code": http.StatusBadRequest, "data": user})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "test", "code": http.StatusOK, "data": user})
}

func GetQuery(c *gin.Context) {
	user := &models.UserTest{}
	err := c.ShouldBindQuery(user)

	if err != nil {
		c.JSON(200, gin.H{"message": "test", "code": http.StatusBadRequest, "data": user})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "test", "code": http.StatusOK, "data": user})

}

func PostForm(c *gin.Context) {
	user := &models.UserTest{}
	err := c.ShouldBind(user)

	if err != nil {
		c.JSON(200, gin.H{"message": "test", "code": http.StatusBadRequest, "data": user})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "test", "code": http.StatusOK, "data": user})

}

func PostJson(c *gin.Context) {
	user := &models.UserTest{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(200, gin.H{"message": "test", "code": http.StatusBadRequest, "data": user})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "test", "code": http.StatusOK, "data": user})
}

func BingUser(c *gin.Context) {
	user := &models.UserTest{}
	err := c.ShouldBind(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "test", "code": http.StatusOK, "data": user})
}
func TestRe(c *gin.Context) {
	user := models.User{
		Username: "test",
		Password: "123456",
	}
	newUser := reflect.ValueOf(&user)
	fmt.Println(newUser)
	c.JSON(200, gin.H{"message": "test", "data": user})
}

func FIleLoad(c *gin.Context) {
	file, _ := c.FormFile("file")
	fmt.Println("filesize(kB)", file.Size/1024)

	fileOpen, err := file.Open()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	data, _ := io.ReadAll(fileOpen)
	fmt.Println(string(data))
	c.JSON(http.StatusOK, gin.H{"message": "test", "fileSize": file.Size / 1024})
}

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	//没有文件夹会自动创建
	err := c.SaveUploadedFile(file, "./uploads/"+file.Filename)

	//dst, _ := os.Create(fmt.Sprintf("./uploads/%s", file.Filename))
	//defer dst.Close()
	//fileOpen, _ := file.Open()
	//io.Copy(dst, fileOpen)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "上传成功", "fileSize": file.Size / 1024})
}

func UploadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	for _, file := range files {
		err := c.SaveUploadedFile(file, "./uploads/"+file.Filename)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
			return
		}
	}
	c.String(http.StatusOK, "上传成功")
}

func DownloadFile(c *gin.Context) {
	filepath := c.Param("filepath")
	c.Header("Content-Type", "application/octet-stream")              // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
	c.Header("Content-Disposition", "attachment; filename="+filepath) // 用来指定下载下来的文件名
	c.Header("Content-Transfer-Encoding", "binary")
	c.File("./uploads/" + filepath)
}

func MiddlewareTest1(c *gin.Context) {
	fmt.Println("中间件1")
	c.Abort() //仅阻止后面的中间件，当前的中间件继续执行
	c.Next()
	c.JSON(200, gin.H{"message": "middleware"})
}
func MiddlewareTest2(c *gin.Context) {
	fmt.Println("中间件2")
}
func MiddlewareTest3(c *gin.Context) {
	fmt.Println("中间件3")
}
