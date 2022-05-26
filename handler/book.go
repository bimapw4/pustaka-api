package handler

import (
	"Belajar-Golang/book"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostData(ctx *gin.Context) {
	bookInput := book.Buku{}

	err := ctx.ShouldBindJSON(&bookInput)

	if err != nil {

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition %s", e.Field(), e.ActualTag())
			ctx.JSON(http.StatusBadRequest, errorMessage)
			fmt.Println(err)

			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookInput,
	})

}
