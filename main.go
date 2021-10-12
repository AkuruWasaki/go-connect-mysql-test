package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Open database
db, err := sql.Open("mysql", "dbname=go_database user=go_test")
if err != nil {
  return err
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.Run()
}
