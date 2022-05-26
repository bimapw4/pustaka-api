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

	book := book.Book{}
	// ==============================
	// create book
	// ==============================
	// book.Title = "Atomic Habbit"
	// book.Price = 9000
	// book.Rating = 5
	// book.Description = "ini buku bagus
	// db.Create(&book)

	// ==============================
	// read one book
	// ==============================
	// err = db.First(&book, 2).Error
	// if err != nil {
	// 	fmt.Printf("error get data")
	// }
	// fmt.Println("title : ", book.Title)
	// fmt.Println("data : ", book)

	// ==============================
	// read all book
	// ==============================
	// err = db.Find(&book).Error
	// if err != nil {
	// 	fmt.Printf("erorr get all book")
	// }
	// // fmt.Println(book)
	// for _, v := range book {
	// 	fmt.Println("title = ", v.Title)
	// 	fmt.Println("data = ", v)
	// }

	// ==============================
	// update book
	// ==============================
	// err = db.Where("id = ? ", 1).First(&book).Error
	// book.Title = "Man Tiger (lalala)"
	// db.Save(&book)
	// fmt.Println("data = ", book)
	// if err != nil {
	// 	fmt.Printf("error update data")
	// }

	// ==============================
	// Delete book
	// ==============================
	err = db.Debug().Where("id = ?", 1).Delete(&book).Error
	if err != nil {
		fmt.Printf("error delete book")
	}
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
