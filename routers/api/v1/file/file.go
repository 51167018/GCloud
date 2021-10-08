package file

import (
	"GCloud/pkg/conf"
	"GCloud/pkg/errcode"
	"GCloud/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func UploadChunk(c *gin.Context) {
	fileHash := c.PostForm("hash")
	file, err := c.FormFile("file")
	hashPath := fmt.Sprintf("upload/%s", fileHash)
	if err != nil {
		fmt.Println("获取上传文件失败", err)
	}

	isExistPath, err := util.PathExists(hashPath)
	if err != nil {
		fmt.Println("获取hash路径错误", err)
	}
	if !isExistPath {
		os.Mkdir(hashPath, os.ModePerm)
	}
	err = c.SaveUploadedFile(file, fmt.Sprintf("upload/%s/%s", fileHash, file.Filename))
	if err != nil {
		c.String(400, "0")
		fmt.Println(err)
	} else {
		chunkList := []string{}
		files, err := ioutil.ReadDir(hashPath)
		if err != nil {
			fmt.Println("文件读取错误", err)
		}
		for _, f := range files {
			fileName := f.Name()

			if f.Name() == ".DS_Store" {
				continue
			}
			chunkList = append(chunkList, fileName)
		}

		c.JSON(200, gin.H{
			"chunkList": chunkList,
		})
	}
}

//CheckChunk 检查文件是否分片上传成功
func CheckChunk(c *gin.Context)  {
	hash := c.Query("hash")
	hashPath := fmt.Sprintf("upload/%s", hash)
	chunkList := []string{}
	isExistPath, err := util.PathExists(hashPath)
	if err != nil {
		fmt.Println("获取hash路径错误", err)
	}

	if isExistPath {
		files, err := ioutil.ReadDir(hashPath)
		state := 0
		if err != nil {
			fmt.Println("文件读取错误", err)
		}
		for _, f := range files {
			fileName := f.Name()
			chunkList = append(chunkList, fileName)
			fileBaseName := strings.Split(fileName, ".")[0]
			if fileBaseName == hash {
				state = 1
			}
		}

		c.JSON(200, gin.H{
			"state": state,
			"chunkList": chunkList,
		})
	} else {
		c.JSON(200, gin.H{
			"state": 0,
			"chunkList": chunkList,
		})
	}
}

func MegerChunk(c *gin.Context)  {
	hash := c.Query("hash")
	fileName := c.Query("fileName")
	hashPath := fmt.Sprintf("upload/%s", hash)

	isExistPath, err := util.PathExists(hashPath)
	if err != nil {
		log.Println("获取hash路径错误", err)
	}

	if !isExistPath {
		util.NewResponse(c).ToResponse(errcode.FOLDER_NOT_FOUND,nil)
		return
	}

	isExistFile, err := util.PathExists(hashPath + "/" + fileName)
	if err != nil {
		log.Println("获取hash路径文件错误", err)
	}
	fmt.Println("文件是否存在", isExistFile)
	if isExistFile {
		util.NewResponse(c).ToResponse(errcode.SUCCESS,fmt.Sprintf("http://"+conf.ServerHost+":"+conf.ServerPort+"/%s/%s", hash, fileName))
		return
	}

	files, err := ioutil.ReadDir(hashPath)
	if err != nil {
		log.Println("合并文件读取失败", err)
	}
	complateFile, err := os.Create(hashPath + "/" + fileName)
	defer complateFile.Close()
	for _, f := range files {

		if f.Name() == ".DS_Store" {
			continue
		}

		fileBuffer, err := ioutil.ReadFile(hashPath + "/" + f.Name())
		if err != nil {
			fmt.Println("文件打开错误", err)
		}
		complateFile.Write(fileBuffer)
	}

	util.NewResponse(c).ToResponse(errcode.SUCCESS,fmt.Sprintf("http://+"+conf.ServerHost+":"+conf.ServerPort+"/%s/%s", hash, fileName))

}
