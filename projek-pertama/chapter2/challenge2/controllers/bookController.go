package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
	BookID int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{
	{BookID: 1, Title: "Golang", Author: "Gopher", Desc: "A book for Go"},
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = len(BookDatas) + 1
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Created",
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].BookID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Updated",
	})
}

func GetAllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &BookDatas)
}

func GetBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID, _ := strconv.Atoi(ctx.Param("bookID"))
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookID),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Deleted",
	})
}
