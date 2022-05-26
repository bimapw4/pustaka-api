package main

import (
	"fmt"
	"log"
	"net/http"

	"Belajar-Golang/book"
	"Belajar-Golang/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connect eror")
	}

	fmt.Println("DB Connect Success")

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)

	books, err := bookRepository.FindByID(2)
	if err != nil {
		fmt.Printf("error get book")
	}

	fmt.Println("data = ", books)
	router := gin.Default()

	v1 := router.Group("/v1")

	routerV1(v1)

	v1.POST("/book", handler.PostData)

	router.Run()
}

func routerV1(v1 *gin.RouterGroup) {

	v1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": "hello World",
		})
	})

	v1.GET("/:params", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": ctx.Param("params"),
		})
	})

	v1.GET("/query", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": ctx.Query("id"),
		})
	})

}
