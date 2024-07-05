package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type FileData struct {
	Name string
	Size float64
}

func FindAll (ctx *gin.Context) {
	folder, err := os.Open("data/")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "msg": err.Error() })
		return
	}

	defer folder.Close()

	files, err := folder.Readdir(-1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "msg": err.Error() })
		return
	} 

	//list all files and return them
	var items []FileData

	for _, file := range files {
		size := float64(file.Size()) / (1024 * 1024)
		items = append(items, FileData{file.Name(), size})
	}

	ctx.JSON(http.StatusOK, gin.H{ "msg": "success", "files": items })
}