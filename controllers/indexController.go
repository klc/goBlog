package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/mkilic91/goBlog/models"
	"net/http"
	"strconv"
)

func GetIndex(context *gin.Context) {
	var posts []Post
	var postQuery = DB
	var categories []Category
	var category = context.Query("category")

	pageParam := context.Param("page")
	page, err := strconv.Atoi(pageParam)

	if err != nil {
		page = 1
	}

	DB.Select("id, title, slug").Find(&categories)

	if category != "" {
		var find uint = 0

		for _, c := range categories {
			if c.Slug == category {
				find = c.ID
			}
		}

		if find != 0 {
			postQuery = DB.Where(Post{CategoryId: find})
		}
	}

	paginator := Paging(&(Param{
		DB:      postQuery,
		Page:    page,
		Limit:   5,
		OrderBy: []string{"id desc"},
	}), &posts)

	context.HTML(http.StatusOK, "pages/index.html", gin.H{"title": "deneme", "posts": posts, "paginator": paginator, "categories": categories})
}
