package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mkilic91/goBlog/routes"
	"html/template"
	"time"
)

func init() {

}
func main() {

	app := gin.Default()
	app.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	app.LoadHTMLGlob("views/**/*.html")
	app.Static("/assets/", "./assets")

	routes.Routes(app)

	err := app.Run()

	if err != nil {
		panic("route tun error")
	}
}

func formatAsDate(t time.Time) string {

	return t.Format("02 Jan 2006 3:04 PM")
}
