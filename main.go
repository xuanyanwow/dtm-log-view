package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//f, err := os.Open("D:\\Code\\Go\\dtm\\dtm.log")
	//if err != nil {
	//	panic(err)
	//}
	//defer f.Close()
	//
	//
	//// TODO 根据配置文件+正则  获取日志文件列表
	//// TODO 读取指定日志文件全部到内存
	//
	//var lineText string
	//scanner := bufio.NewScanner(f)
	//max := 3
	//nowLine := 0
	//
	//for scanner.Scan() {
	//	nowLine++
	//	if nowLine > max {
	//		break
	//	}
	//	lineText = scanner.Text()
	//
	//	fmt.Println(lineText)
	//}
	//
	//fmt.Println("扫描结束")

	r := setupRouter()
	r.Run(":8080")
}


func setupRouter() *gin.Engine {
	r := gin.Default()
	// home page

	r.StaticFile("/", "./templates/index.html")
	r.StaticFile("/index.html", "./templates/index.html")

	// file list api
	r.GET("/file/get_list", func(c *gin.Context) {
		c.String(http.StatusOK, "get_list")
	})

	// file content api
	r.GET("/file/get_content", func(c *gin.Context) {
		c.String(http.StatusOK, "get_content")
	})
	return r
}