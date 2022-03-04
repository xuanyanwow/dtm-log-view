package controller

import (
	"bufio"
	"dtm-log-view/config"
	"dtm-log-view/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetList(c *gin.Context) {

	files, err := ioutil.ReadDir(config.Config.LogFilePath)
	if err != nil {
		log.Println(err)
	}
	var fileList []string
	for _, file := range files {
		// TODO 判断是否符合正则条件
		fileList = append(fileList, config.Config.LogFilePath+string(os.PathSeparator)+file.Name())
	}

	c.JSON(http.StatusOK, utils.JsonResponse{
		Code: "200",
		Data: map[string]interface{}{
			"list": fileList,
		},
		Msg: "ok",
	})
}

func GetContent(c *gin.Context) {
	path := c.PostForm("path")
	log.Println(path)

	// 转int格式
	page, _ := strconv.Atoi(c.PostForm("page"))
	limit, _ := strconv.Atoi(c.PostForm("limit"))

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		c.JSON(http.StatusOK, utils.JsonResponse{
			Code: "400",
			Data: map[string]interface{}{},
			Msg:  "file not found",
		})
		return
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lineText string
	var logArray []string

	scanner := bufio.NewScanner(f)
	nowLine := 0
	haveLine := 0
	startLine := (page - 1) * limit

	// 指定行开始，读取limit行
	for scanner.Scan() {
		nowLine++
		if nowLine <= startLine {
			continue
		}
		if haveLine >= limit {
			break
		}
		lineText = scanner.Text()
		haveLine++

		logArray = append(logArray, lineText)
	}

	c.JSON(http.StatusOK, utils.JsonResponse{
		Code: "200",
		Data: map[string]interface{}{
			"list": logArray,
		},
		Msg: "ok",
	})
}
