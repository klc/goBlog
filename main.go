package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
}

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err = gorm.Open("mysql", dbUrl)

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Post{})
}

func main() {

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLGlob("views/**/*.html")
	router.Static("/assets/", "./assets")

	router.GET("/", getIndex)
	router.GET("/p/:page", getIndex)
	router.GET("/c/:slug", getPost)

	err := router.Run()

	if err != nil {
		panic("route tun error")
	}
}

func getIndex(context *gin.Context) {
	var posts []Post

	pageParam := context.Param("page")
	page, err := strconv.Atoi(pageParam)

	if err != nil {
		page = 1
	}

	db.Select("id", "slug")

	paginator := Paging(&(Param{
		DB:      db,
		Page:    page,
		Limit:   5,
		OrderBy: []string{"id desc"},
	}), &posts)

	context.HTML(http.StatusOK, "pages/index.html", gin.H{"title": "deneme", "posts": posts, "paginator": paginator})
}

func getPost(context *gin.Context) {
	var post Post
	post.Slug = context.Param("slug")
	db.Where(&post).First(&post)

	context.HTML(http.StatusOK, "pages/post.html", gin.H{"title": post.Title, "post": post})
}

func formatAsDate(t time.Time) string {

	return t.Format("02 Jan 2006 3:04 PM")
}
