package controllers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Filename string `uri:"name" binding:"required"`
}

func FindByName (ctx *gin.Context) {
	var params Params

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{ "msg": err.Error() })
		return
	}

	path := fmt.Sprintf("data/%s", params.Filename)
	file, err := os.Open(path)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "msg": err.Error() })
		return
	}

	defer file.Close()

	//read file data
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "msg": err.Error() })
		return
	}

	ctx.JSON(http.StatusOK, gin.H{ "msg": "success" })
}