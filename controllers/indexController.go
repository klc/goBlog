package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mkilic91/goBlog/models"
	"net/http"
	"strconv"
)

func GetIndex(context *gin.Context) {
	var posts []models.Post

	pageParam := context.Param("page")
	page, err := strconv.Atoi(pageParam)

	if err != nil {
		page = 1
	}

	models.DB.Select("id", "slug")

	paginator := models.Paging(&(models.Param{
		DB:      models.DB,
		Page:    page,
		Limit:   5,
		OrderBy: []string{"id desc"},
	}), &posts)

	context.HTML(http.StatusOK, "pages/index.html", gin.H{"title": "deneme", "posts": posts, "paginator": paginator})
}
