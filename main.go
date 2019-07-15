package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func main() {
	// updates everything from github
	pullFromGithub := exec.Command("git", "pull")
	pullFromGithub.Run()

	router := gin.Default()

	// serves all the html, css, and javascript files to the browser
	router.StaticFile("/", "./web/html/index.html")
	router.StaticFS("/html", http.Dir("./web/html/"))
	router.StaticFS("/css", http.Dir("./web/css/"))
	router.StaticFS("/js", http.Dir("./web/js/"))
	router.Run()
}

func ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
