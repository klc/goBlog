package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mkilic91/goBlog/models"
	"net/http"
)

func GetPost(context *gin.Context) {

	var post models.Post
	post.Slug = context.Param("slug")
	models.DB.Where(&post).First(&post)

	context.HTML(http.StatusOK, "pages/post.html", gin.H{"title": post.Title, "post": post})
}
