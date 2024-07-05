package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload (ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "msg": err.Error() })
		return
	}

	//update filename 
	extensions := strings.Split(file.Filename, ".")
	extension := extensions[len(extensions)-1]

	file.Filename = fmt.Sprintf("%v.%s", time.Now().UnixNano() / int64(time.Millisecond), extension)
	destination := filepath.Join("data/", file.Filename)

	err = ctx.SaveUploadedFile(file, destination)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "msg": err.Error() })
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{ "msg": "file has been created", "name": file.Filename })
}