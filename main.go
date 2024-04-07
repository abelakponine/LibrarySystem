package main

import (
	"LibrarySystem/models"
	"LibrarySystem/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	book := models.BookModel{
		Id:            1,
		Title:         "King of Devs",
		Author:        "Abel Akponine",
		DatePublished: "17-04-1993",
		YearPublished: 2024,
		SerialNumber:  "AA0001",
	}

	updatedLibrary := models.CreateBook(&book)

	newBook := book
	newBook.Id = 2
	newBook.Title = "New Dawn"
	newBook.SerialNumber = "AA0002"

	updatedLibrary = models.AddToLibrary(updatedLibrary, newBook)

	for _, book := range *updatedLibrary {
		fmt.Println(book)
	}

	router := gin.Default()
	routers.InitializeRouter(router.Group("development"), "Development")
	routers.InitializeRouter(router.Group("app"))
	router.Run()
}
