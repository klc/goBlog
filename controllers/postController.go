package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/mkilic91/goBlog/models"
	"net/http"
)

func GetPost(context *gin.Context) {

	fmt.Println(context.Params)
	var post Post
	var categories []Category

	post.Slug = context.Param("slug")
	DB.Where(&post).First(&post)

	DB.Select("title, slug").Find(&categories)

	context.HTML(http.StatusOK, "pages/post.html", gin.H{"title": post.Title, "post": post, "categories": categories})
}
